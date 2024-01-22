package main

type AVLTree struct {
	Root *AVLTreeNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

type AVLTreeNode struct {
	Key    int
	Height int
	Left   *AVLTreeNode
	Right  *AVLTreeNode
}

func (t *AVLTree) Search(key int) *AVLTreeNode {
	if t.Root == nil {
		return nil
	}

	if key < t.Root.Key {
		newTree := &AVLTree{
			Root: t.Root.Left,
		}
		return newTree.Search(key)
	} else if key > t.Root.Key {
		newTree := &AVLTree{
			Root: t.Root.Right,
		}
		return newTree.Search(key)
	} else {
		return t.Root
	}
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
			t.Root.Left = newTree.Root

			if height(t.Root.Left)-height(t.Root.Right) >= 2 {
				if height(t.Root.Left.Left) >= height(t.Root.Left.Right) {
					t.Root = leftLeftRotation(t.Root)
				} else {
					t.Root = leftRightRotation(t.Root)
				}
			}
		} else {
			newTree := &AVLTree{
				Root: t.Root.Right,
			}
			newTree.Insert(key)
			t.Root.Right = newTree.Root

			if height(t.Root.Right)-height(t.Root.Left) >= 2 {
				if height(t.Root.Right.Right) >= height(t.Root.Right.Left) {
					t.Root = rightRightRotation(t.Root)
				} else {
					t.Root = RightLeftRotation(t.Root)
				}
			}
		}

		if height(t.Root.Left) > height(t.Root.Right) {
			t.Root.Height = t.Root.Left.Height + 1
		} else {
			t.Root.Height = t.Root.Right.Height + 1
		}
	}
}

func height(node *AVLTreeNode) int {
	if node == nil {
		return 0
	} else {
		return node.Height
	}
}

func leftLeftRotation(node *AVLTreeNode) *AVLTreeNode {
	tmpNode := node.Left
	node.Left = tmpNode.Right
	tmpNode.Right = node

	if height(node.Left) > height(node.Right) {
		node.Height = height(node.Left) + 1
	} else {
		node.Height = height(node.Right) + 1
	}

	if height(tmpNode.Left) > height(tmpNode.Right) {
		tmpNode.Height = height(tmpNode.Left) + 1
	} else {
		tmpNode.Height = height(tmpNode.Right) + 1
	}
	return tmpNode
}

func rightRightRotation(node *AVLTreeNode) *AVLTreeNode {
	tmpNode := node.Right
	node.Right = tmpNode.Left
	tmpNode.Left = node

	if height(node.Left) > height(node.Right) {
		node.Height = height(node.Left) + 1
	} else {
		node.Height = height(node.Right) + 1
	}

	if height(tmpNode.Left) > height(tmpNode.Right) {
		tmpNode.Height = height(tmpNode.Left) + 1
	} else {
		tmpNode.Height = height(tmpNode.Right) + 1
	}
	return tmpNode
}

func leftRightRotation(node *AVLTreeNode) *AVLTreeNode {
	tmpNode := node.Left
	node.Left = node.Left.Right
	tmpNode.Right = node.Left.Right
	node.Left.Left = tmpNode

	if height(tmpNode.Left) > height(tmpNode.Right) {
		tmpNode.Height = height(tmpNode.Left) + 1
	} else {
		tmpNode.Height = height(tmpNode.Right) + 1
	}

	if height(node.Left.Left) > height(node.Left.Right) {
		node.Left.Height = height(node.Left.Left) + 1
	} else {
		node.Left.Height = height(node.Left.Right) + 1
	}

	tmpNode = node.Left
	node.Left = tmpNode.Right
	tmpNode.Right = node

	if height(node.Left) > height(node.Right) {
		node.Height = height(node.Left) + 1
	} else {
		node.Height = height(node.Right) + 1
	}

	if height(tmpNode.Left) > height(tmpNode.Right) {
		tmpNode.Height = height(tmpNode.Left) + 1
	} else {
		tmpNode.Height = height(tmpNode.Right) + 1
	}
	return tmpNode
}

func RightLeftRotation(node *AVLTreeNode) *AVLTreeNode {
	tmpNode := node.Right
	node.Right = tmpNode.Left
	tmpNode.Left = node.Right.Right
	node.Right.Right = tmpNode

	if height(tmpNode.Left) > height(tmpNode.Right) {
		tmpNode.Height = height(tmpNode.Left) + 1
	} else {
		tmpNode.Height = height(tmpNode.Right) + 1
	}

	if height(node.Left.Left) > height(node.Left.Right) {
		node.Left.Height = height(node.Left.Left) + 1
	} else {
		node.Left.Height = height(node.Left.Right) + 1
	}

	tmpNode = node.Right
	node.Right = tmpNode.Left
	tmpNode.Left = node

	if height(node.Left) > height(node.Right) {
		node.Height = height(node.Left) + 1
	} else {
		node.Height = height(node.Right) + 1
	}

	if height(tmpNode.Left) > height(tmpNode.Right) {
		tmpNode.Height = height(tmpNode.Left) + 1
	} else {
		tmpNode.Height = height(tmpNode.Right) + 1
	}
	return tmpNode
}

func (t *AVLTree) Delete(key int) {
	if t.Root == nil {
		return
	}

	if key < t.Root.Key {
		newTree := &AVLTree{
			Root: t.Root.Left,
		}
		newTree.Delete(key)
		t.Root.Left = newTree.Root

		if height(t.Root.Right)-height(t.Root.Left) >= 2 {
			if height(t.Root.Right.Right) > height(t.Root.Right.Left) {
				t.Root = rightRightRotation(t.Root)
			} else {
				t.Root = RightLeftRotation(t.Root)
			}
		}
	} else if key > t.Root.Key {
		newTree := &AVLTree{
			Root: t.Root.Right,
		}
		newTree.Delete(key)
		t.Root.Right = newTree.Root

		if height(t.Root.Left)-height(t.Root.Right) >= 2 {
			if height(t.Root.Left.Left) > height(t.Root.Left.Right) {
				t.Root = leftLeftRotation(t.Root)
			} else {
				t.Root = leftRightRotation(t.Root)
			}
		}
	} else {
		if t.Root.Left == nil && t.Root.Right == nil {
			t.Root = nil
		} else {
			if height(t.Root.Left) > height(t.Root.Right) {
				newTree := &AVLTree{
					Root: t.Root.Left,
				}

				maxNode := newTree.getMaxNode()
				t.Root.Key = maxNode.Key
				newTree.Delete(maxNode.Key)
				t.Root.Left = newTree.Root
			} else {
				newTree := &AVLTree{
					Root: t.Root.Right,
				}

				minNode := newTree.getMinNode()
				t.Root.Key = minNode.Key
				newTree.Delete(minNode.Key)
				t.Root.Right = newTree.Root
			}
		}
	}
}

func (t *AVLTree) getMaxNode() *AVLTreeNode {
	if t.Root == nil {
		return nil
	}
	if t.Root.Right == nil {
		return t.Root
	} else {
		newTree := &AVLTree{
			Root: t.Root.Right,
		}
		return newTree.getMaxNode()
	}
}

func (t *AVLTree) getMinNode() *AVLTreeNode {
	if t.Root == nil {
		return nil
	}
	if t.Root.Left == nil {
		return t.Root
	} else {
		newTree := &AVLTree{
			Root: t.Root.Left,
		}
		return newTree.getMinNode()
	}
}
