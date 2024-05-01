package stash

import (
	"fmt"
)

type pageIndex [2]int

func (id pageIndex) Page() int {
	return id[0]
}

func (id pageIndex) Offset() int {
	return id[1]
}

func (id pageIndex) String() string {
	if id.Page() < 0 || id.Offset() < 1 {
		return "INVALID-ID"
	}
	return fmt.Sprintf("%d:%d", id.Page(), id.Offset())
}

func (id pageIndex) Equals(oid pageIndex) bool {
	return id.Page() == oid.Page() && id.Offset() == oid.Offset()
}

func newPageIndex(page, offset int) pageIndex {
	return pageIndex([...]int{page, offset})
}
