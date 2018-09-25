package problem

import (
	"bytes"
	"fmt"
	"strconv"
)

type Problem struct {
	Id          int
	Url         string
	ProblemName string
}

func (p *Problem) ToString() {
	fmt.Printf("Id : %d , Url : %s , ProblemName : %s \n", p.Id, p.Url, p.ProblemName)
}

//返回string
func (p *Problem) StringData() string {
	var str bytes.Buffer
	str.WriteString("Id : ")
	str.WriteString(strconv.Itoa(p.Id))
	str.WriteString(" , Url : ")
	str.WriteString(p.Url)
	str.WriteString(" , ProblemName : ")
	str.WriteString(p.ProblemName)
	return str.String()
}
