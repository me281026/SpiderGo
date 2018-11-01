package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("11111111111111111111")
		wg.Done()
	}()
	go func() {
		fmt.Println("222222222222222222222")
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("000000000000000000000000")
}
