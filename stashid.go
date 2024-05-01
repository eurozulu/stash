package stash

import (
	"fmt"
)

type StashID [2]int

func (id StashID) Page() int {
	return id[0]
}

func (id StashID) Offset() int {
	return id[1]
}

func (id StashID) String() string {
	if id.Page() < 0 || id.Offset() < 1 {
		return "INVALID-ID"
	}
	return fmt.Sprintf("%d:%d", id.Page(), id.Offset())
}

func (id StashID) Equals(oid StashID) bool {
	return id.Page() == oid.Page() && id.Offset() == oid.Offset()
}

func newStashId(page, offset int) StashID {
	return StashID([...]int{page, offset})
}
