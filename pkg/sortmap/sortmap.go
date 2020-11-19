package sortmap

import (
	"container/list"
	"sync"
)

// Item TODO
type Item interface {
	ItemID() uint64
}

// Comparer TODO
type Comparer func(interface{}, interface{}) bool

// Map TODO
type Map struct {
	items    *list.List
	idxMap   map[uint64]*list.Element
	comparer Comparer
	*sync.RWMutex
}

// New return a new Map. if is not nil, Map will be sorted.
func New(f Comparer) *Map {
	return &Map{
		items:    list.New(),
		idxMap:   make(map[uint64]*list.Element),
		comparer: f,
		RWMutex:  &sync.RWMutex{},
	}
}

// Add TODO
func (m *Map) Add(item Item) bool {
	m.Lock()
	defer m.Unlock()
	return m.add(item)
}

func (m *Map) add(item Item) bool {
	id := item.ItemID()
	if _, ok := m.idxMap[id]; ok {
		return false
	}
	elem := m.items.PushFront(item)
	if m.comparer != nil {
		iter := elem.Next()
		for iter != nil && m.comparer(iter.Value, elem.Value) {
			iter = iter.Next()
		}
		if iter == nil {
			m.items.MoveToBack(elem)
		} else {
			m.items.MoveBefore(elem, iter)
		}
	}
	m.idxMap[id] = elem
	return true
}

// Get TODO
func (m *Map) Get(id uint64) Item {
	m.RLock()
	defer m.RUnlock()
	return m.get(id)
}

func (m *Map) get(id uint64) Item {
	if item, ok := m.idxMap[id]; ok {
		return item.Value.(Item)
	}
	return nil
}

// Size TODO
func (m *Map) Size() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.idxMap)
}

func (m *Map) delete(id uint64) bool {
	if item, ok := m.idxMap[id]; ok {
		delete(m.idxMap, id)
		m.items.Remove(item)
	} else {
		return false
	}
	return true
}

// Delete TODO
func (m *Map) Delete(id uint64) bool {
	m.Lock()
	defer m.Unlock()
	return m.delete(id)
}
