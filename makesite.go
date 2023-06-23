// package main

// import (
// 	"flag"
// 	"fmt"
// 	"html/template"
// 	"io/ioutil"
// 	"os"
// 	"strings"

// 	swearfilter "github.com/JoshuaDoes/gofuckyourself"
// )

// /*
// Use the Go swearfilter module

// DONE
// 1. Import swearfilter by uncommenting.

// DONE
// 2. declare swears in a flag and save as list of strings in {}

// DONE
// 3. use our current content var as message to be filtered

// DONE
// 4. filter message and return html as normal

// If I cannot get this to work then use these:
// "github.com/bregydoc/gtranslate"
//  "github.com/fatih/color"
//  translatedContent, err := gtranslate.TranslateWithParams(
//   postContent,
//   gtranslate.TranslationParams{
//    From: "en",
//    To:   "es",
//   },
//  )
//  if err != nil {
//   panic(err)
//  }

// */

// type Content struct {
// 	Header    string
// 	Paragraph string
// }

// func main() {
// 	// flag for list of swears
// 	swears := flag.String("swears", "", "Comma-separated list of strings")

// 	// parse the command-line flags
// 	flag.Parse()

// 	// split into a slice of strings
// 	swearList := strings.Split(*swears, ",")

// 	// INIT swear filter
// 	filter := swearfilter.New(false, false, false, false, false, swearList...)

// 	// flag to specify the file path
// 	filePath := flag.String("file", "", "Path to the input text file")
// 	flag.Parse()

// 	// check if file
// 	if *filePath == "" {
// 		panic("Please provide a file path using the -file flag")
// 	}

// 	// read the text from file
// 	text, err := ioutil.ReadFile(*filePath)
// 	if err != nil {
// 		panic("Error reading file should be a text file")
// 	}

// 	//----------------------------
// 	// Filter Text for swears
// 	swearFound, swearsFound, err := filter.Check(string(text))
// 	fmt.Println("Swear found: ", swearFound)
// 	fmt.Println("Swears tripped: ", swearsFound)
// 	fmt.Println("Error: ", err)
// 	//----------------------------

// 	// create the data structure for the template
// 	content := Content{
// 		Header:    "Template Header",
// 		Paragraph: string(text),
// 	}

// 	// read the template file
// 	templateFile, err := ioutil.ReadFile("template.tmpl")
// 	if err != nil {
// 		panic("failed to read template")
// 	}

// 	// computer understand file
// 	template, err := template.New("template.tmpl").Parse(string(templateFile))
// 	if err != nil {
// 		panic("failed to parse template")
// 	}
// 	// create the output file
// 	HTMLfile, err := os.Create("output.html")
// 	if err != nil {
// 		panic("failed to create output.html")
// 	}

// 	// write template to output.html
// 	err = template.Execute(HTMLfile, content)
// 	if err != nil {
// 		panic("failed to write content to output.html")
// 	}

// 	// close file done using
// 	HTMLfile.Close()
// }
