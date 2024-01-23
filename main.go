package main

func main() {
	tree := NewRBTree()
	for i := 1; i <= 6; i++ {
		tree.Insert(i)
	}
}
