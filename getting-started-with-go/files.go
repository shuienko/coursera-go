package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const inputFilepath = "testfile01.txt"
const outputFilepath = "testfile02.txt"

func main() {
	dat, err := ioutil.ReadFile(inputFilepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dat), "|", dat)

	err = ioutil.WriteFile(outputFilepath, dat, 0644)
	if err != nil {
		fmt.Println(err)
	}

	// os package

	f1, _ := os.Open(inputFilepath)
	defer f1.Close()

	f1Stat, _ := f1.Stat()
	fmt.Println(f1.Name(), f1Stat.Mode(), f1Stat.ModTime())

	f1Content, _ := os.ReadFile(inputFilepath)
	fmt.Println(string(f1Content))

	// read two bytes from file
	arr := make([]byte, 2)
	f1.Read(arr)
	fmt.Println(string(arr))

	f2, _ := os.Create("outfile.txt")
	defer f2.Close()
	f2.Write([]byte("test\n"))
	f2.WriteString("Hi")

}
