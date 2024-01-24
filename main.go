package main

func main() {
	tree := NewRBTree()
	for i := 1; i <= 100; i++ {
		tree.Insert(i)
	}
}
