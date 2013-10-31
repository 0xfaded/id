package id

import (
	"crypto"
	"encoding/hex"
)

type Id string

func (id Id) Bytes() []byte {
	if bytes, err := hex.DecodeString(string(id)); err != nil {
		panic("Invalid Id")
	} else {
		return bytes
	}
}

func Hash(s string) Id {
	hash := crypto.SHA256.New()
	hash.Write([]byte(s))
	return Id(hex.EncodeToString(hash.Sum(nil)))
}

func SumIds(id ...Id) Id {
	hash := crypto.SHA256.New()
	for _, i := range id {
		hash.Write(i.Bytes())
	}
	return Id(hex.EncodeToString(hash.Sum(nil)))
}
