package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	testData := [][2]*ListNode{
		{
			createList([]int{2, 4, 3}),
			createList([]int{5, 6, 4}),
		},
		{
			createList([]int{0}),
			createList([]int{0}),
		},
		{
			createList([]int{9, 9, 9, 9, 9, 9, 9}),
			createList([]int{9, 9, 9, 9}),
		},
	}

	for _, v := range testData {
		n := addTwoNumbers(v[0], v[1])
		for n.Next != nil {
			fmt.Printf("%d,", n.Val)
			n = n.Next
		}
		fmt.Println("\n===\n")
	}
}

func createList(intList []int) *ListNode {
	var res ListNode

	for tmp, i := &res, 0; ; {
		tmp.Val = intList[i]

		i++
		if i == len(intList) {
			break
		} else {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		}
	}

	return &res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res ListNode
	for tmpList, tempVal := &res, 0; l1 != nil || l2 != nil || tempVal > 0; tmpList = tmpList.Next {
		if l1 != nil {
			tempVal += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tempVal += l2.Val
			l2 = l2.Next
		}

		tmpList.Next = &ListNode{Val: tempVal % 10}
		tempVal = tempVal / 10
	}

	return res.Next
}
