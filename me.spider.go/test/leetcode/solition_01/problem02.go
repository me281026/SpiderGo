package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	s := "PPALLL"
	fmt.Println(checkRecord(s))

}

//860. 柠檬水找零
func lemonadeChange(bills []int) bool {
	amt05, amt10 := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			amt05++
		} else if bill == 10 {
			if amt05 == 0 {
				return false
			}
			amt05--
			amt10++
		} else {
			if amt05 > 0 && amt10 > 0 {
				amt05--
				amt10--
			} else if amt05 >= 3 {
				amt05 -= 3
			} else {
				return false
			}
		}
	}
	return true

}

//121. 买卖股票的最佳时机
func maxProfit02(prices []int) int {
	a := 0
	for i := 0; i < len(prices); i++ {
		c := prices[i] - prices[i-1]
		if c < 0 {
			c = 0
		}
		if c > a {
			a = c
		}
	}
	return a
}

//551. 学生出勤纪录
func checkRecord(s string) bool {
	arr := []rune(s)
	if len(arr) <= 3 || s != "AA" || s != "LLL" {
		return true
	}
	a, b := 0, 0
	for i, v := range arr {
		if v == 'A' {
			a++
		}
		if arr[len(arr)-1] == 'L' && i == len(arr)-1 {
			if arr[len(arr)-2] == 'L' && arr[len(arr)-3] == 'L' {
				b++
				break
			}
		} else {
			if arr[i] == 'L' && arr[i+1] == 'L' {
				b++
			}
		}
	}
	if a <= 1 && b <= 2 {
		return true
	} else {
		return false
	}
}

//551. 学生出勤纪录
func checkRecord02(s string) bool {
	b, _ := regexp.MatchString(`.*LLL.*|.*A.*A.*`, s)
	return !b
}

//551. 学生出勤纪录
func checkRecord03(s string) bool {
	return !(strings.Count(s, `A`) > 1) && !strings.Contains(s, `LLL`)
}
