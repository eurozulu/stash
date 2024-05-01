package stash

import (
	"fmt"
	"log"
)

var InvalidStashID = StashId{}

type Stash interface {
	Put(v ...[]byte) []StashId
	Get(id StashId) ([]byte, error)
	ContainsId(id StashId) bool
	Length() int
}

type stash struct {
	PageSize int
	index    map[StashId]pageIndex
	pages    []StashPage
}

func (s *stash) Length() int {
	return len(s.index)
}

func (s *stash) Put(v ...[]byte) []StashId {
	var ids []StashId
	for _, v := range v {
		bv := ByteValue(v)
		id := bv.ID()
		if !s.ContainsId(id) {
			// Doesn't already exist.  Add to current page.
			index := s.addToCurrentPage(bv)
			s.index[id] = index
		}
		ids = append(ids, id)
	}
	return ids
}

func (s stash) Get(id StashId) ([]byte, error) {
	if !s.ContainsId(id) {
		return nil, fmt.Errorf("%s is not a valid page number", id.String())
	}
	index := s.index[id]
	page := s.pages[index.Page()]
	return page.Get(index.Offset()), nil
}

func (s stash) ContainsId(id StashId) bool {
	index, ok := s.index[id]
	if !ok {
		return false
	}
	if !s.isIndexValid(index) {
		log.Printf("Invalid index: %s found, removing from index...", index.String())
		delete(s.index, id)
		return false
	}
	return true
}

func (s stash) PageCount() int {
	return len(s.pages)
}

func (s *stash) currentPageIndex() int {
	pc := s.PageCount()
	if pc == 0 || s.pages[pc-1].Count() >= s.PageSize {
		return s.addPage()
	}
	return pc - 1
}

func (s *stash) isIndexValid(index pageIndex) bool {
	pid := index.Page()
	if pid < 0 || pid >= s.PageCount() {
		return false
	}
	offset := index.Offset()
	if offset < 0 || offset >= s.pages[pid].Count() {
		return false
	}
	return true
}

func (s *stash) addPage() int {
	l := s.PageCount()
	s.pages = append(s.pages, &memoryStashPage{})
	return l
}

func (s *stash) addToCurrentPage(bv ByteValue) pageIndex {
	pid := s.currentPageIndex()
	offset := s.pages[pid].Put(bv)
	return newPageIndex(pid, offset)
}

func NewStash(pageSize int) Stash {
	return &stash{
		PageSize: pageSize,
		index:    make(map[StashId]pageIndex),
		pages:    nil,
	}
}
