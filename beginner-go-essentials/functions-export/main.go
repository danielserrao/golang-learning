package main

import (
	"fmt"
	"golang-learning/functions-export/mathutils" // import your custom package
)

func main() {
	result := mathutils.Add(9, 4) // ✅ Works, Add is exported
	fmt.Println(result)

	/*
		 	result2 := mathutils.subtract(9, 4) // ❌ Compile error, subtract is unexported
			fmt.Println(result2)
	*/
}
