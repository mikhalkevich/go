package main

import (
	"errors"
	"fmt" 
)

// Custom error message with errors.New()
func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this is a custom error message")
	}
	return nil
}
 

func main() {
	err := check(0, 10)
	if err == nil {
		fmt.Println("check() ended normally!")
	} else {
		fmt.Println(err)
	}
 
}
