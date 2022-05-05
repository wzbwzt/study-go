package iterator

type IntIterator struct {
	*Iterator
	Data []int
}

func (ii *IntIterator) Next() int {
	defer func() {
		ii.Index++
	}()
	return ii.Data[ii.Index]
}
