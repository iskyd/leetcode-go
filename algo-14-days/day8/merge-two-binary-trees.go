package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	ans := &TreeNode{Val: root1.Val + root2.Val}
	ans.Left = mergeTrees(root1.Left, root2.Left)
	ans.Right = mergeTrees(root1.Right, root2.Right)

	return ans
}

func main() {
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 2}
	root1.Left.Left = &TreeNode{Val: 5}

	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Right = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 7}

	fmt.Println(mergeTrees(root1, root2))
}
