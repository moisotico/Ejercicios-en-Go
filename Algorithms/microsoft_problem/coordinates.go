//Definitions

package main

import (
	"fmt"
)

type point struct {
	X int
	Y int
}

type HeapNode struct{
	distance int
	Point point
}

func main() {
	// int k
	k := 2
	fmt.Println("Hello World!")
}
// There are various pointsk
// Compute all distance to the vertex

// Push to heap O(N*log(k))

// Perform k pop() O(log(k))

func closestk(Points []point, vertex point, k int) []point{
	// Usually represented as an array
	heap Heap
	heap.Init()

	for p := range Points {
		distance := calculateDist(p, vertex)
		if len(heap) == k{
			if heap.peek() > distance {
				heap.pop()
				heap.push(HeapNode(distance, p))
			}
		} else{
			heap.push(HeapNode(distance, p))
		}
	}

	ans := make([]point)
	for len(heap) != 0 {
		
		s := heap.pop()
		ans = append(ans, s.Point)
		
	}
	return ans
}
