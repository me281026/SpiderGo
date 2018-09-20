package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	now := time.Now()
	unix := strconv.Itoa(int(time.Now().Unix()))
	fmt.Println(now)
	fmt.Println(unix)

}
