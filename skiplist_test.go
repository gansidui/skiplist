package skiplist

import (
	"fmt"
	"math/rand"
	"testing"
)

type Int int

func (i Int) Less(other interface{}) bool {
	return i < other.(Int)
}

func TestInt(t *testing.T) {
	sl := New()
	if sl.Len() != 0 || sl.Front() != nil && sl.Back() != nil {
		t.Fatal()
	}

	testData := []Int{Int(1), Int(2), Int(3)}

	sl.Insert(testData[0])
	if sl.Len() != 1 || sl.Front().Value.(Int) != testData[0] || sl.Back().Value.(Int) != testData[0] {
		t.Fatal()
	}

	sl.Insert(testData[2])
	if sl.Len() != 2 || sl.Front().Value.(Int) != testData[0] || sl.Back().Value.(Int) != testData[2] {
		t.Fatal()
	}

	sl.Insert(testData[1])
	if sl.Len() != 3 || sl.Front().Value.(Int) != testData[0] || sl.Back().Value.(Int) != testData[2] {
		t.Fatal()
	}

	sl.Insert(Int(-999))
	sl.Insert(Int(-888))
	sl.Insert(Int(888))
	sl.Insert(Int(999))
	sl.Insert(Int(1000))

	expect := []Int{Int(-999), Int(-888), Int(1), Int(2), Int(3), Int(888), Int(999), Int(1000)}
	ret := make([]Int, 0)

	for e := sl.Front(); e != nil; e = e.Next() {
		ret = append(ret, e.Value.(Int))
	}
	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[i] {
			t.Fatal()
		}
	}

	e := sl.Find(Int(2))
	if e == nil || e.Value.(Int) != 2 {
		t.Fatal()
	}

	ret = make([]Int, 0)
	for ; e != nil; e = e.Next() {
		ret = append(ret, e.Value.(Int))
	}
	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[i+3] {
			t.Fatal()
		}
	}

	sl.Remove(sl.Find(Int(2)))
	sl.Delete(Int(888))
	sl.Delete(Int(1000))

	expect = []Int{Int(-999), Int(-888), Int(1), Int(3), Int(999)}
	ret = make([]Int, 0)

	for e := sl.Back(); e != nil; e = e.Prev() {
		ret = append(ret, e.Value.(Int))
	}

	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[len(ret)-i-1] {
			t.Fatal()
		}
	}

	if sl.Front().Value.(Int) != -999 {
		t.Fatal()
	}

	sl.Remove(sl.Front())
	if sl.Front().Value.(Int) != -888 || sl.Back().Value.(Int) != 999 {
		t.Fatal()
	}

	sl.Remove(sl.Back())
	if sl.Front().Value.(Int) != -888 || sl.Back().Value.(Int) != 3 {
		t.Fatal()
	}

	if e = sl.Insert(Int(2)); e.Value.(Int) != 2 {
		t.Fatal()
	}
	sl.Delete(Int(-888))

	if r := sl.Delete(Int(123)); r != nil {
		t.Fatal()
	}

	if sl.Len() != 3 {
		t.Fatal()
	}

	sl.Insert(Int(2))
	sl.Insert(Int(2))
	sl.Insert(Int(1))

	if e = sl.Find(Int(2)); e == nil {
		t.Fatal()
	}

	expect = []Int{Int(2), Int(2), Int(2), Int(3)}
	ret = make([]Int, 0)
	for ; e != nil; e = e.Next() {
		ret = append(ret, e.Value.(Int))
	}
	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[i] {
			t.Fatal()
		}
	}

	sl2 := sl.Init()
	if sl.Len() != 0 || sl.Front() != nil || sl.Back() != nil ||
		sl2.Len() != 0 || sl2.Front() != nil || sl2.Back() != nil {
		t.Fatal()
	}

	// for i := 0; i < 100; i++ {
	// 	sl.Insert(Int(rand.Intn(200)))
	// }
	// output(sl)
}

func TestRank(t *testing.T) {
	sl := New()
	// 1 2 2 2 3 3 4 5 6 6 7 8
	testData := []Int{Int(5), Int(2), Int(3), Int(1), Int(4), Int(2), Int(3),
		Int(6), Int(2), Int(7), Int(8), Int(6)}

	for i := 0; i < len(testData); i++ {
		sl.Insert(testData[i])
	}

	for i := 12; i < 100; i++ {
		sl.Insert(Int(i))
	}

	expect := map[Int]int{Int(1): 1, Int(2): 2, Int(3): 5, Int(4): 7,
		Int(5): 8, Int(6): 9, Int(7): 11, Int(8): 12}

	for k, v := range expect {
		if sl.GetRank(k) != v {
			t.Fatal()
		}
	}

	if sl.GetRank(Int(99)) != 100 || sl.GetRank(Int(92)) != 93 || sl.GetRank(Int(0)) != 0 ||
		sl.GetRank(Int(12)) != 13 || sl.GetRank(Int(13)) != 14 || sl.GetRank(Int(10)) != 0 {
		t.Fatal()
	}

	// 1 1 1 1 2 2 2 2 3 3 3 3
	sl = sl.Init()
	for i := 1; i <= 3; i++ {
		for j := 0; j < 4; j++ {
			sl.Insert(Int(i))
		}
	}
	fmt.Println(sl.GetRank(Int(1)))
	fmt.Println(sl.GetRank(Int(2)))
	fmt.Println(sl.GetRank(Int(3)))

	fmt.Println("=============")
	for e := sl.GetElementByRank(13); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// fmt.Println(sl.GetElementByRank(1).Value)
	// fmt.Println(sl.GetElementByRank(2).Value)
	// fmt.Println(sl.GetElementByRank(6).Value)
	// fmt.Println(sl.GetElementByRank(11).Value)

}

func BenchmarkIntInsertOrder(b *testing.B) {
	b.StopTimer()
	sl := New()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Insert(Int(i))
	}
}

func BenchmarkIntInsertRandom(b *testing.B) {
	b.StopTimer()
	sl := New()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Insert(Int(rand.Int()))
	}
}

func BenchmarkIntDeleteOrder(b *testing.B) {
	b.StopTimer()
	sl := New()
	for i := 0; i < 1000000; i++ {
		sl.Insert(Int(i))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Delete(Int(i))
	}
}

func BenchmarkIntDeleteRandome(b *testing.B) {
	b.StopTimer()
	sl := New()
	for i := 0; i < 1000000; i++ {
		sl.Insert(Int(rand.Int()))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Delete(Int(rand.Int()))
	}
}

func BenchmarkIntFindOrder(b *testing.B) {
	b.StopTimer()
	sl := New()
	for i := 0; i < 1000000; i++ {
		sl.Insert(Int(i))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Find(Int(i))
	}
}

func BenchmarkIntFindRandom(b *testing.B) {
	b.StopTimer()
	sl := New()
	for i := 0; i < 1000000; i++ {
		sl.Insert(Int(rand.Int()))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Find(Int(rand.Int()))
	}
}

func output(sl *SkipList) {
	var x *Element
	for i := 0; i < SKIPLIST_MAXLEVEL; i++ {
		fmt.Printf("LEVEL[%v]: ", i)
		x = sl.header.level[i].forward
		for x != nil {
			fmt.Printf("%v -> ", x.Value)
			x = x.level[i].forward
		}
		fmt.Println("NIL")
	}
}
