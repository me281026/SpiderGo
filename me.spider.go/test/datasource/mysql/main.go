package main

import (
	"./connection"
	"fmt"
)

func main() {
	connection.InitDB()
	users := connection.QueryUserInfo()
	for i, user := range users {
		fmt.Printf("%s\n", "用户信息", i)
		user.ToString()
	}
}
