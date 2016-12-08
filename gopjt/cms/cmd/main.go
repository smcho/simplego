package main

import (
	"os"

	"github.com/smcho/simplego/gopjt/cms"
)

func main() {
	p := &cms.Page{
		Title:   "Hello, world!",
		Content: "This is the body of our webpage",
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}
