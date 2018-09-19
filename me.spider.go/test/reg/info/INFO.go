package info

import (
	"bytes"
	"fmt"
	"strconv"
)

type INFO struct {
	Id    int
	Block string
	Url   string
}

func (i *INFO) ToString() {
	fmt.Printf("id : %d , block : %s , url : %s \n", i.Id, i.Block, i.Url)
}

//返回string
func (i *INFO) StringData() string {
	var str bytes.Buffer
	str.WriteString("Id : ")
	str.WriteString(strconv.Itoa(i.Id))
	str.WriteString(" , Block : ")
	str.WriteString(i.Block)
	str.WriteString(" , url : ")
	str.WriteString(i.Url)
	return str.String()
}
