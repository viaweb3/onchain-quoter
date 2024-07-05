package token

import (
	"bytes"
	"encoding/gob"
)

type Token struct {
	Symbol   string
	Address  string
	Decimals int64
}

func NewToken(name string, address string, decimals int64) Token {
	return Token{
		Symbol:   name,
		Address:  address,
		Decimals: decimals,
	}
}

func Decode(poolBytes []byte) (Token, error) {
	buf := bytes.NewBuffer(poolBytes)
	dec := gob.NewDecoder(buf)

	var t Token
	if err := dec.Decode(&t); err != nil {
		return Token{}, err
	}

	return t, nil
}

func (t Token) Encode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(t); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}
