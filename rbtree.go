package main

// RBTree 红黑树的定义
// 1. 结点要么是红色要么是黑色
// 2. 根结点是黑色
// 3. 红色结点的孩子必须是黑色
// 4. 任意一个结点到各个叶结点之间黑色结点数相同
type RBTree struct {
	Root *RBTreeNode
}

func NewRBTree() *RBTree {
	return &RBTree{}
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

func (t *RBTree) llRotate(node *RBTreeNode) {
	parent := node.Parent

	p := node.Right
	node.Right = p.Left
	if p.Left != nil {
		p.Left.Parent = node
	}

	p.Left = node
	node.Parent = p

	if parent == nil {
		t.Root = p
	} else if parent.Left == node {
		parent.Left = p
	} else {
		parent.Right = p
	}
}

func (t *RBTree) rrRotate(node *RBTreeNode) {
	parent := node.Parent

	p := node.Left
	node.Left = p.Right
	if p.Right != nil {
		p.Right.Parent = node
	}

	p.Right = node
	node.Parent = p

	if parent == nil {
		t.Root = p
	} else if parent.Left == node {
		parent.Left = p
	} else {
		parent.Right = p
	}
}

func (t *RBTree) fixup(node *RBTreeNode) {
	parent := node.Parent
	if parent.Color == black {
		return
	}

	grandParent := parent.Parent
	if grandParent.Right == parent {
		if grandParent.Left != nil && grandParent.Left.Color == red {
			parent.Color = black
			grandParent.Left.Color = black
			grandParent.Color = red
			t.fixup(grandParent)
			return
		}

		if node == parent.Right {
			t.llRotate(grandParent)
			parent.Color = black
			grandParent.Color = red
		} else {
			t.rrRotate(parent)
			t.llRotate(grandParent)

			parent.Color = black
			grandParent.Color = red
		}
	} else {
		if grandParent.Right != nil && grandParent.Right.Color == red {
			parent.Color = black
			grandParent.Right.Color = black
			grandParent.Color = red
			t.fixup(grandParent)
			return
		}

		if node == parent.Left {
			t.rrRotate(grandParent)
			parent.Color = black
			grandParent.Color = red
		} else {
			t.llRotate(parent)
			t.rrRotate(grandParent)

			parent.Color = black
			grandParent.Color = red
		}
	}

}

func (t *RBTree) Delete(key int) {

}

func (t *RBTree) Search(key int) *RBTreeNode {
	return nil
}
