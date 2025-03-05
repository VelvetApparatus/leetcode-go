package _25

import (
	"strconv"
	"strings"
)

/*

	Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

	k is a positive integer and is less than or equal to the length of the linked list.
	If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

	You may not alter the values in the list's nodes, only nodes themselves may be changed.


	Example 1:
	Input: head = [1,2,3,4,5], k = 2
	Output: [2,1,4,3,5]

	URL: https://leetcode.com/problems/reverse-nodes-in-k-group/description/

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	msg := strings.Builder{}
	msg.WriteString("[ ")
	for node := l; node != nil; node = node.Next {
		msg.WriteString(strconv.Itoa(node.Val))
		if node.Next != nil {
			msg.WriteString(", ")
		}
	}
	msg.WriteString(" ]")
	return msg.String()
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	for node := head; node != nil && count < k; node = node.Next {
		count++
	}
	if count < k {
		return head
	}

	var prev *ListNode
	curr := head
	for i := 0; i < k; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	head.Next = reverseKGroup(curr, k)

	return prev
}
