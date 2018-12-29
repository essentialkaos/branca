// Package for working with Branca tokens
package branca

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"time"

	"golang.org/x/crypto/chacha20poly1305"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Branca is branca struct
type Branca struct {
	key []byte
	ttl uint32
}

// Token is branca token
type Token struct {
	payload []byte
	ts      uint32
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrInvalidToken   = errors.New("Token is invalid")
	ErrInvalidVersion = errors.New("Token has invalid version")
	ErrBadKeyLength   = errors.New("Key must be 32 bytes long")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// nonceReadFunc is function for reading random nonce
var nonceReadFunc func(d []byte) (int, error) = rand.Read

// ////////////////////////////////////////////////////////////////////////////////// //

// brancaVersion is current branca version
var brancaVersion = []byte{0xBA}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewBranca creates new branca struct
func NewBranca(key []byte) (*Branca, error) {
	if len(key) != 32 {
		return nil, ErrBadKeyLength
	}

	return &Branca{
		key: key,
	}, nil
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

// ////////////////////////////////////////////////////////////////////////////////// //

// SetTTL sets Time To Live
func (b *Branca) SetTTL(ttl uint32) {
	b.ttl = ttl
}

// IsExpired returns true if given token is expired
func (b *Branca) IsExpired(t Token) bool {
	return uint32(time.Now().Unix()) > t.ts+b.ttl
}

// Encode encodes payload to branca token
func (b *Branca) Encode(payload []byte) ([]byte, error) {
	nonce, err := genNonce()

	if err != nil {
		return nil, err
	}

	ts := make([]byte, 4)
	binary.BigEndian.PutUint32(ts, uint32(time.Now().Unix()))

	// Version (1B) || Timestamp (4B) || Nonce (24B) || Ciphertext (*B) || Tag (16B)
	token := append(brancaVersion, ts...)
	token = append(token, nonce...)

	aead, err := chacha20poly1305.NewX(b.key)

	if err != nil {
		return nil, err
	}

	cipher := aead.Seal(nil, nonce, payload, token)
	token = append(token, cipher...)

	return token, nil
}

// Decode extract payload from branca token
func (b *Branca) Decode(token []byte) (Token, error) {
	if len(token) < 45 {
		return Token{}, ErrInvalidToken
	}

	if token[0] != brancaVersion[0] {
		return Token{}, ErrInvalidVersion
	}

	header := token[0:29]
	cipher := token[29:]
	ts := binary.BigEndian.Uint32(header[1:5])
	nonce := header[5:]

	aead, err := chacha20poly1305.NewX(b.key)

	if err != nil {
		return Token{}, err
	}

	payload, err := aead.Open(nil, nonce, cipher, header)

	if err != nil {
		return Token{}, err
	}

	return Token{payload, ts}, nil
}

// EncodeToString create Base62 encoded token with given payload
func (b *Branca) EncodeToString(payload []byte) (string, error) {
	token, err := b.Encode(payload)

	if err != nil {
		return "", err
	}

	return EncodeBase62(token), nil
}

// DecodeString extract payload from Base62 encoded token
func (b *Branca) DecodeString(token string) (Token, error) {
	data, err := DecodeBase62(token)

	if err != nil {
		return Token{}, err
	}

	return b.Decode(data)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// genNonce generates random nonce
func genNonce() ([]byte, error) {
	nonce := make([]byte, 24)

	_, err := nonceReadFunc(nonce)

	return nonce, err
}
