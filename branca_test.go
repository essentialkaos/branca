package branca

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//                  MIT License <https://opensource.org/licenses/MIT>                 //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/rand"
	"errors"
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type BrancaSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&BrancaSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *BrancaSuite) TestCreation(c *C) {
	_, err := NewBranca([]byte("abcd"))

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrBadKeyLength)
}

func (s *BrancaSuite) TestEncoding(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	token, err := brc.Encode([]byte("TEST1234abcdАБВГ"))

	c.Assert(err, IsNil)
	c.Assert(token, Not(HasLen), 0)
}

func (s *BrancaSuite) TestEncodingToString(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	token, err := brc.EncodeToString([]byte("TEST1234abcdАБВГ"))

	c.Assert(err, IsNil)
	c.Assert(token, Not(HasLen), 0)
}

func (s *BrancaSuite) TestEncodingAEADError(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	brc = Branca([]byte("abcd"))

	_, err = brc.Encode([]byte("TEST1234abcdАБВГ"))

	c.Assert(err, NotNil)
}

func (s *BrancaSuite) TestEncodingNonceError(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	nonceReadFunc = func(d []byte) (int, error) {
		return -1, errors.New("ERROR")
	}

	_, err = brc.Encode([]byte("TEST1234abcdАБВГ"))

	c.Assert(err, NotNil)

	nonceReadFunc = rand.Read
}

func (s *BrancaSuite) TestEncodingToStringAEADError(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	brc = Branca([]byte("abcd"))

	_, err = brc.EncodeToString([]byte("TEST1234abcdАБВГ"))

	c.Assert(err, NotNil)
}

func (s *BrancaSuite) TestDecoding(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	data, _ := brc.Encode([]byte("TEST1234abcdАБВГ"))
	token, err := brc.Decode(data)

	c.Assert(err, IsNil)
	c.Assert(token.Payload(), DeepEquals, []byte("TEST1234abcdАБВГ"))
	c.Assert(token.Timestamp().Unix(), Not(Equals), 0)
}

func (s *BrancaSuite) TestDecodingFromString(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	data, _ := brc.EncodeToString([]byte("TEST1234abcdАБВГ"))
	token, err := brc.DecodeString(data)

	c.Assert(err, IsNil)
	c.Assert(token.Payload(), DeepEquals, []byte("TEST1234abcdАБВГ"))
}

func (s *BrancaSuite) TestDecodingAEADError(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	token, _ := brc.Encode([]byte("TEST1234abcdАБВГ"))

	brc = Branca([]byte("abcd"))

	_, err = brc.Decode(token)

	c.Assert(err, NotNil)
}

func (s *BrancaSuite) TestDecodingFromStringB62Error(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	_, err = brc.DecodeString("АБВГ")

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNonBase62Char)
}

func (s *BrancaSuite) TestDecodingExpired(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	data, err := brc.Encode([]byte("TEST1234abcdАБВГ"))

	c.Assert(err, IsNil)
	c.Assert(data, Not(HasLen), 0)

	time.Sleep(3 * time.Second)

	token, err := brc.Decode(data)

	c.Assert(err, IsNil)
	c.Assert(token.IsExpired(1), Equals, true)
}

func (s *BrancaSuite) TestDecodingSizeError(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	_, err = brc.Decode([]byte("abcd"))

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrInvalidToken)
}

func (s *BrancaSuite) TestDecodingDecryptFail(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	token, err := brc.Encode([]byte("TEST1234abcdАБВГ"))

	c.Assert(err, IsNil)
	c.Assert(token, Not(HasLen), 0)

	for i := 6; i < 22; i++ {
		token[i] = 0xFF
	}

	_, err = brc.Decode(token)

	c.Assert(err, NotNil)
}

func (s *BrancaSuite) TestVersionCheckFail(c *C) {
	brc, err := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

	c.Assert(err, IsNil)
	c.Assert(brc, NotNil)

	token, _ := brc.Encode([]byte("TEST1234abcdАБВГ"))

	token[0] = 0xFF

	_, err = brc.Decode(token)

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrInvalidVersion)
}

func (s *BrancaSuite) TestBase62Encoding(c *C) {
	c.Assert(EncodeBase62([]byte{}), Equals, "")
}

func (s *BrancaSuite) TestBase62Decoding(c *C) {
	data, err := DecodeBase62("")
	c.Assert(err, IsNil)
	c.Assert(data, DeepEquals, []byte{})
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *BrancaSuite) BenchmarkBrancaEncoding(c *C) {
	brc, _ := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))
	payload := []byte("Lorem ipsum")

	for i := 0; i < c.N; i++ {
		brc.Encode(payload)
	}
}

func (s *BrancaSuite) BenchmarkBrancaDecoding(c *C) {
	brc, _ := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))
	payload := []byte{
		186, 92, 39, 144, 77, 213, 126, 132, 225, 72, 237, 97, 100, 168, 211, 156, 67,
		88, 27, 126, 15, 165, 173, 248, 128, 2, 186, 122, 121, 172, 171, 83, 223, 74,
		139, 164, 55, 169, 150, 10, 158, 209, 216, 107, 110, 40, 17, 138, 14, 71, 85,
		21, 197, 38, 95, 191,
	}

	for i := 0; i < c.N; i++ {
		brc.Decode(payload)
	}
}

func (s *BrancaSuite) BenchmarkBrancaEncodingToString(c *C) {
	brc, _ := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))
	payload := []byte("Lorem ipsum")

	for i := 0; i < c.N; i++ {
		brc.EncodeToString(payload)
	}
}

func (s *BrancaSuite) BenchmarkBrancaDecodingFromString(c *C) {
	brc, _ := NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))
	payload := "1y3aOcgTUZQMLXDmGh9J4oAlFIr75zgTbKgTXnwgSQaskprUd4FaTe5k5XSg1jSGtzMRgu7uH690"

	for i := 0; i < c.N; i++ {
		brc.DecodeString(payload)
	}
}

func (s *BrancaSuite) BenchmarkBase62Encoding(c *C) {
	payload := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit")

	for i := 0; i < c.N; i++ {
		EncodeBase62(payload)
	}
}

func (s *BrancaSuite) BenchmarkBase62Decoding(c *C) {
	payload := "C7i3aVJocxTgVcgLMnka3X6x46LZZY1DOU9lh43UhgPZCkX0B17tUaxYJq63cZRY6k3GpUEAui"

	for i := 0; i < c.N; i++ {
		DecodeBase62(payload)
	}
}
