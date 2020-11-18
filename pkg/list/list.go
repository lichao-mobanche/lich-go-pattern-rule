package list

//unfinished list

//New TODO
func New() *List {
	return &List{0, nil, nil}
}

// List TODO
type List struct {
	size  int
	first *Item
	last  *Item
}

// Item TODO
type Item struct {
	Value    interface{}
	Next     *Item
	Previous *Item
	List     *List
}

func (it *Item) clear() {
	it.Next = nil
	it.Previous = nil
	it.List = nil
}

// PopBack TODO
func (l *List) PopBack() (r interface{}) {
	last := l.last
	if last != nil {
		l.last = l.last.Previous
		r = last.Value
		last.clear()
		l.size--
	}
	return r
}

// PopFront TODO
func (l *List) PopFront() (r interface{}) {
	first := l.first
	if first != nil {
		l.first = l.first.Next
		r = first.Value
		first.clear()
		l.size--
	}
	return r
}

// PushBack TODO
func (l *List) PushBack(v interface{}) *Item {
	newItem := &Item{v, nil, l.last, l}
	l.last.Next = newItem
	l.last = l.last.Next
	l.size++
	return newItem
}

// PushFront TODO
func (l *List) PushFront(v interface{}) *Item {
	newItem := &Item{v, l.first, nil, l}
	l.first.Previous = newItem
	l.first = l.first.Previous
	l.size++
	return newItem
}

// InsertAfter TODO
func (l *List) InsertAfter(v interface{}, item *Item) *Item {
	newItem := &Item{v, item.Next, item, l}
	item.Next = newItem
	l.size++
	return newItem
}

// InsertBefore TODO
func (l *List) InsertBefore(v interface{}, item *Item) *Item {
	newItem := &Item{v, item, item.Previous, l}
	item.Previous = newItem
	l.size++
	return newItem
}

// Size TODO
func (l *List) Size() int {
	return l.size
}

//Erase TODO
func (l *List) Erase(item *Item) interface{} {
	if l == item.List {
		p := item.Previous
		n := item.Next
		if p != nil && n != nil {
			p.Next = item.Next
			n.Previous = item.Previous
		}
		if p == nil {
			n.Previous = nil
		}
		if n == nil {
			p.Next = nil
		}
		item.clear()
		l.size--
	}
	return item.Value
}

//Front TODO
func (l *List) Front() *Item {
	return l.first
}

//Back TODO
func (l *List) Back() *Item {
	return l.last
}
