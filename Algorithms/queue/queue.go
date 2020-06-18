package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	slice []int
}

//Enqueue adds the integer to the back of the queue
func (q *Queue) Enqueue(i int) {
	q.slice = append(q.slice, i)
}

// Dequeue returns the first item on the queue and
// remove it
func (q *Queue) Dequeue() (int, error) {
	if len(q.slice) < 1 {
		err := errors.New("ERROR: Empty queue")
		return 0, err
	}
	ret := q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return ret, nil
}

func main() {
	var q *Queue = new(Queue)
	//q.Enqueue(123)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
}
