package main

import (
	"fmt"
)

/*
Algorithm to find the permutations of a given array by creating a tree from it.
For each element of the array it creates the branch starting from that and reducing the array step by step.
In the end, each branch is a different permutation.

Works decently with an array max size of 10 elements.
Still sucks anyway, if u're looking for efficiency look elsewhere
*/


type Node struct {
	Parent *Node
	Value int
	Children []*Node
}

// Inner function to reduce the array size while calculating alla the permutations
func removeFromArray (index int, array []int) []int {
	outArray := []int{}
	for n := range array {
		if n == index {
			continue
		}
		outArray = append(outArray, array[n])
	}
	return outArray
}

// Inner permutator function
func permutator (parent *Node, value int, arrayOfNumbers []int, collector *[]string) *Node {
	node := Node {
		Parent: parent,
		Value: value,
		Children: []*Node{},
	}

	if len(arrayOfNumbers) == 0 {
		currentNode := node
		currentString := ""
		for currentNode.Parent != nil {
			currentString += fmt.Sprint(currentNode.Value)
			currentNode = *currentNode.Parent
		}
		currentString += fmt.Sprint(currentNode.Value)
		*collector = append(*collector, currentString)
	}

	for r := range arrayOfNumbers {
		node.Children = append(node.Children, permutator(&node, arrayOfNumbers[r],removeFromArray(r,arrayOfNumbers), collector))
	}

	return &node
}

// Outer permutator function
// Give an array and a collector (must be a pointer to a string array)
func Permutator (array []int, collector *[]string) {
	for a := range array {
		permutator(nil, array[a], removeFromArray(a, array), collector)
	}
}

func main () {
	array := []int{0,1,2,3,4,5,6, 7, 8, 9}
	collector := []string{}
	Permutator(array, &collector)
	fmt.Println(collector)
}
