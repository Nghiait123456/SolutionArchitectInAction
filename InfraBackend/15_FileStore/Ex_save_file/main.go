package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	pr, pw := io.Pipe()

	// we need to wait for everything to be done
	wg := sync.WaitGroup{}
	wg.Add(2)

	// we get some file as input
	f, err := os.Open("./fruit.txt")
	if err != nil {
		log.Fatal(err)
	}

	// TeeReader gets the data from the file and also writes it to the PipeWriter
	tr := io.TeeReader(f, pw)

	go func() {
		defer wg.Done()
		defer pw.Close()

		// get data from the TeeReader, which feeds the PipeReader through the PipeWriter
		_, err := http.Post("https://example.com", "text/html", tr)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		// read from the PipeReader to stdout
		if _, err := io.Copy(os.Stdout, pr); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
