package main

type AVLTree struct {
	Root *AVLTreeNode
}

type AVLTreeNode struct {
	Key    int
	Height int
	Left   *AVLTreeNode
	Right  *AVLTreeNode
}

func (t *AVLTree) Search(key int) *AVLTreeNode {
	return nil
}

func (t *AVLTree) Insert(key int) {
	if t.Root == nil {
		t.Root = &AVLTreeNode{
			Key:    key,
			Height: 1,
		}
	} else {
		if key <= t.Root.Key {
			newTree := &AVLTree{
				Root: t.Root.Left,
			}
			newTree.Insert(key)
		} else {
			newTree := &AVLTree{
				Root: t.Root.Right,
			}
			newTree.Insert(key)
		}
	}
}

func leftLeftRotation(node *AVLTreeNode) *AVLTreeNode {
	tmpNode := node.Left
	node.Left = tmpNode.Right
	tmpNode.Right = node
	return tmpNode
}

func rightRightRotation(node *AVLTreeNode) *AVLTreeNode {
	tmpNode := node.Right
	node.Right = tmpNode.Left
	tmpNode.Left = node
	return tmpNode
}

func leftRightRotation(node *AVLTreeNode) *AVLTreeNode {
	tmpNode := node.Left
	node.Left = node.Left.Right
	tmpNode.Right = node.Left.Right
	node.Left.Left = tmpNode

	tmpNode = node.Left
	node.Left = tmpNode.Right
	tmpNode.Right = node
	return tmpNode
}

func RightLeftRotation(node *AVLTreeNode) *AVLTreeNode {
	tmpNode := node.Right
	node.Right = tmpNode.Left
	tmpNode.Left = node.Right.Right
	node.Right.Right = tmpNode

	tmpNode = node.Right
	node.Right = tmpNode.Left
	tmpNode.Left = node
	return tmpNode
}

func (t *AVLTree) Delete(key int) {

}
