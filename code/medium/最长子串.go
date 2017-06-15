package medium

/**
	Longest Substring Without Repeating Characters

question:
	Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

区分子串和子序列：
	子串   ：串中任意个连续的字符组成的子序列称为该串的子串
	子序列：给定一个数列(x_n)。在这个数列里，任取无穷多项，不改变它们在原来数列中的先后次序，得到的数列称为是原来数列的一个子数列
**/

//LengthOfLongestSubstring 方法1
func LengthOfLongestSubstring(s string) int {
	cur, max, rn := 1, 0, 0
	sl := len(s)
	m := make(map[byte]int)
	for i := 1; i <= sl; i++ {
		sbyte := s[i-1] //获取字符字节码
		//判断map中是否存在该字节码
		if m[sbyte] < 1 || m[sbyte] < cur { //若不存在或位置小于当前标志位
			m[s[i-1]] = i //讲字节码和位置作为kv存储
			max = max + 1 //当前最大字串+1
			if max > rn {
				rn = max
			}
		} else { //若存在
			if max > rn {
				rn = max
			}
			tempMax := i - m[sbyte] //截止冲突最长子串
			cur = m[s[i-1]] + 1
			m[s[i-1]] = i
			if tempMax > rn { //大于当前最大长度则替换
				rn = tempMax
			}
			max = tempMax //截止冲突最长子串赋值给max
		}
	}
	return rn
}

//LengthOfLongestSubstring2  方法2
//相对上面一种思路一致，代码更简洁
func LengthOfLongestSubstring2(s string) int {
	rn, cur := 0, 0
	m := make(map[byte]int)
	for j := 0; j < len(s); j++ {
		if m[s[j]] > 0 {
			if m[s[j]] > cur {
				cur = m[s[j]]
			}
		}
		if j-cur+1 > rn {
			rn = j - cur + 1
		}
		m[s[j]] = j + 1
	}
	return int(rn)
}
