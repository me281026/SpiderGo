package info

import "fmt"

type INFO struct {
	Id    int
	Block string
	Url   string
}

func (i *INFO) ToString() {
	fmt.Printf("id : %d , block : %s , url : %s \n", i.Id, i.Block, i.Url)
}
