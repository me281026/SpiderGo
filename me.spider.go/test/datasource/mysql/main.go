package main

import (
	"./connection"
	"./user"
	"fmt"
)

func main() {
	var u user.User
	connection.InitDB()
	err := connection.DBConnection.QueryRow("SELECT USERINFOID,USERID,PASSWORD,USERNAME"+
		",SEX,AGE,BIRTHDAY,PHONENUM FROM crf_bi.bi_userinfo where USERINFOID = 1").Scan(u.USERINFOID, u.USERID, u.PASSWORD, u.USERNAME, u.SEX, u.AGE, u.BIRTHDAY, u.PHONENUM)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
