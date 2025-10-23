package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode() *ListNode {
	return &ListNode{
		Next: nil,
	}
}

func (l *ListNode) Push(numbers []int) {
	if len(numbers) < 1 {
		return
	}

	l.Val = numbers[0]
	for i := 1; i < len(numbers); i++ {
		l.Add(numbers[i])
	}
}

func (l *ListNode) Add(value int) {
	root := l

	for {
		if root.Next == nil {
			root.Next = &ListNode{Val: value, Next: nil}
			break
		}
		root = root.Next
	}
}

func (l *ListNode) List() {
	fmt.Printf("nodes: %v\n", l.ToArray())
}

func (l *ListNode) ToArray() []int {
	listNode := []int{}
	root := l

	for root != nil {
		listNode = append(listNode, root.Val)
		root = root.Next
	}
	return listNode
}

// func (l *ListNode) Invert() {
// 	var previous *ListNode
// 	current := l

// 	for current != nil {
// 		next := current.Next
// 		current.Next = previous
// 		previous = current
// 		current = next
// 	}
// 	return previous
// }

func (l *ListNode) ToInverse() *ListNode {
	var previous *ListNode
	current := l

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}
