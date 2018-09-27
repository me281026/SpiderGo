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
