package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func printNode(n *Node) {
	fmt.Printf("Value: %d ", n.Value)
	if n.Left != nil {
		fmt.Printf("Left: %d ", n.Left.Value)
	}
	if n.Right != nil {
		fmt.Printf("Right: %d ", n.Right.Value)
	}
	fmt.Println()

}

func main() {
	//node variables
	var val, indexLeft, indexRight int

	//Read file
	csvfile, err := os.Open("./file.in")
	if err != nil {
		log.Fatalln("Couldn't open the csv file ::", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)
	defer csvfile.Close()
	r.Comma = ' '

	csvData, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// nodes
	N := len(csvData)
	fmt.Println("Number of Nodes:", N)
	nodes := make([]Node, N)
	count := 0

	// Read file.in values and create nodes
	for _, rowValues := range csvData {
		val, _ = strconv.Atoi(rowValues[1])
		indexLeft, _ = strconv.Atoi(rowValues[2])
		indexRight, _ = strconv.Atoi(rowValues[3])

		nodes[count].Value = val
		if indexLeft >= 0 {
			nodes[count].Left = &nodes[indexLeft]
		}
		if indexRight >= 0 {
			nodes[count].Right = &nodes[indexRight]
		}
		count++
	}
	for _, node := range nodes {
		printNode(&node)
	}

}
