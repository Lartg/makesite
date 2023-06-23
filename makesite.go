package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/bregydoc/gtranslate"
)

/*
If I cannot get this to work then use these:

 "github.com/fatih/color"


*/

type Content struct {
	ParagraphEN string
	ParagraphJA string
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

	// translate text to Japanese
	translatedContent, err := gtranslate.TranslateWithParams(
		string(text),
		gtranslate.TranslationParams{
			From: "en",
			To:   "ja",
		},
	)
	if err != nil {
		panic(err)
	}
	// create the data structure for the template
	content := Content{
		ParagraphEN: string(text),
		ParagraphJA: translatedContent,
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
