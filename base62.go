package branca

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//                  MIT License <https://opensource.org/licenses/MIT>                 //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"math/big"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrNonBase62Char means that given string contains symbol which is not a part of
// the Base62 alphabet
var ErrNonBase62Char = errors.New("Got non Base62 character")

// ////////////////////////////////////////////////////////////////////////////////// //

// EncodeBase62 encodes bytes slice to Base62 encoded string
func EncodeBase62(src []byte) string {
	if len(src) == 0 {
		return ""
	}

	var b62 big.Int

	b62.SetBytes(src)
	d := []byte(b62.Text(62))

	for i, r := range d {
		if r >= 65 && r <= 90 {
			d[i] = r + 32
		} else if r >= 97 && r <= 122 {
			d[i] = r - 32
		}
	}

	return string(d)
}

// DecodeBase62 decodes Base62 encoded string to byte slice
func DecodeBase62(src string) ([]byte, error) {
	if src == "" {
		return []byte{}, nil
	}

	d := []byte(src)

	for i, r := range d {
		if r >= 65 && r <= 90 {
			d[i] = r + 32
		} else if r >= 97 && r <= 122 {
			d[i] = r - 32
		}
	}

	var b62 big.Int

	_, ok := b62.SetString(string(d), 62)

	if !ok {
		return nil, ErrNonBase62Char
	}

	return b62.Bytes(), nil
}
