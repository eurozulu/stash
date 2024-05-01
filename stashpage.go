package stash

// stashPage represents a shard of data in the stash.
// The page contains 0 to PageSize elements, each a ByteValue (slice of bytes)
type StashPage interface {
	Count() int
	Get(offset int) ByteValue
	Put(v ByteValue) int
}

type memoryStashPage []ByteValue

func (pg memoryStashPage) Count() int {
	return len(pg)
}

func (sp memoryStashPage) Get(offset int) ByteValue {
	return sp[offset]
}

func (sp *memoryStashPage) Put(v ByteValue) int {
	*sp = append(*sp, v)
	return len(*sp) - 1
}
