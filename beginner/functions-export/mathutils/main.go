package mathutils

// Exported function. This is visible outside the package
func Add(a int, b int) int {
	return a + b
}

// Unexported function. This is only visible inside the package
func subtract(a int, b int) int {
	return a - b
}
