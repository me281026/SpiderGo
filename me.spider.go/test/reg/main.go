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

	match := regexp.MustCompile(`<a href="//(www.bilibili.com/[^>]*)><span>([^<]*)</span></a>`)
	data := match.FindAllSubmatch(all, -1)
	for _, v := range data {
		fmt.Printf("area: %s , url: %s\n", v[2], v[1])
	}
	fmt.Printf("total url num : %d\n", len(data))

	if err != nil {
		fmt.Println("err:", err.Error())
	}

}
