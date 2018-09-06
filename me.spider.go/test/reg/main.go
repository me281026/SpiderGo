package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {
	resp, err := http.Get("https://www.bilibili.com")
	DealErr(err)

	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)

	match := regexp.MustCompile(`<a href="//(www.bilibili.com/[^>]*)><span>([^<]*)</span></a>`)
	data := match.FindAllSubmatch(all, -1)

	file, e := os.Open("D://GoProject//SpiderGo//me.spider.go//test//reg//HTML")
	DealErr(e)
	defer file.Close()

	for _, v := range data {
		for _, m := range v {
			n, err2 := file.Write(m)
			DealErr(err2)
			fmt.Printf("写入 %d 个字节n", n)
		}

		fmt.Printf("area: %s , url: %s\n", v[2], v[1])

	}
	file.Sync()
	fmt.Printf("total url num : %d\n", len(data))

	DealErr(err)

}

func DealErr(err error) {
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}
