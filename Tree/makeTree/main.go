package main

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func (tree *TreeNode) getLNode() *TreeNode {
	if tree.left != nil {
		return tree
	}
	return tree.left
}
func (tree *TreeNode) getRNode() *TreeNode {
	if tree.right != nil {
		return tree
	}
	return tree.right
}

//创建一个深度为h的二叉树
func makeTreeByHeight(h int) *TreeNode {
	a := newTreeNode(1)
	fmt.Println(a)
	a.getLNode().left = newTreeNode(2)
	a.getLNode().right = newTreeNode(3)
	return a
}
func main() {
	s := make([]int, 0)
	s = append(s, 10)
	fmt.Println(s)
	fmt.Println(s)
	//t := makeTreeByHeight(3)
	//fmt.Printf("%+v", t)
}
