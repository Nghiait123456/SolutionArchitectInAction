package main

import (
	"fmt"
	"time"
)

func main() {
	doSomething()
	time.Sleep(100 * time.Second)
}

func doSomething() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	fmt.Println("Before calling z()")
	z()
	fmt.Println("After calling z(), will never get here.")
}

func z() {
	panic("Panic message from z()")
}
