package main

import (
	"log"
)

func main() {
	r := lengthOfLongestSubstring(test)
	log.Println(test[0])
	log.Println(r)
	return
}

var test = "abcabcbb"

func lengthOfLongestSubstring(s string) int {
	curr, max, rn := 1, 0, 0
	sl := len(s)
	m := make(map[byte]int)
	for i := 1; i <= sl; i++ {
		sbyte := s[i] //获取字符字节码
		//判断map中是否存在改字节码
		if m[sbyte] < 1 { //若不存在
			m[s[i]] = i   //讲字节码和位置作为kv存储
			max = max + 1 //当前最大字串+1
			curr = curr
		} else { //若存在
			if max > rn {
				rn = max
			}

			tempMax := i - m[sbyte] //截止冲突最长子串

			curr = m[sbyte] + 1 //取出位置+1后赋值给坐标标记位
			m[s[i]] = i
			if tempMax > rn { //
				rn = tempMax
			}
			max = tempMax
		}
	}
	return rn
}
