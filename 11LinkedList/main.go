package main

import "fmt"

func main() {
	fmt.Println("hiiii")

	ll := LinkedList{}
	ll.AddAtLast(100)
	ll.AddAtLast(200)
	ll.AddAtLast(300)
	ll.AddAtLast(400)

	ll.Display()
	// fmt.Println("length:", ll.Length())

	// head := ll.RemoveFirstNode()
	// fmt.Println("head--", head)

	// tail := ll.RemoveLastNode()
	// fmt.Println("tail--", tail)

	rv := ll.RemoveKthNode(3)
	fmt.Println("deleted data::", rv)

	// ll.RemoveLastNode()
	ll.Display()

}

type Node struct {
	Data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// --------------------------------------INSERTION------------------------------------------
// add an element/node at first
func (ll *LinkedList) AddAtFirst(data int) {
	//create a new node
	newNode := &Node{Data: data, next: nil}

	//check if no node is there
	if ll.head == nil {
		ll.head = newNode
		return
	}

	//point the new node to the head
	newNode.next = ll.head
	ll.head = newNode
}

// add an element/node at last
func (ll *LinkedList) AddAtLast(data int) {
	//creating a new node
	newNode := &Node{Data: data, next: nil}

	//check if no node is there
	if ll.head == nil {
		ll.head = newNode
		return
	}

	var temp *Node
	temp = ll.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func (ll *LinkedList) Display() {

	if ll.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	var temp *Node
	temp = ll.head

	for temp != nil {
		fmt.Print(temp.Data, "-->")
		temp = temp.next
	}
	fmt.Println("")
}

//length of LL
func (ll *LinkedList) Length() int {

	if ll.head == nil {
		return 0
	}

	var temp *Node
	temp = ll.head
	count := 0
	for temp != nil {
		count = count + 1
		temp = temp.next
	}
	return count
}

//-----------------------------------DELETION--------------------------------------------
// remove head of the LL
func (ll *LinkedList) RemoveFirstNode() int {

	if ll.head == nil {
		return 0
	}
	rv := ll.head.Data
	ll.head = ll.head.next
	return rv
}

//delete the last/tail element
func (ll *LinkedList) RemoveLastNode() int {
	if ll.head == nil || ll.head.next == nil {
		return 0
	}

	var temp *Node
	temp = ll.head

	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
	return temp.Data
}

//delete any element
func (ll *LinkedList) RemoveKthNode(k int) int {
	var rv int
	if k == 1 {
		rv = ll.head.Data
		ll.head = ll.head.next
		// return rv
	}

	var temp *Node
	temp = ll.head
	count := 1
	for temp != nil {
		// fmt.Println(temp.next.Data)
		count++
		if count == k { //k==3
			rv = temp.next.Data
			temp.next = temp.next.next
			break
		}
		temp = temp.next
	}
	return rv
}
