package main

import (
	"os"
	"log"

	"github.com/smcho/simplego/gopjt/cms"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	for _, tmpl := range cms.Tmpl.Templates() {
		log.Printf("%v", *tmpl)
	}

	p := &cms.Page{
		Title:   "Hello, world!",
		Content: "This is the body of our webpage",
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}
