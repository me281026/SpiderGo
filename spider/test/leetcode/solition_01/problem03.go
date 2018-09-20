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
