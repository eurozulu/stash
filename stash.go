package stash

import "fmt"

var InvalidStashID = newStashId(-1, -1)

type Stash interface {
	Put(v ...[]byte) []StashID
	Get(id StashID) ([]byte, error)
	ContainsId(id StashID) bool
	ContainsValue(v []byte) (StashID, error)
	Length() int
}

type stash struct {
	PageSize int
	index    map[ByteID]StashID
	pages    []StashPage
}

func (s *stash) Length() int {
	return len(s.index)
}

func (s *stash) Put(v ...[]byte) []StashID {
	var ids []StashID
	for _, v := range v {
		bv := ByteValue(v)
		id, ok := s.index[bv.ID()]
		if !ok {
			// Doesn't already exist.  Add to current page.
			id = s.addToCurrentPage(bv)
		}
		ids = append(ids, id)
	}
	return ids
}

func (s stash) Get(id StashID) ([]byte, error) {
	if !s.ContainsId(id) {
		return nil, fmt.Errorf("%s is not a valid page number", id.String())
	}
	pg := s.pages[id.Page()]
	return pg.Get(id.Offset()), nil
}

func (s stash) ContainsId(id StashID) bool {
	pid := id.Page()
	if pid < 0 || pid >= s.PageCount() {
		return false
	}
	offset := id.Offset()
	if offset < 0 || offset >= s.pages[pid].Count() {
		return false
	}
	return true
}

func (s *stash) ContainsValue(v []byte) (StashID, error) {
	id, ok := s.index[ByteValue(v).ID()]
	if !ok {
		return InvalidStashID, fmt.Errorf("v not known")
	}
	return id, nil
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

func (s *stash) addPage() int {
	l := s.PageCount()
	s.pages = append(s.pages, &memoryStashPage{})
	return l
}

func (s *stash) addToCurrentPage(bv ByteValue) StashID {
	pid := s.currentPageIndex()
	offset := s.pages[pid].Put(bv)
	id := newStashId(pid, offset)
	s.index[bv.ID()] = id
	return id
}

func NewStash(pageSize int) Stash {
	return &stash{
		PageSize: pageSize,
		index:    make(map[ByteID]StashID),
		pages:    nil,
	}
}
