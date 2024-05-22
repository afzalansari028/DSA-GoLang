package main

import "fmt"

func main() {

	queue := Queue{}

	fmt.Println("IsEmpty::", queue.IsEmpty())
	fmt.Println("Size::", queue.Size())

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	queue.Dispay()

	fmt.Println("IsEmpty::", queue.IsEmpty())
	fmt.Println("Size::", queue.Size())

	fmt.Println("peek::", queue.Peek())

	queue.Dequeue()
	fmt.Println("peek::", queue.Peek())

}

var rear = -1

type Queue struct {
	Data [10]int
}

// size check
func (queue *Queue) Size() int {
	return len(queue.Data)
}

// empty check
func (queue *Queue) IsEmpty() bool {
	return rear == -1
}

// insertion of element
func (queue *Queue) Enqueue(val int) {
	if rear == queue.Size()-1 {
		fmt.Println("QUEUE IS FULL")
	}
	rear++
	queue.Data[rear] = val
}

// deletion of element
func (queue *Queue) Dequeue() int {
	if queue.IsEmpty() {
		fmt.Println("QUEUE IS EMPTY")
		return -1
	}
	val := queue.Data[0]

	for i := 0; i < rear; i++ {
		queue.Data[i] = queue.Data[i+1]
	}
	rear--
	return val
}

// see the front elememnt
func (queue *Queue) Peek() int {
	if queue.IsEmpty() {
		fmt.Println("QUEUE IS EMPTY DURING PEEK")
	}
	val := queue.Data[0]
	return val
}

//Display
func (queue *Queue) Dispay() {
	if queue.IsEmpty() {
		fmt.Println("cant display, queue is empty")
	}
	for i := 0; i <= rear; i++ {
		fmt.Print(queue.Data[i], " ")
	}
	fmt.Println()
}
