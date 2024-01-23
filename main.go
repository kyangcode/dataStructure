package main

func main() {
	tree := NewRBTree()
	for i := 1; i <= 5; i++ {
		tree.Insert(i)
	}
}
