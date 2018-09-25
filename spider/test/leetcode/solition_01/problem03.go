package main

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
