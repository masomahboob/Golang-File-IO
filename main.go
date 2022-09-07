package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// writeFile("hello.txt")
	// readFile("hello.txt")
	// byteWriteFile("byte_file.txt")
	byteReadFile("byte_file.txt")
}

func writeFile(filename string) {

	// create a file
	file, err := os.Create(filename)

	// if any error occured, log the error
	if err != nil {
		log.Fatalln(err)
	}

	// write to a file and capture any error
	if _, err := file.WriteString("Hello, world!\n"); err != nil {
		log.Fatalln(err)
	}
}

func readFile(filename string) {

	// open a file
	file, err := os.Open(filename)

	// capture any error and log the error
	if err != nil {
		log.Fatalln(err)
	}

	// read file in chuncks with buffer

	// create a NewScanner to scan the file
	scanner := bufio.NewScanner(file)

	// let's get the text from the scanner
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// writing and reading with slice of bytes

func byteWriteFile(filename string) {

	// create the file
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalln(err)
	}

	if _, err := file.Write([]byte("Hello, world! This is a byte of strings.")); err != nil {
		log.Fatalln(err)
	}
}

func byteReadFile(filename string) {

	//
	data := make([]byte, 100)

	// open the file
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln(err)
	}

	// `file.Read()` function will read to the length of data and assign it to data
	if _, err := file.Read(data); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))
}
