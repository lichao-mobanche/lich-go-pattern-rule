package sortmap

// IterFunc TODO
type IterFunc func(Item) bool

// Iterator TODO
type Iterator interface {
	Iter(IterFunc)
}

// BaseIter TODO
type BaseIter struct {
	*Map
}

// Iter TODO
func (iter *BaseIter) Iter(f IterFunc) {
	iter.RLock()
	defer iter.RUnlock()
	elem := iter.items.Front()
	for elem != nil {
		if !f(elem.Value.(Item)) {
			break
		}
		elem = elem.Next()
	}
}

// NewBaseIter TODO
func NewBaseIter(m *Map) *BaseIter {
	return &BaseIter{m}
}
