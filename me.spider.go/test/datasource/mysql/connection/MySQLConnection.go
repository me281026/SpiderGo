package connection

import (
	"../user"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//数据库连接
var DBConnection *sql.DB

//连接信息
const (
	UserName = "user"
	PassWord = "123456"
	Ip       = "192.168.0.1"
	Port     = "3306"
	DataBase = "test"
)

//建立初始化连接池
func InitDB() {
	//连接串"用户名:密码@tcp(IP:端口)/数据库?charset=utf8mb4"
	path := strings.Join([]string{UserName, ":", PassWord, "@tcp(", Ip, ":", Port, ")/", DataBase, "?charset=utf8"}, "")
	//打开数据库连接
	DBConnection, _ := sql.Open("mysql", path)
	//设置数据库连接数
	DBConnection.SetConnMaxLifetime(10)
	//设置最大空闲连接数
	DBConnection.SetMaxIdleConns(1)
	//测试连接
	if err := DBConnection.Ping(); err != nil {
		fmt.Println("数据库连接失败....")
		return
	}
	fmt.Println("数据库连接成功....")

}

//查询多个记录
func QueryUserInfo() []user.User {
	//执行查询语句
	query, err := DBConnection.Query("SELECT USERINFOID,USERID,PASSWORD,USERNAME" +
		",SEX,AGE,BIRTHDAY,PHONENUM FROM crf_bi.bi_userinfo")
	if err != nil {
		fmt.Println(err)
	}
	var users []user.User
	//循环抓取
	for query.Next() {
		var u user.User
		//user赋值
		err := query.Scan(&u.USERINFOID, &u.USERID, &u.PASSWORD, &u.USERNAME, &u.SEX, &u.AGE, &u.BIRTHDAY, &u.PHONENUM)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, u)
	}
	return users
}

//查询一条记录
func QueryRow() user.User {
	var u user.User
	err := DBConnection.QueryRow("SELECT USERINFOID,USERID,PASSWORD,USERNAME"+
		",SEX,AGE,BIRTHDAY,PHONENUM FROM crf_bi.bi_userinfo where USERINFOID = 1").Scan(&u.USERINFOID, &u.USERID, &u.PASSWORD, &u.USERNAME, &u.SEX, &u.AGE, &u.BIRTHDAY, &u.PHONENUM)
	if err != nil {
		fmt.Println(err)
	}
	u.ToString()
	return u
}
