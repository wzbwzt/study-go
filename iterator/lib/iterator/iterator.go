package iterator

type Iterator struct {
	Len   int
	Index int
}

func (i *Iterator) HasNext() bool {
	if i.Len == 0 {
		return false
	}
	if i.Index == i.Len-1 {
		return false
	}
	return true
}
