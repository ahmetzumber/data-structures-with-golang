package stack

import "fmt"

type Node struct {
	Data int
	NextNode *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(newData int) {
	newNode := &Node{Data: newData}
	if s.IsEmpty() {
		s.Top = newNode
	}else {
		newNode := &Node{Data: newData}
		newNode.NextNode = s.Top
		s.Top = newNode
	}
}

func (s *Stack) Pop() (int, bool) {
	if	s.IsEmpty() {
		fmt.Println("Empty stack!!")
	}
	temp := s.Top.Data
	s.Top = s.Top.NextNode
	return temp, true
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

func (s *Stack) Print()  {
	fmt.Println("--- STACK ---")
	if s.IsEmpty() {
		fmt.Println("Empty list!!")
	}else{
		temp := s.Top
		for temp != nil {
			fmt.Print("| ")
			fmt.Print(temp.Data)
			fmt.Println(" |")
			temp = temp.NextNode
		}
	}
	fmt.Println("Stack Top: ",s.Top.Data)
}