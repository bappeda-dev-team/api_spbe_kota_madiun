package helper

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	KodeReferensi string
	Children      []*TreeNode
}

func NewTreeNode(kodeReferensi string) *TreeNode {
	return &TreeNode{
		KodeReferensi: kodeReferensi,
		Children:      []*TreeNode{},
	}
}

func (node *TreeNode) AddChild(child *TreeNode) {
	node.Children = append(node.Children, child)
}

func PrintTree(node *TreeNode, level int) {
	if node == nil {
		return
	}

	fmt.Println(strings.Repeat("  ", level) + node.KodeReferensi)

	for _, child := range node.Children {
		PrintTree(child, level+1)
	}
}
