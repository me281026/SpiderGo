package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.jd.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//读取网页数据
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", all)
}
