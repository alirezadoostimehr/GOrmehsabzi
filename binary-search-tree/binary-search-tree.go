package BST

import "fmt"

type BstNode struct {
	Val   int
	Left  *BstNode
	Right *BstNode
}

type BinarySearchTree struct {
	root *BstNode
}

func insert(treeNode *BstNode, insertionValue int) *BstNode {
	if treeNode == nil {
		return &BstNode{Val: insertionValue}
	}
	if insertionValue < treeNode.Val {
		treeNode.Left = insert(treeNode.Left, insertionValue)
	} else if insertionValue > treeNode.Val {
		treeNode.Right = insert(treeNode.Right, insertionValue)
	}
	return treeNode
}

func (BT *BinarySearchTree) Insert(insetionValue int) {
	BT.root = insert(BT.root, insetionValue)
}

func search(treeNode *BstNode, searchValue int) *BstNode {
	if treeNode == nil || treeNode.Val == searchValue {
		return treeNode
	}
	if searchValue < treeNode.Val {
		return search(treeNode.Left, searchValue)
	}
	return search(treeNode.Right, searchValue)
}

func (BT *BinarySearchTree) Search(searchValue int) bool {
	return search(BT.root, searchValue) != nil
}

func delete(treeNode *BstNode, deletionValue int) *BstNode {
	if deletionValue < treeNode.Val {
		treeNode.Left = delete(treeNode.Left, deletionValue)
	} else if treeNode.Val < deletionValue {
		treeNode.Right = delete(treeNode.Right, deletionValue)
	} else {
		if treeNode.Left == nil && treeNode.Right == nil {
			return nil
		} else if treeNode.Left == nil {
			return treeNode.Right
		} else if treeNode.Right == nil {
			return treeNode.Left
		} else {
			parNode := treeNode
			tmpNode := treeNode.Right
			for tmpNode.Left != nil {
				parNode = tmpNode
				tmpNode = tmpNode.Left
			}
			if parNode.Val < tmpNode.Val {
				parNode.Right = tmpNode.Right
			} else {
				parNode.Left = tmpNode.Right
			}
			tmpNode.Left = treeNode.Left
			tmpNode.Right = treeNode.Right
			return tmpNode
		}
	}
	return treeNode
}

func (BT *BinarySearchTree) Delete(deletionValue int) {
	if !BT.Search(deletionValue) {
		return
	}
	BT.root = delete(BT.root, deletionValue)
}

func inOrder(treeNode *BstNode) string {
	if treeNode == nil {
		return ""
	}
	return inOrder(treeNode.Left) +
		fmt.Sprint(treeNode.Val, " ") +
		inOrder(treeNode.Right)
}

func (BT *BinarySearchTree) InOrder() string {
	return inOrder(BT.root)
}

func preOrder(treeNode *BstNode) string {
	if treeNode == nil {
		return ""
	}
	return fmt.Sprint(treeNode.Val, " ") +
		preOrder(treeNode.Left) +
		preOrder(treeNode.Right)
}

func (BT *BinarySearchTree) PreOrder() string {
	return preOrder(BT.root)
}

func postOrder(treeNode *BstNode) string {
	if treeNode == nil {
		return ""
	}
	return postOrder(treeNode.Left) +
		postOrder(treeNode.Right) +
		fmt.Sprint(treeNode.Val, " ")
}

func (BT *BinarySearchTree) PostOrder() string {
	return postOrder(BT.root)
}
