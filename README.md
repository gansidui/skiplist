skiplist
===============

reference from redis [zskiplist](https://github.com/antirez/redis)


Usage
===============

~~~Go

package main

import (
	"fmt"
	"github.com/gansidui/skiplist"
)

type Node struct {
	score float64
	id    int
}

func (this *Node) Less(other interface{}) bool {
	return this.score < other.(*Node).score
}

func NewNode(score float64, id int) *Node {
	return &Node{
		score: score,
		id:    id,
	}
}

func main() {
	data := make([]*Node, 8)
	data[0] = NewNode(5.5, 1)
	data[1] = NewNode(2.2, 2)
	data[2] = NewNode(2.2, 3)
	data[3] = NewNode(2.2, 4)
	data[4] = NewNode(1.1, 5)
	data[5] = NewNode(4.4, 6)
	data[6] = NewNode(3.3, 7)
	data[7] = NewNode(6.6, 8)

	sl := skiplist.New()
	for i := 0; i < len(data); i++ {
		sl.Insert(data[i])
	}

	node := NewNode(2.2, 999)

	// According to the score to find
	for e := sl.Find(node); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*Node).id, "-->", e.Value.(*Node).score)
	}
	fmt.Println()

	for e := sl.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*Node).id, "-->", e.Value.(*Node).score)
	}
	fmt.Println()

	for e := sl.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value.(*Node).id, "-->", e.Value.(*Node).score)
	}
}

/* output:

4 --> 2.2
3 --> 2.2
2 --> 2.2
7 --> 3.3
6 --> 4.4
1 --> 5.5
8 --> 6.6

5 --> 1.1
4 --> 2.2
3 --> 2.2
2 --> 2.2
7 --> 3.3
6 --> 4.4
1 --> 5.5
8 --> 6.6

8 --> 6.6
1 --> 5.5
6 --> 4.4
7 --> 3.3
2 --> 2.2
3 --> 2.2
4 --> 2.2
5 --> 1.1

*/


~~~


License
===============

MIT