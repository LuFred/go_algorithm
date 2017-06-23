package medium

/**
	Add Two Numbers

question:
	给的2个链表，已倒序的方式存储2个非负的int，要求把这2个数相加，结果存入一个新的链表中。
	假定 2个数 不以0开头  除非 是0

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
即：342+465=807

**/

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//AddTwoNumbers 计算2个链表数
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resu := &ListNode{}
	t(resu, l1, l2)
	return resu
}
func t(cur *ListNode, l1 *ListNode, l2 *ListNode) {
	l1v, l2v := 0, 0

	if l1 != nil {
		l1v = l1.Val
	}
	if l2 != nil {
		l2v = l2.Val
	}

	sum := l1v + l2v
	i, v := 0, sum
	if sum >= 10 {
		v = sum % 10
		i = sum / 10
	}
	curNode := &ListNode{
		Val:  v,
		Next: &ListNode{},
	}
	cur.Next = curNode
	cur.Val = cur.Val + i
	//log.Println(cur)
	if l1.Next != nil || l2.Next != nil {
		t(cur, l1.Next, l2.Next)
	}
}
