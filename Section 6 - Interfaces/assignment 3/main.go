package main

import (
	"fmt"
	"io"
	"os"
	//"io/ioutil"
)

func main() {
	// If using ioutil.ReadFile()
	// bs, err := ioutil.ReadFile(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(bs))

	// If using os.Open()
	fptr, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// Instead of creating a bytestring and reading the contents of
	// the file into it, we can probably use io.Copy as a shortcut!
	// bs := make([]byte, 32*1024)
	// fptr.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy() expects as its second argument a type that implements the io.Reader interface,
	// which expects types associated with it to have the Read() function. We can pass fptr into
	// io.Copy() as the 2nd argument since fptr implements the Reader interface. Indeed, fptr
	// (which is a *File type) has the Read() function defined.
	io.Copy(os.Stdout, fptr)
}