package LinkedLists

import (
	"fmt"
)

const N int = 100010

var head, idx int
var e, ne = make([]int, N), make([]int, N)

// init
func init() {
	head = -1
	idx = 0
}

// put item x in front of head
func addToHead(x int) {
	e[idx] = x
	ne[idx] = head
	head = idx
	idx++
}

// put x in next position of k
func add(k int, x int) {
	e[idx] = x
	ne[idx] = ne[k]
	ne[k] = idx
	idx++
}

// delete item in next position of k
func remove(k int) {
	ne[k] = ne[ne[k]]
}

func main() {
	var m int
	fmt.Scanf("%d\n", &m)

	for ; m > 0; m-- {
		var k, x int
		var op byte
		fmt.Scanf("%c", &op)

		if op == 'H' {
			fmt.Scanf("%d\n", &x)
			addToHead(x)
		} else if op == 'I' {
			fmt.Scanf("%d%d\n", &k, &x)
			add(k - 1, x)
		} else {
			fmt.Scanf("%d\n", &k)
			if k == 0 {
				head = ne[head]
			} else {
				remove(k - 1)
			}
		}
	}

	for i := head; i != -1; i = ne[i] {
		fmt.Printf("%d", e[i])
		fmt.Print(" ")
	}

}