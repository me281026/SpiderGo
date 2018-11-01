package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("11111111111111111111111111")
	}()

	go func() {
		fmt.Println("222222222222222222222222222")
	}()

	time.Sleep(time.Second * 10)

	fmt.Println("0000000000000000")

}
