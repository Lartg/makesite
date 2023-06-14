package main

import (
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readTextFlag() {
	text := flag.String("text", "", "text to print")
	flag.Parse()
	fmt.Println(*text)
}

func readTextFileFlag() {
	file := flag.String("file", "", "textfile to be read")
	flag.Parse()
	text, err := os.ReadFile(*file)
	check(err)

	fmt.Println(string(text))

}

func main() {
	readTextFileFlag()
}
