package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var ans *ListNode

	if list1.Val < list2.Val {
		ans = list1
		ans.Next = mergeTwoLists(list1.Next, list2)
	} else {
		ans = list2
		ans.Next = mergeTwoLists(list2.Next, list1)
	}

	return ans
}
