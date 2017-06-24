package medium

/**
	Add Two Numbers

question:
	给的2个链表，已倒序的方式存储2个非负的int，要求把这2个数相加，结果存入一个新的链表中。
	假定 2个数 不以0开头  除非 是0

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
即：342+465=807

思路：
	递归2个链表，若单次相加大于10则求除求余，直到2条链表的next都为nil
	最后判断result中最后一个节点的val是否为0，若为0则删除该节点

时间复杂度：O(max(m,n))
空间复杂度：O(max(m,n))。[新链表最大为O(max(m,n))+1]
**/

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//AddTwoNumbers 计算2个链表数
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resu := &ListNode{}
	add(resu, l1, l2)
	return resu
}
func add(cur *ListNode, l1 *ListNode, l2 *ListNode) {
	l1v, l2v := 0, 0
	if l1 != nil {
		l1v = l1.Val
	}
	if l2 != nil {
		l2v = l2.Val
	}
	i, v := 0, l1v+l2v
	if v >= 10 {
		i = v / 10
		v = v % 10
	}
	cur.Val = cur.Val + v
	if cur.Val >= 10 {
		i = i + cur.Val/10
		cur.Val = cur.Val % 10
	}
	nextNode := &ListNode{
		Val: i,
	}
	cur.Next = nextNode

	if l1 == nil {
		if l2.Next != nil {
			add(cur.Next, nil, l2.Next)
		} else {
			if cur.Next.Val == 0 {
				cur.Next = nil
			}
		}
	} else if l2 == nil {
		if l1.Next != nil {
			add(cur.Next, l1.Next, nil)
		} else {
			if cur.Next.Val == 0 {
				cur.Next = nil
			}
		}
	} else {
		if l1.Next == nil && l2.Next != nil {
			add(cur.Next, nil, l2.Next)
		} else if l1.Next != nil && l2.Next == nil {
			add(cur.Next, l1.Next, nil)
		} else if l1.Next != nil && l2.Next != nil {
			add(cur.Next, l1.Next, l2.Next)
		} else {
			if cur.Next.Val == 0 {
				cur.Next = nil
			}
		}
	}
	//log.Println(cur)
}
