package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var memoryAdd map[*ListNode]int = make(map[*ListNode]int)

	for headA != nil {
		memoryAdd[&*headA] = headA.Val
		headA = headA.Next
	}

	for headB != nil {
		if memoryAdd[&*headB] >= 1 {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
