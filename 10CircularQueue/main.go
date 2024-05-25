package main

import "fmt"

func main() {

	fmt.Println("circular queue")

	q := Queue{}

	fmt.Println("isfull::", q.IsFull())
	fmt.Println("isEmpty::", q.IsEmpty())

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)
	q.Enqueue(60)
	q.Enqueue(70)
	q.Enqueue(80)
	q.Enqueue(90)
	q.Enqueue(100)

	fmt.Println("isfull::", q.IsFull())
	fmt.Println("isEmpty::", q.IsEmpty())

	q.Dequeue()
	fmt.Println("front::", q.Peek())
	q.Dispay()

}

var front = -1
var rear = -1

const size = 10

type Queue struct {
	Data [size]interface{}
}

// check isfull
func (queue *Queue) IsFull() bool {
	return (rear+1)%size == front
}

// check isempty
func (queue *Queue) IsEmpty() bool {
	return rear == -1 && front == -1
}

// add data in the queue
func (queue *Queue) Enqueue(val interface{}) {
	if queue.IsFull() {
		fmt.Println("Queue is full, cann't insert")
		return
	}

	if front == -1 {
		front = 0
	}

	rear = (rear + 1) % size
	queue.Data[rear] = val
}

// remove data from the queue
func (queue *Queue) Dequeue() interface{} {
	if queue.IsEmpty() {
		fmt.Println("Queue is empty, can't delete")
		return -1
	}

	removed := queue.Data[front]
	front++
	return removed
}

// see the front/peek element
func (queue *Queue) Peek() interface{} {
	if queue.IsEmpty() {
		fmt.Println("Queue is empty, can't show peek element")
		return -1
	}

	peek := queue.Data[front]
	return peek
}

//Display
func (queue *Queue) Dispay() {
	if queue.IsEmpty() {
		fmt.Println("can't display, queue is empty")
		return
	}
	for i := front; i <= rear; i++ {
		fmt.Print(queue.Data[i], " ")
	}
	fmt.Println()
}
