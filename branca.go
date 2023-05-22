// Package branca implements branca token specification
package branca

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//                  MIT License <https://opensource.org/licenses/MIT>                 //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"time"

	"golang.org/x/crypto/chacha20poly1305"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	// KEY_SIZE is Branca key which is always 32 bytes (i.e 256 bit)
	KEY_SIZE = 32

	// MIN_TOKEN_SIZE is minimal token size is 45 bytes (with empty payload)
	MIN_TOKEN_SIZE = 45
)

const (
	// VERSION_SIZE is Branca version size (1B)
	VERSION_SIZE = 1

	// TIMESTAMP_SIZE is Branca timestamp size (4B)
	TIMESTAMP_SIZE = 4

	// NONCE_SIZE is Branca nonce size (24B)
	NONCE_SIZE = 24

	// HEADER_SIZE is Branca header size (29B)
	HEADER_SIZE = VERSION_SIZE + TIMESTAMP_SIZE + NONCE_SIZE
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Branca is byte slice with key data
type Branca []byte

// Token is branca token
type Token struct {
	payload []byte
	ts      uint32
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrInvalidToken means that given data doesn't look like branca token
	ErrInvalidToken = errors.New("Token is invalid")

	// ErrInvalidVersion means that token has an unsupported version
	ErrInvalidVersion = errors.New("Token has invalid version")

	// ErrBadKeyLength is returned if key not equal to 32 bytes
	ErrBadKeyLength = errors.New("Key must be 32 bytes long")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// nonceReadFunc is function for reading random nonce
var nonceReadFunc func(d []byte) (int, error) = rand.Read

// ////////////////////////////////////////////////////////////////////////////////// //

// brancaVersion is current branca version
var brancaVersion = []byte{0xBA}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewBranca creates new branca struct
func NewBranca(key []byte) (Branca, error) {
	if len(key) != KEY_SIZE {
		return nil, ErrBadKeyLength
	}

	return Branca(key), nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Payload returns token payload
func (t *Token) Payload() []byte {
	return t.payload
}

// Timestamp returns token timestamp
func (t *Token) Timestamp() time.Time {
	return time.Unix(int64(t.ts), 0)
}

// IsExpired returns true if given token is expired
func (t *Token) IsExpired(ttl uint32) bool {
	return uint32(time.Now().Unix()) > t.ts+ttl
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Encode encodes payload to branca token
func (b Branca) Encode(payload []byte) ([]byte, error) {
	if len(b) != KEY_SIZE {
		return nil, ErrBadKeyLength
	}

	nonce, err := genNonce()

	if err != nil {
		return nil, err
	}

	ts := make([]byte, TIMESTAMP_SIZE)
	binary.BigEndian.PutUint32(ts, uint32(time.Now().Unix()))

	// Version (1B) || Timestamp (4B) || Nonce (24B) || Ciphertext (*B) || Tag (16B)

	var buf bytes.Buffer

	buf.Write(brancaVersion)
	buf.Write(ts)
	buf.Write(nonce)

	aead, _ := chacha20poly1305.NewX(b)
	buf.Write(aead.Seal(nil, nonce, payload, buf.Bytes()))

	return buf.Bytes(), nil
}

// Decode extract payload from branca token
func (b Branca) Decode(token []byte) (Token, error) {
	if len(b) != KEY_SIZE {
		return Token{}, ErrBadKeyLength
	}

	if len(token) < MIN_TOKEN_SIZE {
		return Token{}, ErrInvalidToken
	}

	if token[0] != brancaVersion[0] {
		return Token{}, ErrInvalidVersion
	}

	aead, _ := chacha20poly1305.NewX(b)
	payload, err := aead.Open(nil, token[VERSION_SIZE+TIMESTAMP_SIZE:HEADER_SIZE], token[HEADER_SIZE:], token[:HEADER_SIZE])

	if err != nil {
		return Token{}, err
	}

	return Token{payload, binary.BigEndian.Uint32(token[VERSION_SIZE : VERSION_SIZE+TIMESTAMP_SIZE])}, nil
}

// EncodeToString create Base62 encoded token with given payload
func (b Branca) EncodeToString(payload []byte) (string, error) {
	token, err := b.Encode(payload)

	if err != nil {
		return "", err
	}

	return EncodeBase62(token), nil
}

// DecodeString extract payload from Base62 encoded token
func (b Branca) DecodeString(token string) (Token, error) {
	data, err := DecodeBase62(token)

	if err != nil {
		return Token{}, err
	}

	return b.Decode(data)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// genNonce generates random nonce
func genNonce() ([]byte, error) {
	nonce := make([]byte, NONCE_SIZE)

	_, err := nonceReadFunc(nonce)

	return nonce, err
}
