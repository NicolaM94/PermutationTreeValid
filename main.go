package main

import (
	"errors"
	"fmt"
)


type Node struct {
	Parent *Node
	Value int
	Children []*Node
}

func RemoveFromArray (index int, array []int) []int {
	outArray := []int{}
	for n := range array {
		if n == index {
			continue
		}
		outArray = append(outArray, array[n])
	}
	return outArray
}

func FindIndex (value int, array []int) (int, error) {
	for i := range array {
		if array[i] == value {
			return i, nil
		}
	}
	return 0, errors.New("index not found")
}


func Permutator (parent *Node, value int, arrayOfNumbers []int) *Node {
	node := Node {
		Parent: parent,
		Value: value,
		Children: []*Node{},
	}

	if len(arrayOfNumbers) == 0 {
		currentNode := node
		for currentNode.Parent != nil {
			fmt.Print(currentNode.Value)
			currentNode = *currentNode.Parent
		}
		fmt.Print(currentNode.Value)
		fmt.Println()
	}

	for r := range arrayOfNumbers {
		node.Children = append(node.Children, Permutator(&node, arrayOfNumbers[r],RemoveFromArray(r,arrayOfNumbers)))
	}

	return &node
}



func main () {
	array := []int{0,1,2,3}
	for a := range array {
		Permutator(nil, array[a], RemoveFromArray(a,array))
	}
}
