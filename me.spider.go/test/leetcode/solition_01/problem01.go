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
