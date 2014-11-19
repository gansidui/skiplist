package skiplist

type Interface interface {
	Less(other interface{}) bool
}

type SkipList struct {
	header *Element
	tail   *Element
	update []*Element
	length int
	level  int
}

// New returns an initialized skiplist.
func New() *SkipList {
	return &SkipList{
		header: newElement(SKIPLIST_MAXLEVEL, nil),
		tail:   nil,
		update: make([]*Element, SKIPLIST_MAXLEVEL),
		length: 0,
		level:  1,
	}
}

// Init initializes or clears skiplist sl.
func (sl *SkipList) Init() *SkipList {
	sl.header = newElement(SKIPLIST_MAXLEVEL, nil)
	sl.tail = nil
	sl.update = make([]*Element, SKIPLIST_MAXLEVEL)
	sl.length = 0
	sl.level = 1
	return sl
}

// Front returns the first elements of skiplist sl or nil.
func (sl *SkipList) Front() *Element {
	return sl.header.level[0].forward
}

// Back returns the last elements of skiplist sl or nil.
func (sl *SkipList) Back() *Element {
	return sl.tail
}

// Len returns the numbler of elements of skiplist sl.
func (sl *SkipList) Len() int {
	return sl.length
}

// Insert inserts v, increments sl.length, and returns a new element of wrap v.
func (sl *SkipList) Insert(v Interface) *Element {
	node := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && node.level[i].forward.Value.Less(v) {
			node = node.level[i].forward
		}
		sl.update[i] = node
	}

	level := randomLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			sl.update[i] = sl.header
		}
		sl.level = level
	}

	node = newElement(level, v)
	for i := 0; i < level; i++ {
		node.level[i].forward = sl.update[i].level[i].forward
		sl.update[i].level[i].forward = node
	}

	if sl.update[0] != sl.header {
		node.backward = sl.update[0]
	}
	if node.level[0].forward != nil {
		node.level[0].forward.backward = node
	} else {
		sl.tail = node
	}
	sl.length++

	return node
}

// deleteElement deletes e from its skiplist, and decrements sl.length.
func (sl *SkipList) deleteElement(e *Element, update []*Element) {
	for i := 0; i < sl.level; i++ {
		if update[i].level[i].forward == e {
			update[i].level[i].forward = e.level[i].forward
		}
	}

	if e.level[0].forward != nil {
		e.level[0].forward.backward = e.backward
	} else {
		sl.tail = e.backward
	}

	for sl.level > 1 && sl.header.level[sl.level-1].forward == nil {
		sl.level--
	}
	sl.length--
}

// Remove removes e from sl if e is an element of skiplist sl.
// It returns the element value e.Value.
func (sl *SkipList) Remove(e *Element) interface{} {
	if e == nil {
		return nil
	}

	node := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && node.level[i].forward.Value.Less(e.Value) {
			node = node.level[i].forward
		}
		sl.update[i] = node
	}

	node = node.level[0].forward
	if node == e {
		sl.deleteElement(node, sl.update)
		return node.Value
	}

	return nil
}

// Delete deletes e if e.Value == v, and return e.Value.
func (sl *SkipList) Delete(v Interface) interface{} {
	node := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && node.level[i].forward.Value.Less(v) {
			node = node.level[i].forward
		}
		sl.update[i] = node
	}

	node = node.level[0].forward
	if node != nil && !node.Value.Less(v) && !v.Less(node.Value) {
		sl.deleteElement(node, sl.update)
		return node.Value
	}

	return nil
}

// Find finds e if e.Value == v, and return e.
func (sl *SkipList) Find(v Interface) *Element {
	node := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && node.level[i].forward.Value.Less(v) {
			node = node.level[i].forward
		}
	}

	node = node.level[0].forward
	if node != nil && !node.Value.Less(v) && !v.Less(node.Value) {
		return node
	}

	return nil
}
