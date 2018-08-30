package connection

import (
	"database/sql"
	"fmt"
	"strings"
)

//数据库连接
var DBConnection *sql.DB

//连接信息
const (
	UserName = "mldeploy"
	PassWord = "123456"
	Ip       = "192.168.60.99"
	Port     = "3306"
	DataBase = "crf_bi"
)

//建立初始化连接池
func InitDB() {
	//连接串"用户名:密码@tcp(IP:端口)/数据库?charset=utf8mb4"
	path := strings.Join([]string{UserName, ":", PassWord, "@tcp(", Ip, ":", Port, "/", DataBase, "?charset=utf8mb4"}, "")
	//打开数据库连接
	db, _ := sql.Open("mysql", path)
	//设置数据库连接数
	db.SetConnMaxLifetime(100)
	//设置最大空闲连接数
	db.SetMaxIdleConns(10)
	//测试连接
	if err := db.Ping(); err != nil {
		fmt.Println("数据库连接失败....")
		return
	}
	fmt.Println("数据库连接成功....")

}
