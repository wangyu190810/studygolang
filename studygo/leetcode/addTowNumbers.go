package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var nums = 0
	var num1 = 0
	var num2 = 0
	var first = true
	var listNode = ListNode{}
	for {
		if l1.Next == nil {
			num1 = l1.Val
		} else {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2.Next == nil {
			num2 = l2.Val
		} else {
			num2 = l2.Val
			l2 = l2.Next
		}

		nums = num1 + num2
		fmt.Printf("nums", nums)
		if first {
			listNode = ListNode{nums, nil}
			first = false
		}
		listNode.Next = &ListNode{nums, &listNode}
		// listNode.Next = &now

		// listNode =
		// listNode.Val = nums
		// listNode.Next = &listNode_print

		if l1.Next == nil && l2.Next == nil {
			break
		}

	}
	return &listNode

}

func RunaddTwoNumbers() {
	l1 := ListNode{1, nil}
	l2 := ListNode{2, nil}
	l3 := ListNode{3, &l1}
	l4 := ListNode{4, &l2}
	fmt.Println("data %V", l3, l4)

	data := addTwoNumbers(&l3, &l4)
	// data := &l4
	nums := 1
	for {
		if data.Next == nil {
			fmt.Println("data", data.Val)
			break
		} else {
			val := data.Val
			data = data.Next
			fmt.Println("data", val)
		}
		nums += 1
		if nums == 7 {
			break
		}
	}
	fmt.Printf("data %V", data)
}
