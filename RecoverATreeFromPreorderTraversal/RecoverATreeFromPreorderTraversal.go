package leetcode

import (
	"strconv"
	"strings"
)

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(S string) *TreeNode {
	ss := strings.Split(S, "-")
	val, _ := strconv.Atoi(ss[0])
	root := &TreeNode{Val: val}
	stack := []*TreeNode{root}
	pos := 1
	for pos < len(ss) {
		level := 1
		for ss[pos] == "" {
			level++
			pos++
		}
		val, _ = strconv.Atoi(ss[pos])
		node := &TreeNode{Val: val}
		if level >= len(stack) {
			stack[len(stack)-1].Left = node
		} else {
			stack = stack[0:level]
			stack[len(stack)-1].Right = node
		}
		stack = append(stack, node)
		pos++
	}
	return root
}
