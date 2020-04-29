package hello

import "fmt"

// SomethingCool returns a string just pass it a name
func SomethingCool(name string) string {
	return fmt.Sprintf("%s is an awesome person, have a great day!", name)
}

// Version2 prints a version 2 statement
func Version2() {
	fmt.Println("This is version 2, boy we're releasing these fast")
}
