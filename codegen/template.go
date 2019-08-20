package codegen

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func parse(text string) *template.Template {
	tmpl := template.New("")
	return template.Must(tmpl.Parse(text))
}

func render(tmpl *template.Template, filename string, data interface{}) {
	var rendered bytes.Buffer
	tmpl.Execute(&rendered, data)

	fmted, err := format.Source(rendered.Bytes())
	if err != nil {
		log.Printf("rendered template:\n\n%s\n\n", rendered.Bytes())
		log.Panicf("problem generating %q: %s", filename, err)
	}

	err = ioutil.WriteFile(filename, fmted, os.ModePerm)
	if err != nil {
		log.Panicf("problem writing %q: %s", filename, err)
	}
}
