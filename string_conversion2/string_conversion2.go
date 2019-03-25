package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[rune]int) //定义每个字符出现的最后位置

	start := 0     //字符串的起始位置
	maxLength := 0 //当前字符串最大长度

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	fmt.Printf("%x\n", lastOccurred)
	return maxLength
}

func main() {
	s := "abcc一一二三四"
	fmt.Println(lengthOfLongestSubstring(s))
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) \n", i, ch)
	}

}
