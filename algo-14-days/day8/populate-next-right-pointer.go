package main

import (
	"fmt"
)

type queue []*Node

func (q queue) Push(v *Node) queue {
	return append(q, v)
}

func (q queue) Pop() (queue, *Node) {
	return q[1:], q[0]
}

func (q queue) IsEmpty() bool {
	return len(q) == 0
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect_bfs(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := make(queue, 0)
	queue = queue.Push(root)
	var rightNode *Node = nil
	var cur *Node = nil
	for !queue.IsEmpty() {
		queue, cur = queue.Pop()
		cur.Next = rightNode
		if cur != root {
			rightNode = cur
		}

		if cur.Right != nil {
			queue = queue.Push(cur.Right)
			queue = queue.Push(cur.Left)
		}
	}

	return root
}

func connect(root *Node) *Node {
	head := root

	for root != nil {
		cur := root
		root = root.Left

		for cur != nil {
			if cur.Left != nil {
				cur.Left.Next = cur.Right
				if cur.Next != nil {
					cur.Right.Next = cur.Next.Left
				}
			} else {
				break
			}

			cur = cur.Next
		}
	}

	return head
}

func main() {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Left = &Node{Val: 6}
	root.Right.Right = &Node{Val: 7}

	fmt.Println(connect(root))
}
