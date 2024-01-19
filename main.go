package main

import (
	"fmt"
	"sort"
)

//type User struct {
//	ID   int
//	Name string
//}
//
//func (u *User) Less(data Data) bool {
//	return u.ID < data.(*User).ID
//}

func main() {
	//tree := NewBTree(3)
	//
	//for i := 1; i <= 100; i++ {
	//	tree.Insert(&User{
	//		ID:   i,
	//		Name: "n" + strconv.Itoa(i),
	//	})
	//}

	s := []int{1, 2, 3}
	r := sort.Search(3, func(i int) bool {
		return 0 < s[i]
	})
	fmt.Println(r)
}
