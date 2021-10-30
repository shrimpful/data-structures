package linked_list

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Length++
}

func (l *LinkedList) PrintListData() {
	toPrint := l.Head
	for l.Length != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
		l.Length--
	}
	fmt.Printf("\n")
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.Length == 0 {
		return
	}
	if l.Head.Data == value {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	previousToDelete := l.Head
	for previousToDelete.Next.Data != value {
		if previousToDelete.Next.Next == nil {
			return
		}
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = previousToDelete.Next.Next
	l.Length--
}
