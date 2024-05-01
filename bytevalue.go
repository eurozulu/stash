package stash

import (
	"bytes"
	"crypto/sha1"
)

type ByteValue []byte

func (bv ByteValue) ID() ByteID {
	return sha1.Sum(bv)
}

type ByteID [20]byte

func (bv ByteID) Equals(oid ByteID) bool {
	return bytes.Equal(bv[:], oid[:])
}
