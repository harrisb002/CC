// package main

// import (
// 	"fmt"
// 	// "math"
// 	// "strconv"
// )

// func main() {

// 	// x := "\nThis is the Slices Guide!"
// 	// data := fmt.Sprint("\n- Arrays have a fixed length that is set when they are declared, /nwhile slices have a dynamic length that can change at runtime. \n- Slices have both a length and capacity \n- Is a view of an array  \n- Can increase and Decrease the capacity")
// 	// data2 := fmt.Sprint("\n-  \n-  \n-  ")
// 	// fmt.Println(x)
// 	// fmt.Println(data)
// 	// fmt.Println(data2)

// 	// // using slices
// 	// fmt.Println("\n\nUsing Slices: []<datatype><{data}>")

// 	// // Using underlying array that has been previously defined
// 	// underlyingArr := []int{1, 2, 3, 4, 5}
// 	// fmt.Println("\nUnderlying array: ", underlyingArr)
// 	// // Make a slice from the arr
// 	// sl := underlyingArr[:]
// 	// fmt.Println("\nSlice of the underlying array: ", sl)
// 	// fmt.Println("\n- Pointer starts at the start of the slice: ", sl[0])
// 	// fmt.Println("\n- Capacity of the slice is: ", cap(sl))
// 	// fmt.Println("\n- Length of the slice is: ", len(sl))

// 	// // Can increase the length of the slice of up to capacity
// 	// sl = sl[:4]
// 	// fmt.Println("\nSlice of the underlying array: ", sl)

// 	//// PlayGround ////

// 	//// Notes ////
// 	notes := fmt.Sprint("A Slice is a View of an array one can Increase and Decrease the capacity of. \nTo make a slice can either take a slice of an existing array \nOR\n make a slice which will auto. create a new array for you. ")
// 	fmt.Println(notes)

// 	subArr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	fmt.Println("This is my array: ", subArr)
// 	fmt.Printf("It has the type: %T ", subArr)

// 	notes2 := fmt.Sprint("\n\nI am going to make a slice of this array. \nI am taking [:3] so only 1-3 length and not at full cap.")
// 	fmt.Println(notes2)

// 	sl := subArr[:3]
// 	fmt.Println("\n\nHere is my slice w/ length and capacity: ", sl, len(sl), cap(sl))

// 	notes3 := fmt.Sprint("Now I am going to make a slice of the subarray again \nbut now I am taking [5:] so only 5-10 length and at full cap.")
// 	fmt.Println(notes3)

// 	sl2 := subArr[5:]
// 	fmt.Println("\n\nHere is my other slice w/ length and capacity: ", sl2, len(sl2), cap(sl2))

// 	notes4 := fmt.Sprint("Notice the distinction in length and Cap. \nRemember it also posses a pointer and thus effects the underlying array \nwhen passed to function as its pass by ref.")
// 	fmt.Println(notes4)

// 	notes5 := fmt.Sprint("\n\nTaking a slice of a slice allows one to expand the slice. \nThis is how we will make it dynamic.")
// 	fmt.Println(notes5)

// 	newSl := sl[:10]
// 	fmt.Println("By taking a new slice of the original sl[:3], newSl := sl[:10], \nI have now expanded it to fill capacity newSl: ", newSl, len(newSl), cap(newSl))

// 	fmt.Println("\n\nLets modify the slice and notice it changes the subArr. slice[0] = 100")

// 	sl[0] = 100
// 	fmt.Println("\n\nNow Notice both arrays are changed")
// 	fmt.Println("subArr & sl: ", subArr, sl)

// 	fmt.Println("\n\nWe can also just create slices without an underlying array!!!")
// 	slStr := []string{"Hey", "I am another sl"}
// 	fmt.Println("slStr := []string{\"Hey\", \"I am another sl}")
// 	fmt.Printf("Here is the slice type now: %T", slStr)

// 	fmt.Println("\n\nArrays will implicitly be created when defining this slice \nand will double in size as length reaches capacity")

// 	fmt.Println("\n\nAppending to slices can be done through the append() function.")
// 	fmt.Println("slStr = append(slStr, \"Ben\") fmt.Println(slStr, len(slStr), cap(slStr))")
// 	fmt.Println("Adding elements by slicing the slice and returning it. Must reassign the new version.")

// 	fmt.Println("\n\nNotice how the capacity doubles as length reaches it")
// 	for x := 0; x < 10; x++ {
// 		slStr = append(slStr, "Ben")
// 		fmt.Println(slStr, len(slStr), cap(slStr))
// 	}

// 	fmt.Println("\n\nThe new underlying array is being created for us due to how we defined the slice")
// 	fmt.Println("slStr := []string{\"Hello\", \"There\"}")

// 	fmt.Println("\n\nOther ways to slice is by using make function. Can also be used for other types as well")
// 	fmt.Println("Takes in the type, size, capacity(Optional)")
// 	fmt.Println("makeSl := make([]int, 10)")

// 	fmt.Println("\n\nIterating and passing slices to functions.")
// 	fmt.Println("When iterating can use the simple range syntax")
// 	fmt.Println("for i, value := range sl {fmt.Println(i, value)}")

// 	fmt.Println("\n\nLets make a new slice. newSlice := []string{\"Hey\", \"Im a\", \"newSlice\"}")
// 	newSlice := []string{"Hey", "Im a", "newSlice"}

// 	fmt.Println("\nPrinting the values in my slice defined newSlice using range-loop")
// 	for i, value := range newSlice {
// 		fmt.Println(i, value)
// 	}

// 	println("\n\nLets call a function and use the slice. \nLets try to modify it. slice[0] = 100")
// 	otherFunc(newSlice)

// 	fmt.Println("\nPrinting the values in my slice defined newSlice using range-loop again")
// 	for i, value := range newSlice {
// 		fmt.Println(i, value)
// 	}

// 	fmt.Println("\nAs noted it has changed the slice as it chnaged the underlying array that the slice is viewing")

// 	test := []int{}
// 	fmt.Printf("\n\nThe type is: %T", test)
// 	fmt.Println("\nThe test is: ", test)
// 	// fmt.Println("\nThe test[0] is: ", test[0]) // Error!!!
// 	// Does not initialize it for you

// 	// Make does
// 	test2 := make([]int, 5)
// 	fmt.Printf("\n\nThe type is: %T", test2)
// 	fmt.Println("\nThe test2 is: ", test2)
// 	fmt.Println("\nThe test2[0] is: ", test2[0])

// 	// Make does
// 	a := []int{1, 2, 3, 4, 5}
// 	fmt.Println(len(a), cap(a))
// 	a = append(a, 1000)
// 	fmt.Println(len(a), cap(a))

// 	s := []int{1, 2, 3, 4, 5}
// 	fmt.Println("Here is s: ", s)
// 	s = append(s[:2], s[3:]...)
// 	fmt.Println("Here is s: ", s)
// }

// func otherFunc(slice []string) {
// 	// If I change something here it will effect the underlying array
// 	slice[0] = "Changed this"
// 	fmt.Println("I am inside the function. \nChanging the underlying array with slice[0] = \"Changed this\". \nThis is a pass by reference becuase it is using pointers")
// }
