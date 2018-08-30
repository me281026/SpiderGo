package user

import "fmt"

type User struct {
	USERINFOID int
	USERID     string
	PASSWORD   string
	USERNAME   string
	SEX        string
	BIRTHDAY   string
	AGE        int
	PHONENUM   string
}

func (user *User) toString() {
	fmt.Println(user.USERID + "," + user.USERNAME + "," + user.SEX)
}
