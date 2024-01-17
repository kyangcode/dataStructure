package main

import (
	"fmt"
	"strconv"
)

type User struct {
	ID   int
	Name string
}

func (u *User) Less(data Data) bool {
	return u.ID < data.(*User).ID
}

func main() {
	tree := NewBTree(3)

	for i := 1; i <= 100; i++ {
		tree.Insert(&User{
			ID:   i,
			Name: "n" + strconv.Itoa(i),
		})
	}

	data := tree.Search(&User{
		ID: 88,
	})
	fmt.Println(data)
}
