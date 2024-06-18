// package main

// import (
// 	"fmt"
// )

// func main() {
// 	note := fmt.Sprint("Functions in Go: \n 1) Must define the type for params and return (static remember) \n 2) func <funcName>(param <typeOfParam>, ...)(return types){}")
// 	fmt.Println("note: ", note)

// 	x := 100
// 	y := 200
// 	fmt.Println("x: ", x)
// 	fmt.Println("y: ", y)
// 	newNum, message := add(x, y)
// 	fmt.Println("newNum and message ", newNum, message)

// 	note2 := fmt.Sprint("\n\nFunctions have types and they can be used as parameters: \nfunc callFunc(callable func(int) int) int {return callable(10)}")
// 	fmt.Println(note2)

// 	fmt.Println("\nThis defined a function that accepts another function as a param and returns an int")

// 	note3 := fmt.Sprint("\n\nThe function that can be passed is this: \nfunc doubleNumber(num int) int {return 2 * num}")
// 	fmt.Println(note3)

// 	value := callFunc(doubleNumber)
// 	fmt.Println("The value from this: value := callFunc(doubleNumber) is = ", value)

// 	fmt.Println("\n\nLooking at the types of the functions: ")
// 	fmt.Printf("1) type of add(): %T", add)
// 	fmt.Printf("\n2) type of doubleNumber(): %T", doubleNumber)
// 	fmt.Printf("\n3) type of callFunc(): %T", callFunc)

// 	fmt.Println("\n\nAs seen above, callfunc has been defined \nas a function that accepts another function as a param and returns an int")

// 	note4 := fmt.Sprint("\n\nAnonymous functions are un-named functions and are allowed in Go! \nThis is can be done by passing a function much like an arrow function in JS")
// 	fmt.Println(note4)
// 	fmt.Println("\nanotherVal := callFunc(func(num int) int { \n return n + 2 \n})")

// 	anotherVal := callFunc(func(num int) int {
// 		return num + 2
// 	})
// 	fmt.Println("\nIf anotherVal: ", anotherVal)

// 	fmt.Println("\n\nWe can also assign the function to a variable as such: \nanFunc := func(num int) int { \n return num + 1 \n}")

// 	anFunc := func(num int) int {
// 		return num + 1
// 	}
// 	fmt.Println("\nNow what is the result of: res := anFunc(5)")
// 	res := anFunc(5)
// 	fmt.Println("res: ", res)

// 	fmt.Println("\n\nCan also return functions: ")

// 	fmt.Println("Using the getFunc function to create the function \nThis returned function can be used to \nappend the string to the parameter passed to getFunc()")

// 	f1 := getFunc("Hey")
// 	newFuncValue := f1("There")
// 	newFuncValue2 := f1("Ben")

// 	fmt.Println(newFuncValue, newFuncValue2)

// 	fmt.Printf("\n\nLets look at the getFunc function type: %T", getFunc)
// 	fmt.Println("\nAs one can see functions can be nested through returns")

// 	fmt.Println("\n\nFunctions can also be set to accept a variable amount of arguements.")
// 	fmt.Println("This can be seen here: func sum(nums ...int) int {} \nWhere the spread allows for any number of parameters")
// 	fmt.Println("\nThese are reffered to as a variadic function.")

// 	fmt.Println("This is going to be a slice seen above (nums)")
// 	fmt.Println("Therefore we can get the sum as such: \nsl := 0 \nfor _, value := range nums {")
// }

// // Accepting a variable amount of arguements
// func sum(nums ...int) int {
// 	sl := 0

// 	for _, value := range nums {
// 		sl += value
// 	}
// 	return sl
// }

// func getFunc(str string) func(string) string {
// 	return func(str2 string) string {
// 		return str + str2
// 	}
// }

// func doubleNumber(num int) int {
// 	return 2 * num
// }

// // Can accept functions as parameter in functions
// // This is accepting a function that accepts one integer parameter and returns one int value
// func callFunc(callable func(int) int) int {
// 	return callable(10)
// }

// func add(num1 int, num2 int) (int, string) {
// 	return num1 + num2, "Hey there"
// }
