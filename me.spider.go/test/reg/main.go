package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://bbs.nga.cn/index.php")
	if err != nil {
		fmt.Println("err:", err.Error())
	}

	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Printf("%s\n", all)

}
