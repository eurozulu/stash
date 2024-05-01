package stash

import "testing"

func TestNewStashId(t *testing.T) {
	testPage := 1
	testOffset := 2
	expectString := "1:2"
	id := newPageIndex(testPage, testOffset)
	if id.Page() != testPage {
		t.Errorf("Unexpected page in stash id.  expected %d, found %d", testPage, id)
	}
	if id.Offset() != testOffset {
		t.Errorf("Unexpected offset in stash id.  expected %d, found %d", testOffset, id)
	}
	if id.String() != expectString {
		t.Errorf("Unexpected stash id string.  expected '%s', found '%s'", expectString, id.String())
	}
}

func TestInvalidStashId(t *testing.T) {
	testPage := -1
	testOffset := 2
	expectString := "INVALID-ID"
	id := newPageIndex(testPage, testOffset)
	if id.Page() != testPage {
		t.Errorf("Unexpected page in stash id.  expected %d, found %d", testPage, id)
	}
	if id.Offset() != testOffset {
		t.Errorf("Unexpected offset in stash id.  expected %d, found %d", testOffset, id)
	}
	if id.String() != expectString {
		t.Errorf("Unexpected stash id string.  expected '%s', found '%s'", expectString, id.String())
	}
}
