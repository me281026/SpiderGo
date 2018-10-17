package main

import (
	"./info"
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	resp, err := http.Get("https://www.bilibili.com")
	DealErr(err)

	defer resp.Body.Close()
	//判断请求是否成功
	if resp.StatusCode != 200 {
		fmt.Println("success!!!!!!")
	}

	all, err := ioutil.ReadAll(resp.Body)

	match := regexp.MustCompile(`<a href="//(www.bilibili.com/[^\"]*)"><span>([^<]*)</span></a>`)
	data := match.FindAllSubmatch(all, -1)

	//创建INFO数组
	var arr []info.INFO

	for i, v := range data {
		fmt.Println()
		//创建INFO
		info := info.INFO{
			i + 1,
			string(v[2]),
			string(v[1]),
		}
		arr = append(arr, info)

	}

	fmt.Printf("total url num : %d\n", len(data))

	DealErr(err)
	//创建文件,写入数据
	file, i := os.Create("D:\\GoProject\\SpiderGo\\spider\\test\\reg\\" + "Info_" + strconv.Itoa(int(time.Now().Unix())) + ".txt")
	defer file.Close()
	DealErr(i)

	//创建bufio
	writer := bufio.NewWriter(file)

	//遍历数组
	for _, v := range arr {
		v.ToString()
		writer.WriteString(v.StringData() + " \n")
	}
	writer.Flush()

}

func DealErr(err error) {
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}
