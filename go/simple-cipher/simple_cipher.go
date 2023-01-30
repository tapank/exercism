package cipher

type shift struct {
	key int
}
type vigenere struct {
	key []int
}

// modifier indicating forward or backward encoding
type direction int

const (
	FWD direction = 1
	BCK direction = -1
)

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance >= 26 || distance <= -26 {
		return nil
	}
	return shift{distance}
}

func (c shift) Encode(input string) string {
	return Shift([]int{c.key}, FWD, input)
}

func (c shift) Decode(input string) string {
	return Shift([]int{c.key}, BCK, input)
}

func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}

	// create key
	rkey := []int{}

	// validate key
	valid := false
	for _, r := range key {
		switch {
		case r >= 'b' && r <= 'z':
			valid = true
		case r == 'a':
			// a is valid but not good enough to flip isgood
		default:
			// illegal character encountered
			return nil
		}
		rkey = append(rkey, int(r-'a'))
	}

	// return nil if key invalid
	if !valid {
		return nil
	}
	return vigenere{rkey}
}

func (v vigenere) Encode(input string) string {
	return Shift(v.key, FWD, input)
}

func (v vigenere) Decode(input string) string {
	return Shift(v.key, BCK, input)
}

func Shift(key []int, dir direction, input string) string {
	pos := 0
	shifted := []rune{}
	for _, r := range input {
		if r >= 'A' && r <= 'Z' {
			r |= 0b100000
		}
		if r >= 'a' && r <= 'z' {
			by := key[pos] * int(dir)
			pos++
			if pos == len(key) {
				pos = 0
			}
			r += rune(by)
			if by > 0 && r > 'z' {
				r -= 26
			} else if by < 0 && r < 'a' {
				r += 26
			}
			shifted = append(shifted, r)
		}
	}
	return string(shifted)
}
