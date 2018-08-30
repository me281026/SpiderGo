package main

import (
	"fmt"
	"os/exec"
)

//在window环境下使用go命令
func main() {

	//go get -u github.com/go-sql-driver/mysql
	//第一个参数表示是cmd命令,第二个是目录,
	command := exec.Command("cmd", "D://GoProject//src", "get", "-u github.com/go-sql-driver/mysql")

	if run := command.Run(); run != nil {
		fmt.Println(run)
	}

}
