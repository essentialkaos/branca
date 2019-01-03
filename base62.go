package branca

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"errors"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrNonBase62Char means that given string contains symbol which is not a part of
// the Base62 alphabet
var ErrNonBase62Char = errors.New("Got non Base62 character")

// ////////////////////////////////////////////////////////////////////////////////// //

// alphabet is Base62 alphabet
var alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// alphabetMap is map byte -> symbol position
var alphabetMap = map[byte]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'I': 18,
	'J': 19, 'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24, 'P': 25, 'Q': 26, 'R': 27,
	'S': 28, 'T': 29, 'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34, 'Z': 35, 'a': 36,
	'b': 37, 'c': 38, 'd': 39, 'e': 40, 'f': 41, 'g': 42, 'h': 43, 'i': 44, 'j': 45,
	'k': 46, 'l': 47, 'm': 48, 'n': 49, 'o': 50, 'p': 51, 'q': 52, 'r': 53, 's': 54,
	't': 55, 'u': 56, 'v': 57, 'w': 58, 'x': 59, 'y': 60, 'z': 61,
}

// ////////////////////////////////////////////////////////////////////////////////// //

// EncodeBase62 encodes bytes slice to base62 encoded string
func EncodeBase62(src []byte) string {
	if len(src) == 0 {
		return ""
	}

	digits := []int{0}

	for i := 0; i < len(src); i++ {
		carry := int(src[i])

		for j := 0; j < len(digits); j++ {
			carry += digits[j] << 8
			digits[j] = carry % 62
			carry = carry / 62
		}

		for carry > 0 {
			digits = append(digits, carry%62)
			carry = carry / 62
		}
	}

	var result bytes.Buffer

	for k := 0; src[k] == 0 && k < len(src)-1; k++ {
		result.WriteByte(alphabet[0])
	}

	for t := len(digits) - 1; t >= 0; t-- {
		result.WriteByte(alphabet[digits[t]])
	}

	return result.String()
}

// DecodeBase62 decodes bases62 encoded string to byte slice
func DecodeBase62(src string) ([]byte, error) {
	if src == "" {
		return []byte{}, nil
	}

	bytes := []byte{0}

	for i := 0; i < len(src); i++ {
		carry, ok := alphabetMap[src[i]]

		if !ok {
			return nil, ErrNonBase62Char
		}

		for j := 0; j < len(bytes); j++ {
			carry += int(bytes[j]) * 62
			bytes[j] = byte(carry & 0xff)
			carry >>= 8
		}

		for carry > 0 {
			bytes = append(bytes, byte(carry&0xff))
			carry >>= 8
		}
	}

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return bytes, nil
}
