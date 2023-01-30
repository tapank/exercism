package variablelengthquantity

import "errors"

const (
	MASK = 0b01111111
	MSB  = 0b10000000
)

// EncodeVarint encodes the input numbers each into bytes
func EncodeVarint(input []uint32) (out []byte) {
	for _, num := range input {
		// set the first byte
		bytes := []byte{byte(num & MASK)}

		// set remaining if needed
		for num >>= 7; num > 0; num >>= 7 {
			bytes = append([]byte{byte(num&MASK) | MSB}, bytes...)
		}
		out = append(out, bytes...)
	}
	return
}

// DecodeVarint decodes the input bytes one or more numbers
func DecodeVarint(input []byte) (out []uint32, err error) {
	// validate input
	if len(input) > 0 && input[len(input)-1]&MSB != 0 {
		err = errors.New("incomplete byte sequence")
	}

	// build numbers and add to output slice
	var num uint32
	for _, b := range input {
		num |= uint32(b & MASK)
		if b&MSB == 0 {
			out = append(out, num)
			num = 0
		} else {
			num <<= 7
		}
	}
	return
}
