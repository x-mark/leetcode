package leetcode

import (
	"strconv"
	"strings"
)

// func main() {
// 	str := "[1,2,3,null,null,4,5]"
// 	obj := Constructor()
// 	root := obj.deserialize(str)
// 	fmt.Println(obj.serialize(root))
// }

// TreeNode node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec
type Codec struct {
}

// Constructor
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	parent := []*TreeNode{root}
	var tree []string
	max := 0
	for i := 0; i < len(parent); i++ {
		if parent[i] == nil {
			tree = append(tree, "null")
		} else {
			tree = append(tree, strconv.Itoa(parent[i].Val))
			parent = append(parent, parent[i].Left)
			if parent[i].Left != nil {
				max = len(parent) - 1
			}
			parent = append(parent, parent[i].Right)
			if parent[i].Right != nil {
				max = len(parent) - 1
			}
		}
		if i == max {
			break
		}
	}

	return "[" + strings.Join(tree, ",") + "]"

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 2 {
		return nil
	}

	data = data[1 : len(data)-1]
	ss := strings.Split(data, ",")
	val, _ := strconv.Atoi(ss[0])
	root := &TreeNode{Val: val}

	ss = ss[1:]

	parent := []*TreeNode{root}
	for i := 0; i < len(parent); i++ {
		if len(ss) > 0 {
			if ss[0] != "null" {
				val, _ = strconv.Atoi(ss[0])
				parent[i].Left = &TreeNode{Val: val}
				parent = append(parent, parent[i].Left)
			}
			ss = ss[1:]
		}

		if len(ss) > 0 {
			if ss[0] != "null" {
				val, _ = strconv.Atoi(ss[0])
				parent[i].Right = &TreeNode{Val: val}
				parent = append(parent, parent[i].Right)
			}
			ss = ss[1:]
		}

	}
	return root
}
