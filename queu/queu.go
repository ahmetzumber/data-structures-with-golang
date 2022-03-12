package queu

import "fmt"

type Node struct {
	Data int
	NextNode *Node
}

type Queu struct {
	Last *Node
}

func (q *Queu) Enqueu(newData int) {
	newNode := &Node{ Data: newData}
	if q.IsEmpty() {
		fmt.Println("Empty list!!")
	}else {
		newNode.NextNode = q.Last
	}
	q.Last = newNode
}

func (q *Queu) Dequeu() (int, bool){
	if q.IsEmpty() {
		fmt.Println("Empty list!!")
		return 0, false
	}else if q.Last.NextNode == nil {
		singleNodeData := q.Last.Data
		q.Last = nil
		return singleNodeData, true
	}else {
		current := q.Last
		for {
			if current.NextNode.NextNode == nil {
				data := current.NextNode.Data
				current.NextNode = nil
				return data, true
			}
		}
	}
}

func (q *Queu) Print() {
	fmt.Println("--- QUEU ---")
	for q.Last != nil {
		fmt.Print(q.Last.Data)
		fmt.Print(" -> ")
		q.Last = q.Last.NextNode
	}
	fmt.Print("null")
}

func (q *Queu) IsEmpty() bool {
	return q.Last == nil
}