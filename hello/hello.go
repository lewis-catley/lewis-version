package hello

import "fmt"

// PrintVimGo will print `vim-go`
func PrintVimGo() {
	fmt.Println("Hello World!")
}

// PrintString will print the string passed to it
func PrintString(str string) {
	fmt.Println(str)
}

// Add you can never imagine what this does
func Add(a, b int) int {
	return a + b
}

// Subtract no idea what this does
func Subtract(a, b int) int {
	return a - b
}
