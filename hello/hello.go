package hello

import "fmt"

// SomethingCool returns a string just pass it a name
func SomethingCool(name string) string {
	return fmt.Sprintf("%s is an awesome person, have a great day!", name)
}
