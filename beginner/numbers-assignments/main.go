package main

import (
	"fmt"
)

func main() {
	// int (Integer) represents whole numbers without decimals.
	// When dividing integers, the decimal part is truncated, meaning deleted.
	// For example 15/2 = 7. The .5 part is truncated.
	printIntegersExample()

	// float64 (Floating Point) represents numbers with decimals.
	// Slightly slower than int, so it should be avoided when decimals are not needed.
	// Not always exact (binary representation).
	// For example 0.1 + 0.2 may return 0.30000000000000004 instead of 0.3.
	printFloatExample()
}

func printIntegersExample() {
	fmt.Printf("printIntegersExample function called\n")
	var a int // Declare variable 'a' of type int
	var b int // Declare variable 'b' of type int

	// These types of assignments are called simple assignments
	a = 5  // Assignment of value 5 to variable a
	b = 10 // Assignment of value 10 to variable b

	// Print values and types of a and b
	// Type is int since the variables were declared as int
	fmt.Printf("Value of a is %d and type is %T\n", a, a)
	fmt.Printf("Value of b is %d and type is %T\n", b, b)

	// This type of assignment is called short variable declaration and assignment
	// Declare and initialize variables in one line. By "initialize", it means to assign an initial value to the variable.
	// Gets its type from the assigned value and not from a previous declaration.
	mean := (a + b) / 2 // Calculates the mean of a and b which means the average value.

	fmt.Printf("Mean of a and b is %v and type is %T\n", mean, mean)
}

func printFloatExample() {
	fmt.Printf("printFloatExample function called\n")
	var a float64 // Declare variable 'a' of type float64
	var b float64 // Declare variable 'b' of type float64

	// These types of assignments are called simple assignments
	a = 5  // Assignment of value 5 to variable a
	b = 10 // Assignment of value 10 to variable b

	// Print values and types of a and b
	// Type is float64 since the variables were declared as float64
	fmt.Printf("Value of a is %f and type is %T\n", a, a)
	fmt.Printf("Value of b is %f and type is %T\n", b, b)

	// This type of assignment is called short variable declaration and assignment
	// Declare and initialize variables in one line. By "initialize", it means to assign an initial value to the variable.
	// Gets its type from the assigned value and not from a previous declaration.
	mean := (a + b) / 2 // Calculates the mean of a and b which means the average value.

	fmt.Printf("Mean of a and b is %v and type is %T\n", mean, mean)

	// Notes about float

	// float64 has a precision of 15–16 significant decimal digits
	// float32 has a precision of 7 significant decimal digits
	// Significant decimal digits are both the integer and fractional parts, excluding leading zeros.
	// It does not mean only digits after the decimal point.
	// float32 uses less memory, so is faster, but has less precision.

	// In most cases float64 should be used, but float32 should be used when memory or throughput
	// is more important than precision, for example in large arrays (millions) or GPU / ML workloads.

	// We could get deeper in the differences and examples, but knowing this is enough for almost all use cases.
}
