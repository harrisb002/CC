// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("* Can use the var syntax  but is a bit verbose: \nvar mp map[string]int = map[string]int{\"a\": 1, \"b\": 5}")
// 	var varMap map[string]int = map[string]int{"a": 1, "b": 5}

// 	fmt.Println("Here is the varMaps values:")
// 	println(varMap["a"], varMap["b"])

// 	fmt.Println("\n* Can also just use the Imp. Operator: \nmp := map[string]int{\"c\": 2, \"d\": 4}")
// 	mp := map[string]int{"c": 2, "d": 4}
// 	fmt.Println("\n\nHere is the mp values:")
// 	fmt.Println(mp)
// 	fmt.Println("\nOther values are the default for the type. \nHere is mp[\"z\"]")
// 	fmt.Println(mp["z"])

// 	fmt.Println("\n* Here is using the make function: \nmakeMap := make(map[int]int)")
// 	makeMap := make(map[int]int)
// 	fmt.Println("This created an empty map: ", makeMap)

// 	fmt.Println("\n\n* Here is a more complicated map \ncompMap := map[string][]int{\"a\": {1, 2, 3, 4, 5}}")
// 	compMap := map[string][]int{"a": {1, 2, 3, 4, 5}}
// 	fmt.Println("\nThis created a map with strings as keys and int slices for the value:")
// 	fmt.Println(compMap)

// 	fmt.Println("\n\n* Can also change values using the same syntax to access a key: \ncompMap[\"a\"] = []int{100, 200, 300}")
// 	compMap["a"] = []int{100, 200, 300}
// 	fmt.Println("compMap is now: ", compMap)

// 	fmt.Println("\n\n* Can also add values using the same syntax to access a key: \ncompMap[\"b\"] = []int{100, 200, 300}")
// 	compMap["b"] = []int{100, 200, 300}
// 	fmt.Println("compMap is now: ", compMap)

// 	fmt.Println("\n\n* Can also delete values using the delete functionality: \ndelete(compMap, \"a\")")
// 	delete(compMap, "a")
// 	fmt.Println("compMap is now: ", compMap)

// 	fmt.Println("\n\n* Can also determine if a value exists or not using: value, ok := compMap[\"b\"]")
// 	value, ok := compMap["b"]
// 	fmt.Println("value and ok is: ", value, ok)

// 	fmt.Println("\n\n* Quick example of using loops with map.")
// 	loopMap := map[uint]uint{}
// 	n := 100

// 	data := fmt.Sprint(" for number := 1; number <= n; number++ {\n  for divisor := 1; divisor <= 5; divisor++ {\n   if number%divisor == 0 {\n    loopMap[uint(divisor)]++ // This works cause the default values!!! Super usefull\n   }\n}")
// 	fmt.Println(data)

// 	for number := 1; number <= n; number++ {
// 		for divisor := 1; divisor <= 5; divisor++ {
// 			if number%divisor == 0 {
// 				loopMap[uint(divisor)]++ // This works cause the default values!!! Super usefull
// 			}
// 		}
// 	}

// 	fmt.Println("\n\n loopMap: ", loopMap)

// }
