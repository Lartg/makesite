package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bregydoc/gtranslate"
)

type Content struct {
	ParagraphEN string
	ParagraphJA string
}

func main() {
	// flag for directory, default current dir
	dir := flag.String("dir", ".", "directory path")
	flag.Parse()

	// get absolute path of the directory
	absDir, err := filepath.Abs(*dir)
	if err != nil {
		panic(err)
	}

	// get all files
	files, err := ioutil.ReadDir(absDir)
	if err != nil {
		panic(err)
	}

	// find text files
	for _, file := range files {
		// read the text from file
		if filepath.Ext(file.Name()) == ".txt" {
			filePath := filepath.Join(absDir, file.Name())
			textByte, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", filePath, err)
				continue
			}
			text := string(textByte)

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

			// make unique new file name
			splitTextFileName := strings.Split(file.Name(), ".")
			formattedTextFileName := splitTextFileName[0]
			newFileName := fmt.Sprintf("%s.html", formattedTextFileName)
			// create the output file
			HTMLfile, err := os.Create(newFileName)
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
	}
}
