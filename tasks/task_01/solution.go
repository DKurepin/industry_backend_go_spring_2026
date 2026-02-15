package main

import "fmt"

func greet(name string) string {
	if name == "" {
		return fmt.Sprint("Hello, World!")
	}
	return fmt.Sprintf("Hello, %s!", name)
}
