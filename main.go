package main

import "fmt"

func main() {
	tree := NewAVLTree()

	for i := 1; i <= 5; i++ {
		tree.Insert(i)
	}

	node := tree.Search(3)
	fmt.Println(node.Key)
}
