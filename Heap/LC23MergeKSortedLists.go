// link: https://leetcode.com/problems/merge-k-sorted-lists/
package Heap

import (
	"container/heap"
	. "github.com/CSStudySession/AlgoInGo/LinkedLists"
)

type NodeHeap []*ListNode

func (h NodeHeap) Len() int {return len(h)}
func (h NodeHeap) Less(i, j int) bool {return h[i].Val < h[j].Val}
func (h NodeHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *NodeHeap) Pop() interface{} {
	item := (*h)[len(*h) - 1]
	*h = (*h)[: len(*h) - 1]
	return item
}

func (h *NodeHeap) Push(item interface{}) {
	*h = append(*h, item.(*ListNode))
}


func mergeKLists(lists []*ListNode) *ListNode {
	mergeHeap := &NodeHeap{}
	heap.Init(mergeHeap)

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(mergeHeap, lists[i])
		}
	}

	merged := &ListNode{}
	head := merged
	if mergeHeap.Len() > 0 {
		nxt := heap.Pop(mergeHeap).(*ListNode)
		if nxt.Next != nil {
			heap.Push(mergeHeap, nxt.Next)
		}
		head.Next = nxt
		head = head.Next
	}
	return merged.Next
}
