package main

import "data-structures/linkedlist"

func main(){
	list := linkedlist.LinkedList{}
	node1 := linkedlist.Node{Data: 32}
	node2 := linkedlist.Node{Data: 5}
	node3 := linkedlist.Node{Data: 65}
	node4 := linkedlist.Node{Data: 18}
	list.AddLast(&node1)
	list.AddLast(&node2)
	list.AddLast(&node3)
	list.AddLast(&node4)
	list.Print()
}