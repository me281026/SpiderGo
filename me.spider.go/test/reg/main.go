package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://www.bilibili.com")
	if err != nil {
		fmt.Println("err:", err.Error())
	}

	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)

	match := regexp.MustCompile(`<a href="//www.bilibili.com/.*</span></a>`)
	data := match.FindAll(all, -1)
	for _, v := range data {
		fmt.Printf("%s\n", v)
	}

	if err != nil {
		fmt.Println("err:", err.Error())
	}

}
