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

func (l *LinkedList) AddFirst(newNode *Node){
	newNode.NextNode = l.Head
	l.Head = newNode
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

func (l *LinkedList) AddAfter(search *Node, newNode *Node){
	if l.Head == nil {
		fmt.Println("Empty list!!")
	}else {
		temp := l.Head
		for temp != nil && temp.Data != search.Data {
			temp = temp.NextNode
		}
		if temp != nil {
			newNode.NextNode = temp.NextNode
			temp.NextNode = newNode
		}
	}
}

func (l *LinkedList) Remove(removedData int){

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