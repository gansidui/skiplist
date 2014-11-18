package skiplist

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test(t *testing.T) {
	sl := New()
	count := 20

	fmt.Println("====== Init Skip List =======")
	for i := 0; i < count; i++ {
		sl.Insert(float64(i))
	}

	fmt.Println("====== Print Skip List ======")

	sl.Output()

	fmt.Println("======== Search Skip List ======")
	for i := 0; i < count; i++ {
		sl.Find(float64(rand.Intn(count + 10)))
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("===== Delete Skip List ======")
	for i := 0; i < count+10; i += 2 {
		if sl.Delete(float64(i)) {
			fmt.Printf("Delete[%v]: SUCCESS\n", i)
		} else {
			fmt.Printf("Delete[%v]: NOT FOUND\n", i)
		}
	}

	sl.Output()

}
