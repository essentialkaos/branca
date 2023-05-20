//go:build gofuzz
// +build gofuzz

package branca

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//                  MIT License <https://opensource.org/licenses/MIT>                 //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

var b, _ = NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

// ////////////////////////////////////////////////////////////////////////////////// //

func FuzzEncode(data []byte) int {
	_, err := b.Encode(data)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzDecode(data []byte) int {
	_, err := b.Decode(data)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzEncodeBase62(data []byte) int {
	s := EncodeBase62(data)

	if s == "" {
		return 1
	}

	return 0
}

func FuzzDecodeBase62(data []byte) int {
	_, err := DecodeBase62(string(data))

	if err != nil {
		return 1
	}

	return 1
}
