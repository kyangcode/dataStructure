package main

import (
	"math"
	"sort"
)

type Data interface {
	Less(data Data) bool
}

type List []Data

func (l List) BinarySearch(data Data) (int, bool) {
	r := sort.Search(len(l), func(i int) bool {
		return data.Less(l[i])
	})

	if r > 0 && !l[r-1].Less(data) {
		return r - 1, true
	}
	return r, false
}

type Node struct {
	List     List
	Children []*Node
	Parent   *Node
}

type BTree struct {
	Root *Node
	M    int // 树的阶
}

func NewBTree(m int) *BTree {
	return &BTree{
		M: m,
	}
}

func (t *BTree) Insert(data Data) {
	if t.Root == nil {
		t.Root = &Node{}
	}

	leaf, index := t.GetLeafNodeForInsert(data)
	leaf.List = append(leaf.List[:index], append(List{data}, leaf.List[index:]...)...)

	// 分裂
	t.Split(leaf)
}

func (t *BTree) Search(data Data) (*Node, int, Data) {
	if t.Root == nil {
		return nil, -1, nil
	}

	index, isExist := t.Root.List.BinarySearch(data)
	if isExist {
		return t.Root, index, t.Root.List[index]
	}
	if len(t.Root.Children) == 0 {
		return nil, -1, nil
	}

	child := t.Root.Children[index]
	childBTree := &BTree{M: t.M, Root: child}
	return childBTree.Search(data)
}

func (t *BTree) Delete(data Data) {
	node, index, _ := t.Search(data)
	if node == nil {
		return
	}

	var reBalanceNode *Node
	min := int(math.Ceil(float64(t.M)/2)) - 1
	if len(node.Children) > 0 {
		left := node.Children[index]
		right := node.Children[index+1]

		var ele Data
		if len(right.List) > min {
			ele = right.List[0]
			right.List = right.List[1:]

			reBalanceNode = right
		} else {
			ele = left.List[len(left.List)-1]
			left.List = left.List[0 : len(left.List)-1]

			reBalanceNode = left
		}
		node.List[index] = ele
	} else {
		if index+1 == len(node.List) {
			node.List = node.List[0 : index-1]
		} else {
			node.List = append(node.List[0:index], node.List[index+1:]...)
		}
		reBalanceNode = node
	}

	// 调整树使满足条件
	t.ReBalance(reBalanceNode)
}

func (t *BTree) ReBalance(node *Node) {
	if node.Parent == nil && len(node.Children) == 0 {
		return
	}

	min := int(math.Ceil(float64(t.M)/2)) - 1

	var childIndex int
	var leftBrother, rightBrother *Node
	for index, child := range node.Parent.Children {
		if child == node {
			if index-1 >= 0 {
				leftBrother = node.Parent.Children[index-1]
			}
			if index+1 <= len(node.Parent.Children) {
				rightBrother = node.Parent.Children[index+1]
			}
			childIndex = index
		}
	}

	if leftBrother != nil && len(leftBrother.List) > min {
		node.List = append(List{node.Parent.List[childIndex-1]}, node.List...)
		node.Parent.List[childIndex-1] = leftBrother.List[len(leftBrother.List)-1]
		leftBrother.List = leftBrother.List[0 : len(leftBrother.List)-1]
	} else if rightBrother != nil && len(rightBrother.List) > min {
		node.List = append(node.List, node.Parent.List[childIndex])
		node.Parent.List[childIndex] = rightBrother.List[0]
		rightBrother.List = rightBrother.List[1:]
	} else if leftBrother != nil {

	} else if rightBrother != nil {

	}
}

func (t *BTree) Split(n *Node) {
	if len(n.List) <= t.M-1 {
		return
	}

	k := int(math.Ceil(float64(len(n.List))/2)) - 1
	mid := n.List[k]

	rightNode := &Node{
		List:   n.List[k+1:],
		Parent: n.Parent,
	}
	n.List = n.List[:k]

	if len(n.Children) != 0 {
		rightNode.Children = n.Children[k+1:]
		for _, child := range rightNode.Children {
			child.Parent = rightNode
		}
		n.Children = n.Children[:k+1]
	}

	if n.Parent != nil {
		index, _ := n.Parent.List.BinarySearch(mid)
		if index >= len(n.Parent.List) {
			n.Parent.List = append(n.Parent.List, mid)
		} else {
			n.Parent.List = append(n.Parent.List[:index], append(List{mid}, n.Parent.List[index:]...)...)
		}

		if index+1 >= len(n.Parent.Children) {
			n.Parent.Children = append(n.Parent.Children, rightNode)
		} else {
			n.Parent.Children = append(n.Parent.Children[:index], append([]*Node{rightNode}, n.Parent.Children[index+1:]...)...)
		}
		t.Split(n.Parent)
	} else {
		n.Parent = &Node{
			List:     []Data{mid},
			Children: []*Node{n, rightNode},
		}
		rightNode.Parent = n.Parent
		t.Root = n.Parent
	}
}

// GetLeafNodeForInsert 获取待插入的叶子结点
func (t *BTree) GetLeafNodeForInsert(data Data) (*Node, int) {
	index, isExist := t.Root.List.BinarySearch(data)
	if isExist {
		return nil, 0
	}
	if len(t.Root.Children) == 0 {
		return t.Root, index
	}

	bTree := &BTree{
		M:    t.M,
		Root: t.Root.Children[index],
	}
	return bTree.GetLeafNodeForInsert(data)
}
