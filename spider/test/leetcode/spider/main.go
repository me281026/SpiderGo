package main

import (
	"./problem"
	bufio "bufio"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	time "time"
)

func main() {
	//获取链接
	url := "https://leetcode-cn.com/problemset/all/"
	resp, err := http.Get(url)
	Dealerr(err)
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	Dealerr(err)
	//正则匹配
	compile := regexp.MustCompile(`<a href="([^\"]*)">([^<]*)</a>`)
	submatch := compile.FindAllSubmatch(all, -1)
	//创建问题数组
	var plems []problem.Problem
	for i, v := range submatch {
		//创建问题结构体
		problem := problem.Problem{
			i + 1,
			string(v[1]),
			string(v[2]),
		}
		plems = append(plems, problem)
	}
	//创建文件,写入数据D:\GoProject\SpiderGo\spider\test\leetcode\spider
	file, i := os.Create("D:\\GoProject\\SpiderGo\\spider\\test\\leetcode\\spider" + "Info_" + strconv.Itoa(int(time.Now().Unix())) + ".txt")
	defer file.Close()
	Dealerr(i)

	//创建bufio
	writer := bufio.NewWriter(file)

	//遍历数组
	for _, v := range plems {
		v.ToString()
		writer.WriteString(v.StringData() + " \n")
	}
	writer.Flush()

}

func Dealerr(err error) {
	if err != nil {
		println(err.Error())
	}
}
