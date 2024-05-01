package stash

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
)

type ByteValue []byte

func (bv ByteValue) ID() StashId {
	return sha1.Sum(bv)
}

type StashId [20]byte

func (bv StashId) Equals(oid StashId) bool {
	return bytes.Equal(bv[:], oid[:])
}

func (bv StashId) String() string {
	return hex.EncodeToString(bv[:])
}
