package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	// swearfilter "github.com/JoshuaDoes/gofuckyourself"
)

/*
Use the Go swearfilter module

1. Import swearfilter by uncommenting.

2. declare swears in a flag and save as list of strings in {}

3. use our current content var as message to be filtered

4. filter message and return html as normal

*/

type Content struct {
	Header    string
	Paragraph string
}

func main() {
	// flag to specify the file path
	filePath := flag.String("file", "", "Path to the input text file")
	flag.Parse()

	// check if file
	if *filePath == "" {
		panic("Please provide a file path using the -file flag")
	}

	// read the text from file
	text, err := ioutil.ReadFile(*filePath)
	if err != nil {
		panic("Error reading file should be a text file")
	}

	// create the data structure for the template
	content := Content{
		Header:    "Template Header",
		Paragraph: string(text),
	}

	// read the template file
	templateFile, err := ioutil.ReadFile("template.tmpl")
	if err != nil {
		panic("failed to read template")
	}

	// computer understand file
	template, err := template.New("template.tmpl").Parse(string(templateFile))
	if err != nil {
		panic("failed to parse template")
	}
	// create the output file
	HTMLfile, err := os.Create("output.html")
	if err != nil {
		panic("failed to create output.html")
	}

	// write template to output.html
	err = template.Execute(HTMLfile, content)
	if err != nil {
		panic("failed to write content to output.html")
	}

	// close file done using
	HTMLfile.Close()
}
