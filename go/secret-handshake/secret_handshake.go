package secret

var moves = []string{"wink", "double blink", "close your eyes", "jump"}

const REVERSEMASK = 0b10000

// Handshake returns the moves of a handshake based on code.
func Handshake(code uint) []string {
	handshake := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		if code&(1<<i) != 0 {
			handshake = append(handshake, moves[i])
		}
	}
	if code&REVERSEMASK != 0 {
		for i, j := 0, len(handshake)-1; i < j; i, j = i+1, j-1 {
			handshake[i], handshake[j] = handshake[j], handshake[i]
		}
	}
	return handshake
}
