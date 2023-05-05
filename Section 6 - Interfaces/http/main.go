package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// This approach is too manual. We can make things easier by using io.Copy
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)  // Read response into our bs slice
	// fmt.Println(string(bs))

	// io.Copy allows us to easily copy the contents of resp.Body into std out.
	// 1st arg: dst Writer, 2nd ard: src Reader
	//io.Copy(os.Stdout, resp.Body)

	// Calling io.Copy() with our own type that implements the Writer interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// Creating our own type that implements the the Writer interface
type logWriter struct{}
func (logWriter) Write(bs []byte) (int, error) {
	// Note how since we're not gonna actually be using a logWriter
	// value in this function, we don't have to specify the name of
	// a parameter for it
	fmt.Println(string(bs))
	fmt.Printf("Wrote %d bytes\n", len(bs))
	return len(bs), nil
}