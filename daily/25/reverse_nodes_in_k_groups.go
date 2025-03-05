package _25

import (
	"strconv"
	"strings"
)

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
