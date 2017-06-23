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
func Test_AddTwoNumbers(t *testing.T) {
	log.Println(12 % 10)
	l1 := &medium.ListNode{
		Val: 2,
		Next: &medium.ListNode{
			Val: 4,
			Next: &medium.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &medium.ListNode{
		Val: 5,
		Next: &medium.ListNode{
			Val: 6,
			Next: &medium.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	res := medium.AddTwoNumbers(l1, l2)
	node := res
	for {
		if node == nil {
			break
		}
		log.Println(node.Val)
		node = node.Next
	}

}
