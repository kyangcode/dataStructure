package main

import "strconv"

type User struct {
	ID   int
	Name string
}

func (u *User) Less(data Data) bool {
	return u.ID < data.(*User).ID
}

func main() {
	tree := NewBTree(5)

	for i := 1; i <= 20; i++ {
		tree.Insert(&User{
			ID:   i,
			Name: "n" + strconv.Itoa(i),
		})
	}

	for i := 20; i >= 1; i-- {
		tree.Delete(&User{
			ID: i,
		})
	}

}
