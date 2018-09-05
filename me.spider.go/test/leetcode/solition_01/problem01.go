package main

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
