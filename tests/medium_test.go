package tests

import (
	"go_algorithm/code/medium"
	"log"
	"testing"
)

//glilfygihi
/**[length=5
g 7
l  4
i  8
f 5
y 6
h 9

]
**/
func Test_LengthOfLongestSubstring(t *testing.T) {
	l1 := medium.LengthOfLongestSubstring("assdw")
	l2 := medium.LengthOfLongestSubstring2("assdw")
	log.Println(l1)
	log.Println(l2)
}
