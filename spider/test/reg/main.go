package main

import (
	"./info"
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
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
	all, err := ioutil.ReadAll(resp.Body)

	match := regexp.MustCompile(`<a href="//(www.bilibili.com/[^\"]*)"><span>([^<]*)</span></a>`)
	data := match.FindAllSubmatch(all, -1)

	//file, e := os.Create("INFO")
	//DealErr(e)
	//defer file.Close()

	//创建INFO数组
	var arr []info.INFO

	for i, v := range data {
		//for _, m := range v {
		//	n, err2 := file.Write(m)
		//	DealErr(err2)
		//	fmt.Printf("写入 %d 个字节n", n)
		//}
		fmt.Println()
		//创建INFO
		info := info.INFO{
			i + 1,
			string(v[2]),
			string(v[1]),
		}
		arr = append(arr, info)

	}
	//file.Sync()
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

//448. 找到所有数组中消失的数字
//给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
//
//找到所有在 [1, n] 范围之间没有出现在数组中的数字。
//
//您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
//
//示例:
//
//输入:
//[4,3,2,7,8,2,3,1]
//
//输出:
//[5,6]
func findDisappearedNumbers(nums []int) []int {
	var arr []int
	for i := 0; i < len(nums); i++ {
		v := int(math.Abs(float64(nums[i])) - 1)
		if nums[v] < 0 {
			continue
		}
		nums[v] = nums[v] * -1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			arr = append(arr, i+1)
		}

	}
	return arr
}
