package main

// RBTree 红黑树的定义
// 1. 结点要么是红色要么是黑色
// 2. 根结点是黑色
// 3. 红色结点的孩子必须是黑色
// 4. 任意一个结点到各个叶结点之间黑色结点数相同
type RBTree struct {
	Root *RBTreeNode
}

const (
	black = 1
	red   = 2
)

type RBTreeNode struct {
	Key    int
	Color  int
	Left   *RBTreeNode
	Right  *RBTreeNode
	Parent *RBTreeNode
}

func (t *RBTree) Insert(key int) *RBTreeNode {
	if t.Root == nil {
		t.Root = &RBTreeNode{
			Key:   key,
			Color: black,
		}
		return t.Root
	}

	parent := t.Root
	for {
		if key < parent.Key {
			if parent.Left == nil {
				break
			} else {
				parent = parent.Left
			}
		} else if key > parent.Key {
			if parent.Right == nil {
				break
			} else {
				parent = parent.Right
			}
		} else {
			return parent
		}
	}

	var newNode *RBTreeNode
	if key < parent.Key {
		newNode = &RBTreeNode{
			Key:    key,
			Parent: parent,
			Color:  red,
		}
		parent.Left = newNode
	} else {
		newNode = &RBTreeNode{
			Key:    key,
			Parent: parent,
			Color:  red,
		}
		parent.Right = newNode
	}

	t.fixup(newNode)
	return newNode
}

func llRotate(node *RBTreeNode) *RBTreeNode {
	p := node.Right
	node.Right = p.Left
	p.Left.Parent = node

	p.Left = node
	node.Parent = p
	return p
}

func rrRotate(node *RBTreeNode) *RBTreeNode {
	p := node.Left
	node.Left = p.Right
	p.Right.Parent = node

	p.Right = node
	node.Parent = p
	return p
}

func (t *RBTree) fixup(node *RBTreeNode) {
	parent := node.Parent
	if parent.Color == black {
		return
	}

	grandParent := parent.Parent
	if grandParent != nil && grandParent.Right.Color == red {
		parent.Color = black
		grandParent.Right.Color = black
		grandParent.Color = red
		t.fixup(grandParent)
		return
	}

	if node == parent.Left {
		if grandParent != nil {
			if grandParent.Left == parent {
				grandParent.Left = rrRotate(parent)
			} else {
				grandParent.Right = rrRotate(parent)
			}
		} else {
			t.Root = rrRotate(parent)
		}
	} else {

		// ll //rr
	}
}

func (t *RBTree) Delete(key int) {

}

func (t *RBTree) Search(key int) *RBTreeNode {
	return nil
}
