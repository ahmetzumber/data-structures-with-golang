package linkedlist

import "fmt"

type Node struct {
	Data int
	NextNode *Node
}

type LinkedList struct {
	Head *Node
	Length int
}

func (l *LinkedList) AddLast(newNode *Node){
	if l.Head == nil {
		l.Head = newNode
	}else {
		tempNode := l.Head
		for tempNode.NextNode != nil {
			tempNode = tempNode.NextNode
		}
		tempNode.NextNode = newNode
	}
}

func (l *LinkedList) Print(){
	tempNode := l.Head
	for tempNode != nil {
		fmt.Print(tempNode.Data)
		fmt.Print(" -> ")
		tempNode = tempNode.NextNode
	}
	fmt.Println("null")
}