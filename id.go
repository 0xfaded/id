package id

import (
	"crypto"
	"encoding/hex"
	"strconv"

	"appengine"
	"appengine/datastore"
)

type Id interface {
	String() string
	Key(c appengine.Context, kind string, parent *datastore.Key) *datastore.Key
}

type StringId string
type IntId int64

func (id StringId) String() string {
	return string(id)
}

func (id StringId) Key(c appengine.Context, kind string, parent *datastore.Key) *datastore.Key {
	return datastore.NewKey(c, kind, id.String(), 0, parent)
}

func Hash(s string) Id {
	hash := crypto.SHA256.New()
	hash.Write([]byte(s))
	return StringId(hex.EncodeToString(hash.Sum(nil)))
}

func (id IntId) String() string {
	return strconv.FormatInt(int64(id), 10)
}

func (id IntId) Key(c appengine.Context, kind string, parent *datastore.Key) *datastore.Key {
	return datastore.NewKey(c, kind, "", int64(id), parent)
}

/*
func SumIds(id ...Id) Id {
	hash := crypto.SHA256.New()
	for _, i := range id {
		hash.Write(i.Bytes())
	}
	return Id(hex.EncodeToString(hash.Sum(nil)))
}
*/
