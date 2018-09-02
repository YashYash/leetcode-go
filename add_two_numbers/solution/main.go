package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// ErrInput - If inputs are invalid
var ErrInput = errors.New("INPUT ERROR")

// ListNode - a single node that composes the list
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedListNode the linked list of Items
type LinkedListNode struct {
	Head *ListNode
	Size int
}

func (ll *LinkedListNode) String() string {
	stringNodes := ""
	if ll.Head == nil {
		stringNodes = strconv.Itoa(ll.Head.Val)
	} else {
		last := ll.Head
		for {
			stringNodes = stringNodes + strconv.Itoa(last.Val) + " "
			if last.Next == nil {
				break
			}
			last = last.Next
		}
	}
	return stringNodes
}

func timeTrack(start time.Time, name string, l1 LinkedListNode, l2 LinkedListNode) {
	elapsed := time.Since(start)
	fmt.Println(name)
	fmt.Println("l1: ", l1.String())
	fmt.Println("l2: ", l2.String())
	fmt.Println("Run Time: " + elapsed.String())
	fmt.Println("------------------------------")
}

// Append - Add node to linked list
func (ll *LinkedListNode) Append(t int) {
	node := ListNode{
		Val:  t,
		Next: nil,
	}
	if ll.Head == nil {
		ll.Head = &node
	} else {
		last := ll.Head
		for {
			if last.Next == nil {
				break
			}
			last = last.Next
		}
		last.Next = &node
	}
	ll.Size++
}

// AddTwoNumbers - Solution lives here
func AddTwoNumbers(l1 LinkedListNode, l2 LinkedListNode, tc string) (LinkedListNode, error) {
	defer timeTrack(time.Now(), "AddTwoNumbers - Test Case: "+tc, l1, l2)
	var (
		l1Last   = l1.Head
		l2Last   = l2.Head
		carry    = 0
		listNode LinkedListNode
	)
	if l1Last.Next == nil && l2Last.Next == nil {
		listNode.Append(l1Last.Val + l2Last.Val)
	} else {
		for {
			sum := l1Last.Val + l2Last.Val
			val := sum + carry
			if val > 9 {
				carry = 1
			}
			listNode.Append(val % 10)
			if l1Last.Next == nil || l2Last.Next == nil {
				break
			}
			l1Last = l1Last.Next
			l2Last = l2Last.Next
		}
	}
	return listNode, nil
}

func main() {
	fmt.Println("Add Two Numbers")
}
