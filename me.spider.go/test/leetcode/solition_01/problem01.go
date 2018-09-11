package main

import (
	"bytes"
	"strings"
)

func main() {

}

//389. 找不同
//给定两个字符串 s 和 t，它们只包含小写字母。
//字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
//请找出在 t 中被添加的字母。

func findTheDifference(s string, t string) byte {
	num := 0
	for _, v := range s {
		num ^= int(v)
	}
	for _, v := range t {
		num ^= int(v)
	}
	return byte(rune(num))
}

//349. 两个数组的交集
//给定两个数组，编写一个函数来计算它们的交集。
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [9,4]
func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	m := make(map[int]int)
	for _, i := range nums1 {
		m[i] = 1
	}
	for _, j := range nums2 {
		if m[j] == 1 {
			result = append(result, j)
			m[j] = 2
		}
	}

	return result
}

//485. 最大连续1的个数
//给定一个二进制数组， 计算其中最大连续1的个数。
//输入: [1,1,0,1,1,1]
//输出: 3
//解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
func findMaxConsecutiveOnes01(nums []int) int {
	var a, b int
	for _, v := range nums {
		if v != 1 {
			a = 0
		} else {
			a++
			if a > b {
				b = a
			}
		}
	}
	return b
}

func findMaxConsecutiveOnes02(nums []int) int {
	var a, b int
	for _, v := range nums {
		if b+v > a {
			a = b + v
		}
		b = (b + v) * v
	}
	return b
}

//788. 旋转数字
func rotatedDigits(N int) int {
	count := 0
	for i := 1; i <= N; i++ {
		validFound := false
		for i > 0 {
			if i%10 == 2 {
				validFound = true
			}
			if i%10 == 5 {
				validFound = true
			}
			if i%10 == 6 {
				validFound = true
			}
			if i%10 == 9 {
				validFound = true
			}
			if i%10 == 3 {
				validFound = false
				break
			}
			if i%10 == 4 {
				validFound = false
				break
			}
			if i%10 == 7 {
				validFound = false
				break
			}
			i = i / 10
		}

		if validFound {
			count++
		}
	}
	return count

}

//122. 买卖股票的最佳时机 II
func maxProfit(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}

//824. 山羊拉丁文
func toGoatLatin(S string) string {
	words := strings.Fields(S)
	vowels := map[byte]bool{
		'a': true, 'A': true,
		'e': true, 'E': true,
		'i': true, 'I': true,
		'o': true, 'O': true,
		'u': true, 'U': true,
	}
	var buffer bytes.Buffer

	for i, word := range words {
		if vowels[word[0]] {
			buffer.WriteString(word)
		} else {
			buffer.WriteString(word[1:])
			buffer.WriteString(string(word[0]))
		}
		buffer.WriteString("ma")
		for j := 1; j <= i+1; j++ {
			buffer.WriteString("a")
		}
		if i != len(words)-1 {
			buffer.WriteString(" ")
		}
	}

	return buffer.String()
}

//242. 有效的字母异位词
//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。
//输入: s = "anagram", t = "nagaram"
//输出: true
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var arr [26]int
	for _, v := range s {
		arr[v-'a']++
	}
	for _, v := range t {
		arr[v-'a']--
	}
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

//447. 回旋镖的数量
func numberOfBoomerangs(points [][]int) int {
	result := 0
	for i, v := range points {
		p := map[int]int{}
		for m, n := range points {
			if i == m {
				continue
			}
			k1, k2 := v[0]-n[0], v[1]-n[1]
			p[k1*k1+k2*k2] += 1
		}
		for t := range p {
			if p[t] > 1 {
				result += p[t] * (p[t] - 1)
			}
		}

	}
	return result
}

//268. 缺失数字
func missingNumber(nums []int) int {
	len := len(nums)
	a := len * (len + 1) / 2
	var b int
	for _, num := range nums {
		b += num
	}
	return a - b

}
