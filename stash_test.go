package stash

import "testing"

var testBytes = [][]byte{
	[]byte("Hello World"),
	[]byte("Goodbye and thanks for the fish"),
	[]byte("The rain in Spain stays mainly in the plane\nThe Plane is rain stays mainly in Spain."),
	[]byte("123456"),
	[]byte("true"),
	[]byte("false"),
	[]byte("ha ha ha ha"),
}
var testByteIDs = buildByteIDs(testBytes)

func TestMemoryStash_Put(t *testing.T) {
	testPageSize := 5
	s := NewStash(testPageSize)
	ids := s.Put(testBytes...)
	if len(ids) != len(testBytes) {
		t.Errorf("Unexpected return count putting in bytes.  Expected %d, found %d", len(testBytes), len(ids))
	}
	if s.Length() != len(testBytes) {
		t.Errorf("Unexpected length putting in bytes.  Expected %d, found %d", len(testBytes), s.Length())
	}

	// put in the same data and ensure length remains unchanged
	ids = s.Put(testBytes...)
	if len(ids) != len(testBytes) {
		t.Errorf("Unexpected return count (re)putting in bytes.  Expected %d, found %d", len(testBytes), len(ids))
	}
	if s.Length() != len(testBytes) {
		t.Errorf("Unexpected length (re)putting in bytes.  Expected %d, found %d", len(testBytes), s.Length())
	}
}

func buildByteIDs(byz [][]byte) []ByteID {
	ids := make([]ByteID, len(byz))
	for i, b := range byz {
		ids[i] = ByteValue(b).ID()
	}
	return ids
}
