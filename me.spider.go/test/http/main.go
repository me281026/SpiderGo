package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.zhihu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("%s\n", resp.Body)
}
