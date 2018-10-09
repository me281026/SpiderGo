package main

import "math"

func main() {

}

//844. 比较含退格的字符串
//给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。
//输入：S = "ab#c", T = "ad#c"
//输出：true
//解释：S 和 T 都会变成 “ac”。
//输入：S = "a#c", T = "b"
//输出：false
//解释：S 会变成 “c”，但 T 仍然是 “b”。
func backspaceCompare(S string, T string) bool {
	//直接判断两个字符串是否一致
	if S == T {
		return true
	}
	//if strings.Contains(S,"#")
	return false
}

//217. 存在重复元素
func containsDuplicate(nums []int) bool {
	var m = make(map[int]int)
	for i, v := range nums {
		_, flag := m[v]
		if flag {
			return true
		}
		m[v] = i
	}
	return false
}

//383. 赎金信
//给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，
//判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。如果可以构成，返回 true ；否则返回 false。
func canConstruct(ransomNote string, magazine string) bool {
	var arr [26]int
	for _, v := range []rune(ransomNote) {
		arr[v-'a']++
	}
	for _, v := range []rune(magazine) {
		a := arr[v-'a']
		a--
		if a < 0 {
			return false
		}
	}
	return true

}

//70. 爬楼梯
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	a := 1
	b := 2
	for i := 3; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b

}

//908. 最小差值
//给定一个整数数组 A，对于每个整数 A[i]，我们可以选择任意 x 满足 -K <= x <= K，并将 x 加到 A[i] 中。
//
//在此过程之后，我们得到一些数组 B。
//
//返回 B 的最大值和 B 的最小值之间可能存在的最小差值。
//示例 1：
//
//输入：A = [1], K = 0
//输出：0
//解释：B = [1]
//示例 2：
//
//输入：A = [0,10], K = 2
//输出：6
//解释：B = [2,8]
//示例 3：
//
//输入：A = [1,3,6], K = 3
//输出：0
//解释：B = [3,3,3] 或 B = [4,4,4]
func smallestRangeI(A []int, K int) int {
	min := float64(A[0])
	max := float64(A[0])
	for _, v := range A {

		min = math.Min(min, float64(v))
		max = math.Max(max, float64(v))
	}
	return int(math.Max(0, max-min-2*float64(K)))

}

//342. 4的幂
//给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。
//输入: 16
//输出: true
//示例 2:
//
//输入: 5
//输出: false
func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && num&0x55555555 == num

}

//231. 2的幂
//给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
//
//示例 1:
//
//输入: 1
//输出: true
//解释: 20 = 1
//示例 2:
//
//输入: 16
//输出: true
//解释: 24 = 16
//示例 3:
//
//输入: 218
//输出: false
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

//5. 最长回文子串
//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba"也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"
func longestPalindrome(s string) string {
	return ""
}

//263. 丑数
//编写一个程序判断给定的数是否为丑数。
//
//丑数就是只包含质因数 2, 3, 5 的正整数。
//
//示例 1:
//
//输入: 6
//输出: true
//解释: 6 = 2 × 3
//示例 2:
//
//输入: 8
//输出: true
//解释: 8 = 2 × 2 × 2
//示例 3:
//
//输入: 14
//输出: false
//解释: 14 不是丑数，因为它包含了另外一个质因数 7。
//说明：
//
//1 是丑数。
//输入不会超过 32 位有符号整数的范围: [−231,  231 − 1]。
func isUgly(num int) bool {

	for i := 2; i < 6 && num > 0; i++ {
		for num%i == 0 {
			num /= i
		}
	}
	return num == 1

}

func isUgly02(num int) bool {
	if num <= 0 {
		return false
	}
	for num%2 == 0 {
		num /= 2
	}
	for num%3 == 0 {
		num /= 3
	}
	for num%5 == 0 {
		num /= 5
	}
	return num == 1

}

//541. 反转字符串 II
//给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。
// 如果剩余少于 k 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，
// 并将剩余的字符保持原样。
//
//示例:
//
//输入: s = "abcdefg", k = 2
//输出: "bacdfeg"
//要求:
//
//该字符串只包含小写的英文字母。
//给定字符串的长度和 k 在[1, 10000]范围内。
func reverseStr(s string, k int) string {
	return ""
}

//405. 数字转换为十六进制数
func toHex(num int) string {
	var arr = [16]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	if num == 0 {
		return ""
	}
	var res string
	for num != 0 {
		res = string(arr[(num&15)]) + res

	}
	return ""

}

//888. 公平的糖果交换
func fairCandySwap(A []int, B []int) []int {
	return nil
}

//896. 单调数列
func isMonotonic(A []int) bool {
	flag := 0
	for i := 0; i < len(A)-1; i++ {
		num := 0
		if A[i] > A[i+1] {
			num = -1
		} else if A[i] < A[i+1] {
			num = 1
		}
		if num != 0 {
			if flag != num && flag != 0 {
				return false
			}
			flag = num
		}
	}
	return true
}
