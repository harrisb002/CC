// package main

// import (
// 	"fmt"
// 	// "math"
// 	// "strconv"
// )

// func main() {
// 	x := "\nThis is the Array Guide!"
// 	data := fmt.Sprint("\n- Arrays are a Fixed-Size Data Structure that can hold variables on a type \n- Arrays are mutable and accessible \n- Cant add or Delete from it! \n- Basically just use slices lol")
// 	fmt.Println(x)
// 	fmt.Println(data)

// 	// Default values fill the array upon initialization
// 	// var arr [3]int // Cannot change the size of an array after initialization

// 	// Ways to init Arrays
// 	arr := [2]int{1, 2} // Curly Braces!

// 	// Make an array of size 2 that contains arrays of size two that contain integers
// 	nestedArr := [2][2]int{{1, 2}, {2, 3}}
// 	fmt.Println("\n\nMy Arr is: ", arr)
// 	fmt.Printf("My Arr type is: %T ", arr)
// 	fmt.Println("My nestedArr is: ", nestedArr)

// 	// Another way to int an array using the spread op
// 	// Automatically counts the size of the initialized array
// 	autoCount := [...][2]int{{1, 2}, {3, 4}}
// 	fmt.Println("\n\nMy autoCount is: ", autoCount)
// 	fmt.Printf("My autoCount type is: %T ", autoCount)

// 	// Accessing and Mutating an array
// 	mutateArr := [...][2]int{{1, 2}, {3, 4}, {5, 6}}
// 	fmt.Println("\n\nMutated array is: ", mutateArr)
// 	mutateArr[0] = [2]int{10, 20} // Must create the array to replace one in the mutateArr!!!
// 	fmt.Println("Mutated array is now: ", mutateArr)

// 	fmt.Println("\nPrint the outer arrays: ")
// 	for _, nested := range mutateArr {
// 		fmt.Println(nested)
// 	}

// 	fmt.Println("\nPrint the inner arrays: ")
// 	for _, nested := range mutateArr {
// 		for _, value := range nested {
// 			fmt.Println(value)
// 		}
// 	}
// }
