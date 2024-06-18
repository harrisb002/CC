// package main

// /*//////////////////////////////////////////////////////
// This lesson Consists Looping
// */ //////////////////////////////////////////////////////

// import (
// 	"fmt"
// 	// "math"
// 	// "strconv"
// )

// func main() {
// 	// for idx := 0; idx < 10; idx++ {
// 	// 	fmt.Println(idx)
// 	// }

// 	// Dont have access to while loop, Must use for-loop
// 	// a := 10
// 	// for a < 10 {
// 	// 	fmt.Println("loop")
// 	// 	a++
// 	// }

// 	// // Looping through strings!!!
// 	// str := "Hey there cool dude"

// 	// // str[0] is accessible but is not assignable
// 	// // This is going to access the first byte that is inside the string
// 	// // Remember: byte = uint8
// 	// fmt.Printf("%T\n", str[0]) // As seen here
// 	// fmt.Println(str[0])        // Comes up with the integer representation of the charachter in the string

// 	// // Therefore to get the character must convert it to a string
// 	// fmt.Println("Heres the character, str[0] converted to a string.", string(str[0]))

// 	// // Strings as just slices of bytes

// 	/*
// 		However, Accessing all of the individual bytes of the strings is not usually desirable
// 		ASCII:
// 			American Standard code for Information Interchange.
// 			ASCII = 1 byte (or 8 bits) = 256 chars

// 		UTF-8:
// 		To encompass all the new chars now UTF-8 is the most used
// 		Uses 4 bytes
// 	*/

// 	/*
// 		Some characters like emojis may take up more than one byte and
// 		therefore a single index may not ecompass the entire charachter if
// 		one like an emoji takes more than one

// 		Must loop through the array of characters in a special way

// 	*/

// 	specialStr := "😁"
// 	str := "Hey there dude."

// 	fmt.Println("Will work for all charahcters that fit in one byte")
// 	for idx := 0; idx < len(str); idx++ {
// 		fmt.Printf("%c", str[idx])
// 	}
// 	fmt.Println("\n")
// 	fmt.Println("But Doesnt work for special characters that may take up too much for one byte")
// 	for idx := 0; idx < len(specialStr); idx++ {
// 		fmt.Printf("%c", specialStr[idx])
// 	}
// 	fmt.Printf("\nThis is the type when using regular for loop: %T", specialStr[0]) // This is an uint8 which = byte

// 	fmt.Println("\n")
// 	fmt.Println("Must use range syntax to get access to index and character in a specific position.")
// 	for _, char := range specialStr {
// 		fmt.Printf("%c", char)
// 		fmt.Printf("\nThis is the type when using range based for loop:%T", char) // This is an int32 which = rune
// 	}

// 	// Break and continue still work!!!

// cities := []string{"London", "Tokyo", "New York"}
// for index, city := range cities {
//   fmt.Printf("City %d is %s.\n", index + 1, city) // => City 1 is London. City 2 is Tokyo. City 3 is New York.
// }

// }
