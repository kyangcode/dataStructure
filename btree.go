package main

import (
	"math"
	"sort"
)

// B树的定义
// 对于m阶的B树，需要满足以下条件：
// 1. 叶子节点在同一层
// 2. 除叶子节点，其他节点最多有m-1个关键字
// 3. 除叶子节点，其他节点最少有ceil(m/2)-1个关键字，根节点最少有1个关键字
// 4. 除叶子节点，若有n个关键字，则有n+1棵子树
// 5. 节点中关键字非降序排列
// 6. 假设y为某个节点关键字的左子树，则以y为根的子树上关键字都不大于当前节点关键字。
// 7. 假设z为某个节点关键字的右子树，则以z为根的子树上关键字都不小于当前节点关键字。

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

// Insert 插入
// 假设最大关键字数为max
// 1. 先查找一个合适的叶子节点，并将关 键字插入
// 2. 若叶子节点的关键字数大于max，则需要分裂
// 3. 按照中间的位置，将当前节点分裂成两个节点
// 4. 将关键关键字插入到父亲节点
// 5. 父节点递归调用分裂的函数
func (t *BTree) Insert(data Data) {
	if t.Root == nil {
		t.Root = &Node{}
	}

	leaf, index := t.GetLeafNodeForInsert(data)
	leaf.List = append(leaf.List[:index], append(List{data}, leaf.List[index:]...)...)

	// 分裂
	t.Split(leaf)
}

// Search B树搜索
// 1. 二分查找当前节点，找到直接返回
// 2. 递归查找子树即可
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

// Delete B树根据关键字删除
// 1. 对于当前节点，若key存在。
// 2. 如果是叶子节点直接删除即可。
// 3. 如果是内部节点，则判断左子树关键字个数是否满足最小min，其中min为ceil(M/2)，若满足，则递归删除最右关键字k1，并用k1替换当前关键字。
// 4. 否则，判断右子树关键字个数是否满足最小min，若满足，则递归安徽念书最左关键字k2, 并用k2替换当前关键字。
// 5. 否则，合并左右子树和当前关键字，并调用递归删除。
// 6. 若当前节点不存在要删除的关键字，且关键字一定在c子树。
// 7. 若c子树关键字大于等于min，则递归调用删除。
// 8. 否则，看c的兄弟节点是否满足大于等于min，如果是，则将父节点关键字移动至当前节点，并且把兄弟节点移动至父亲节点，并把子树移动至当前节点。
// 9. 否则合并任意一个兄弟节点和父节点关键字。
func (t *BTree) Delete(data Data) {
	min := int(math.Ceil(float64(t.M) / 2))

	index, isExist := t.Root.List.BinarySearch(data)
	if isExist {
		if len(t.Root.Children) == 0 { // 叶子节点，直接删除对应元素
			if index == len(t.Root.List)-1 {
				t.Root.List = t.Root.List[0:index]
			} else {
				t.Root.List = append(t.Root.List[0:index], t.Root.List[index+1:]...)
			}
		} else { // 内部节点
			left := t.Root.Children[index]
			right := t.Root.Children[index+1]
			if len(left.List) >= min {
				newData := left.List[len(left.List)-1]
				leftTree := &BTree{
					M:    t.M,
					Root: left,
				}
				leftTree.Delete(newData)
				t.Root.List[index] = newData
			} else if len(right.List) >= min {
				newData := right.List[0]
				rightTree := &BTree{
					M:    t.M,
					Root: right,
				}
				rightTree.Delete(newData)
				t.Root.List[index] = newData
			} else {
				left.List = append(left.List, append(List{data}, right.List...)...)
				left.Children = append(left.Children, right.Children...)
				for _, child := range right.Children {
					child.Parent = left
				}

				if index == len(t.Root.List)-1 {
					t.Root.List = t.Root.List[0:index]
				} else {
					t.Root.List = append(t.Root.List[0:index], t.Root.List[index+1:]...)
				}
				if index+1 == len(t.Root.Children)-1 {
					t.Root.Children = t.Root.Children[0 : index+1]
				} else {
					t.Root.Children = append(t.Root.Children[0:index+1], t.Root.Children[index+1:]...)
				}

				if len(t.Root.List) == 0 {
					t.Root.List = left.List
					t.Root.Children = left.Children
					for _, child := range left.Children {
						child.Parent = t.Root
					}
				}

				newTree := &BTree{
					M:    t.M,
					Root: left,
				}
				newTree.Delete(data)
			}
		}
	} else {
		node := t.Root.Children[index]
		if len(node.List) >= min {
			newTree := &BTree{
				M:    t.M,
				Root: node,
			}
			newTree.Delete(data)
		} else {
			var lefBrother, rightBrother *Node
			if index > 0 {
				lefBrother = t.Root.Children[index-1]
			}
			if index < len(t.Root.Children)-1 {
				rightBrother = t.Root.Children[index+1]
			}

			if lefBrother != nil && len(lefBrother.List) >= min {
				node.List = append(List{t.Root.List[index-1]}, node.List...)
				t.Root.List[index-1] = lefBrother.List[len(lefBrother.List)-1]

				if len(lefBrother.Children) != 0 {
					node.Children = append([]*Node{lefBrother.Children[len(lefBrother.Children)-1]}, node.Children...)
					node.Children[0].Parent = node

					lefBrother.Children = lefBrother.Children[0 : len(lefBrother.Children)-1]
				}
			} else if rightBrother != nil && len(rightBrother.List) >= min {
				node.List = append(node.List, t.Root.List[index])
				t.Root.List[index] = rightBrother.List[0]

				if len(rightBrother.Children) != 0 {
					node.Children = append(node.Children, rightBrother.Children[0])
					node.Children[len(node.Children)-1] = node

					rightBrother.Children = rightBrother.Children[1:]
				}
			} else {
				if lefBrother != nil {
					node.List = append(lefBrother.List, append(List{t.Root.List[index-1]}, node.List...)...)
					t.Root.List = append(t.Root.List[0:index-1], t.Root.List[index:]...)

					node.Children = append(lefBrother.Children, node.Children...)
					for _, child := range lefBrother.Children {
						child.Parent = node
					}
					t.Root.Children = append(t.Root.Children[0:index-1], t.Root.Children[index:]...)
				} else if rightBrother != nil {
					node.List = append(node.List, append(List{t.Root.List[index]}, rightBrother.List...)...)
					if index == len(t.Root.List)-1 {
						t.Root.List = t.Root.List[0:index]
					} else {
						t.Root.List = append(t.Root.List[0:index], t.Root.List[index+1:]...)
					}

					node.Children = append(node.Children, rightBrother.Children...)
					for _, child := range rightBrother.Children {
						child.Parent = node
					}
					if index+1 == len(t.Root.Children) {
						t.Root.Children = t.Root.Children[0 : index+1]
					} else {
						t.Root.Children = append(t.Root.Children[0:index+1], t.Root.Children[index+2:]...)
					}
				}

				if len(t.Root.List) == 0 {
					t.Root.List = node.List
					t.Root.Children = node.Children
					for _, child := range node.Children {
						child.Parent = t.Root
					}
				}
			}

			newTree := &BTree{
				M:    t.M,
				Root: t.Root,
			}
			newTree.Delete(data)
		}
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
