// package main

// import (
// 	"fmt"
// 	"math"
// 	"strconv"
// )

// Stands for Format
/*
	Println = Print Line, Basically console.log in JS
	fmt.Println("Hey", 2, var)
	fmt.Printf stands for print formatted
	fmt.Printf("The type of var is: %T", var)
	fmt.Printf("The default value of var is: %v", var)
	fmt.Printf("The binary of var is: %b", var)
	fmt.Printf("The scientific notation of var is: %e", var)
	fmt.Printf("The floating rep. of var is: %f", var)
	fmt.Printf("The string rep. of var is: %s", var)
	fmt.Printf("The 2 decimal places of precision for float of var is: %.2f", var)
	fmt.Printf("Show the percent sign: %.2f%%", var)
	fmt.Printf("Escape chars work \" and here is var: ", var)

	Create a string but dont print it out yet us fmt.Sprintf or can be used to pass to a function
	y := fmt.Sprintf("Creating and storing the string in y", y)
	fmt.Println("Now here is the string stored in y: ", y)
*/

func main() {
	// fmt.Println("Hello there!")

	// // var/const - name of variable - type of variab;e
	// const myName string = "Ben!!!"
	// const myOthername string = "This is my other name"
	// var changedName string = "Changing the name"

	// fmt.Println(changedName)

	// // Implicit Operator
	// // Explicitly define vaiables when you dont want to give it a variable on declaration
	// // var x uint32 = 21
	// x := 21
	// fmt.Printf("%T", x)
	// fmt.Println()
	// fmt.Println(x)

	// // Type casting
	// y := uint64(x)
	// fmt.Printf("%T", y)
	// fmt.Println()
	// fmt.Println(y)

	// // Overflow can happen!
	// over := -9

	// // Trying to squeeze unsigned into signed!
	// overflow := uint(over)
	// fmt.Println(over)
	// fmt.Println(overflow)

	// Printing
	// str := fmt.Sprintf("I can store a string like this or a variable.")

	// str2 := 25.333
	// str3 := fmt.Sprintf("%f", str2)

	// fmt.Println("String stored in str: ", str)
	// fmt.Println("String stored in str2: ", str2)
	// fmt.Println("String stored in str3: ", str3)

	// myVar := 10.555
	// // fmt.Printf stands for print formatted
	// fmt.Printf("The type of var is: %T", myVar)
	// fmt.Println()
	// fmt.Printf("The default value of myVar is: %v", myVar)
	// fmt.Println()
	// myVar2 := 10
	// fmt.Printf("The binary of myVar2 is: %b", myVar2)
	// fmt.Println()
	// fmt.Printf("The scientific notation of myVar is: %e", myVar)
	// fmt.Println()
	// fmt.Printf("The floating rep. of myVar is: %f", myVar)
	// fmt.Println()
	// myVarStr := "Hey Im var str"
	// fmt.Printf("The string rep. of myVarStr is: %s", myVarStr)
	// fmt.Println()
	// fmt.Printf("The 2 decimal places of precision for float of myVar is: %.2f", myVar)
	// fmt.Println()
	// fmt.Printf("Show the percent sign: %.2f%%", myVar)
	// fmt.Println()
	// fmt.Printf("Escape chars work \" which is cool i guess")
	// fmt.Println()
	// fmt.Printf("I can manually print more endlines \n\n with backslash \\n \n\n")

	// Operators!
	/*
		+
		-
		*
		/
		--
		++
		%
		**
	*/
	/*
		Same type works
		x: = 7
		y := 9
		z := x + y

		Not the same type
		x: = uint8(7)
		y := 1000
		z := x + y

		Must make the operands the same type
		Important to convert x to y not the otherway around
		x: = uint8(7)
		y := 1000
		z := x + uint8(y)  // This is wrong as its going from a bigger to smaller value
		// Results in integer overflow as int type cant fit in uint8
		/////////////
		x: = uint8(7)
		y := 1000
		z := int(x) + y  // Always go smaller to bigger

		x: = 7
		y := 1000
		z := y / x  // Always same type out as the types put in
		//////////
		x: = 7
		y := 1000
		z := float64(y) / float64(x)  // Must use float conversion to get floating point value
	*/
	/*
		String Concat is the same
		x := "Hey"
		y := 2
		z := x + string(y) // This conversion gives the ASCII Rep. of the integer is as a string
		// Must use
		z := x + fmt.Sprint(y)
	*/
	/*
		Math Package!
		fmt.Println(math.Min(4, 5)) = 4
		fmt.Println(math.Max(4, 5)) = 5
		fmt.Println(math.Sqrt(4)) = 2
		fmt.Println(math.Pow(4, 2)) = 16
		fmt.Println(math.Round(4.255)) = 4
		fmt.Println(math.Ceil(4.23)) = 5
		fmt.Println(math.Floor(4.78)) = 4
	*/
	/*
		String to Integer Value Conversion
		x := "1234"
		y := int(x) // This will not work!
		// Must use a special package os "strconv"
		myStrToConvert := "1234"
		// myStrToConvert := "1234 This is an error"

		// Atoi takes a string value and returns an int (Basically ParseInt)
		// returns an error value as well
		myStrConversion, error := strconv.Atoi(myStrToConvert)
		fmt.Println(myStrConversion, error)

		// ParseInt()
		// Pass a string as well as the base and bit size(size to use to store value)
		x := "1234"
		y, err := strconv.ParseInt(x, 10, 64) // (var, base, size) use 0 for default 32 if it fits (x, 10, 0)
		fmt.Println(y, err)
	*/

	/*
		Comparison!!!
		x := uint(8)
		y := 10
		z := x < y // Mis Matched type! Must have same type
		z := int(x) < y

		// If using a int literal Go will auto default to a type based on the needs of the comp.
		x := uint(8)
		z := x < 9   // 9 Go will default to a uint8 for the comparison to work

		== is for Equivalence NO === needed as Go enforces this
		|| && ! => Are all implemented the same

		// Ifs are much the same!
		x := 2
		if x < 3 {
			fmt.Println("run")
		} else if x > 7 {
		} else {
			...
		}
		//////////////
		ifAmount := 20
		var ifState bool = false

		if ifState {
			fmt.Println("Hey there in the if!")
		} else if ifAmount > 10 {
			fmt.Println("Hey there in the else if!")
		} else {
			fmt.Println("Woah made it to else huh...?")
		}
			fmt.Println(z)
	*/

	/*
		Switch Statements!!!
		// Automatically breaks for you
		a := 1
		switch a {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("twi")
		default:
			fmt.Println("default")
		}

		// Naked switch statement!
		// redact the arguement to switch and use the comparison in each case
		// acts as an if - else if block (i.e. does not fall through)
		a := 3
		switch {
		a > 2:
			fmt.Println("greater than 0")
		a < 0:
			fmt.Println("less than zero")
		default:
			fmt.Println("default")
		}

		// To purposefully fall through then use the "fallthrough"
		a := 1
		switch {
		a < -1:
			fmt.Println("a is less than -1")
			fallthrough
		a < 0:
			fmt.Println("a is less than 0")
			fallthrough
		a < 1:
			fmt.Println("a is less than 1")
		default:
			fmt.Println("default")
		}

		// Can have multiple values in case
		a := "h"
		switch a {
		case "a", "b", "c":
			fmt.Println("This is an "a", "b", or "c")
		default:
			fmt.Println("default")
		}

		/ Can have in naked swicth doing this aas well
		a := "h"
		switch {
		case a == "a", a == "b", a =="c":
			fmt.Println("This is an "a", "b", or "c")
		default:
			fmt.Println("default")
		}

	*/

	fmt.Println(math.Min(4, 5))
	myStrToConvert := "1234"
	// myStrToConvert := "1234 This is an error"
	myStrConversion, error := strconv.Atoi(myStrToConvert)
	fmt.Println(myStrConversion, error)

}
