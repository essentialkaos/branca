// +build gofuzz

package branca

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
	_, err := DecodeBase62(data)

	if err != nil {
		return 1
	}

	return 1
}
