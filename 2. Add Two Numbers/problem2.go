package main

type Node struct {
	prev *Node
	next *Node
}

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
   var head *ListNode
   var curr *ListNode
   for l1 != nil || l2 != nil || carry > 0 {
	   val := carry
	   if l1 != nil {
		   val += l1.Val
		   l1 = l1.Next
	   }
	   if l2 != nil {
		   val += l2.Val
		   l2 = l2.Next
	   }
	   carry = val / 10
	   val = val % 10
	   n := &ListNode{Val:val, Next:nil}
	   if head == nil {
		   head = n
		   curr = n
	   } else {
		   curr.Next = n
		   curr = n
	   }
   }
   return head
}