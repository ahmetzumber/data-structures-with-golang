package main

import (
	"data-structures/linkedlist"
	"data-structures/queu"
	stack2 "data-structures/stack"
)

func main(){
	list := linkedlist.LinkedList{}
	node1 := linkedlist.Node{Data: 32}
	node2 := linkedlist.Node{Data: 5}
	node3 := linkedlist.Node{Data: 65}
	node4 := linkedlist.Node{Data: 18}
	list.AddLast(&node1)
	list.AddLast(&node2)
	list.AddLast(&node3)
	list.AddFirst(&node4)
	list.AddAfter(5,1)
	list.Remove(18)
	//list.Print()
	//fmt.Print("Size: ",list.Size())
	// ------- STACK ---------
	stack := stack2.Stack{}
	stack.Push(32)
	stack.Push(5)
	stack.Push(10)
	stack.Push(87)
	stack.Pop()
	stack.Print()
	// ----- QUEUE -----
	queue := queu.Queu{}
	queue.Enqueu(4)
	queue.Enqueu(88)
	queue.Enqueu(0)
	queue.Enqueu(12)
	queue.Print()

}
