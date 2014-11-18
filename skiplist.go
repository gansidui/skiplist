package skiplist

import (
	"fmt"
	"math/rand"
)

const SKIPLIST_MAXLEVEL = 8
const SKIPLIST_P = 0.25

type skiplistLevel struct {
	forward *skiplistNode
}

type skiplistNode struct {
	score    float64
	backward *skiplistNode
	level    []*skiplistLevel
}

type SkipList struct {
	header *skiplistNode
	tail   *skiplistNode
	length int
	level  int
}

type Interface interface {
	Less(lhs, rhs interface{}) bool
}

func newSkipListNode(level int, score float64) *skiplistNode {
	slLevels := make([]*skiplistLevel, level)
	for i := 0; i < level; i++ {
		slLevels[i] = new(skiplistLevel)
	}

	return &skiplistNode{
		score:    score,
		backward: nil,
		level:    slLevels,
	}
}

func randomLevel() int {
	level := 1
	for rand.Int()%2 == 1 {
		level += 1
	}

	if level < SKIPLIST_MAXLEVEL {
		return level
	} else {
		return SKIPLIST_MAXLEVEL
	}
}

func New() *SkipList {
	return &SkipList{
		level:  1,
		length: 0,
		header: newSkipListNode(SKIPLIST_MAXLEVEL, 0),
	}
}

func (sl *SkipList) Insert(score float64) *skiplistNode {
	update := make([]*skiplistNode, SKIPLIST_MAXLEVEL)
	node := sl.header

	i := 0
	for i = sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && node.level[i].forward.score < score {
			node = node.level[i].forward
		}
		update[i] = node
	}

	level := randomLevel()
	if level > sl.level {
		for i = sl.level; i < level; i++ {
			update[i] = sl.header
		}
		sl.level = level
	}

	node = newSkipListNode(level, score)
	for i = 0; i < level; i++ {
		node.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = node
	}

	if update[0] != sl.header {
		node.backward = update[0]
	}
	if node.level[0].forward != nil {
		node.level[0].forward.backward = node
	} else {
		sl.tail = node
	}
	sl.length++

	return node
}

func (sl *SkipList) deleteNode(x *skiplistNode, update []*skiplistNode) {
	for i := 0; i < sl.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].forward = x.level[i].forward
		}
	}

	if x.level[0].forward != nil {
		x.level[0].forward.backward = x.backward
	} else {
		sl.tail = x.backward
	}

	for sl.level > 1 && sl.header.level[sl.level-1].forward == nil {
		sl.level--
	}
	sl.length--
}

func (sl *SkipList) Delete(score float64) bool {
	update := make([]*skiplistNode, SKIPLIST_MAXLEVEL)
	node := sl.header

	for i := sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && node.level[i].forward.score < score {
			node = node.level[i].forward
		}
		update[i] = node
	}

	node = node.level[0].forward
	if node != nil && score == node.score {
		sl.deleteNode(node, update)
		return true
	}

	return false
}

func (sl *SkipList) Find(score float64) bool {
	node := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && node.level[i].forward.score < score {
			node = node.level[i].forward
		}
	}

	node = node.level[0].forward
	if node != nil && score == node.score {
		fmt.Printf("Found %v\n", node.score)
		return true
	}

	fmt.Printf("Not Found %v\n", score)
	return false
}

func (sl *SkipList) Output() {
	var node *skiplistNode
	for i := 0; i < SKIPLIST_MAXLEVEL; i++ {
		fmt.Printf("LEVEL[%v]: ", i)
		node = sl.header.level[i].forward
		for node != nil {
			fmt.Printf("%v -> ", node.score)
			node = node.level[i].forward
		}
		fmt.Println("NULL")
	}
}
