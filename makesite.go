package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
)

/*
If I cannot get this to work then use these:
"github.com/bregydoc/gtranslate"
 "github.com/fatih/color"
 translatedContent, err := gtranslate.TranslateWithParams(
  postContent,
  gtranslate.TranslationParams{
   From: "en",
   To:   "es",
  },
 )
 if err != nil {
  panic(err)
 }

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
