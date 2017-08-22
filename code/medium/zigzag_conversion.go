package medium

import (
	"bytes"
	"fmt"
)

/**
 ZigZag Conversion
给定一个字符串以及一个numRows行号，将其表示为一个倒Z字形的字符排列，按顺序拼接每行的字符串输出
例如:
	字符串：PAYPALISHIRING
	行数:3
		效果如下：
		P   A   H   N
		A P L S I I G
		Y   I   R
	结果为：PAHNAPLSIIGYIR
思路：
	逆向思考，上图的字符串其实就是N个字符‘N’拼接组成，
	看形状可以想到一个‘N’首行和末行分别由2个字符组成，中间位置需要3个字符组成一行
	所以要打印出结果需要分以下几中情况循环抓取字符
		1.首行与末行
			：2个字符之间的间隔是(numRows-1)*2个字符，
			例子中的第一行，获取到第一个字母P之后需要获取的是 在此基础上跳到第(3-1)*2=4个字母即A
		2.每个‘N’形状中间段都有3个字母组成，而这里面又有2中获取情况
			a.在竖直线上的情况
				 (i-1)*2
			b.在斜线上的情况
				 (numRows-i)*2
时间复杂度：O(n)
**/
//ZigZagConversion 代码实现
func ZigZagConversion(s string, numRows int) string {
	var buffer bytes.Buffer //输出结果对象
	sl := len(s)            //输入字符串长度
	var curS int            //当前待获取字符
	if numRows < 2 {
		return s
	}
	for i := 1; i <= numRows; i++ {
		opt := 1
		curS = i
		for {
			if curS > sl {
				fmt.Println("") //换行
				break
			}
			char := s[curS-1]
			fmt.Print(string([]byte{char}))
			buffer.WriteByte(char)
			if i == 1 || i == numRows {
				curS = curS + (numRows-1)*2
				for sp := 1; sp < numRows-1; sp++ { //打印空格
					fmt.Print(" ")
				}
			} else if opt%2 != 0 { //奇数次操作：相隔待获取字符串位：(numRows-i)*2
				curS = curS + (numRows-i)*2
				for sp := 0; sp < numRows-i-1; sp++ { //打印空格
					fmt.Print(" ")
				}
			} else { //偶数次操作:相隔待获取字符串位：(i-1)*2
				curS = curS + (i-1)*2
				for sp := 0; sp < i-2; sp++ { //打印空格
					fmt.Print(" ")
				}
			}
			opt++
		}
	}
	return buffer.String()
}
