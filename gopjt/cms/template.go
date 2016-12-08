package cms

import "html/template"

// Tmpl is a reference to oall of our templates.
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

// Page is struct used to represent the content of a webpage.
type Page struct {
	Title   string
	Content string
}
