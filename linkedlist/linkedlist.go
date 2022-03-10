package linkedlist

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
		// TODO
	}
}

func (l *LinkedList) print(){
	// TODO
}