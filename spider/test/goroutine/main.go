package main

import "fmt"

func main() {
	//创建channel
	channel := make(chan struct{})
	count := 2
	//1
	go func() {
		fmt.Println("111111111111111111111")
		channel <- struct{}{}
	}()

	//2
	go func() {
		fmt.Println("222222222222222222222")
		channel <- struct{}{}
	}()

	//循环判断channel
	for range channel {
		count--
		if count == 0 {
			close(channel)
		}

	}
	fmt.Println("00000000000000")
}
