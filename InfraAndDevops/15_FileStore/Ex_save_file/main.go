package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("./tmp/123.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := f.Write([]byte("writing some data into a file"))
	if err != nil {
		panic(err)
	}
	fmt.Println("wrote %d bytes", n)
}
