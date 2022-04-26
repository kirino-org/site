package main

import (
	"html/template"
	"io"
	"os"

	blackfriday "github.com/russross/blackfriday/v2"
)

func main() {
	tmpl, _ := template.New("").ParseFiles("./views/index.gohtml")

	f, err := os.Open("./index.md")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	raw, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	mdHtml := blackfriday.Run(raw, blackfriday.WithExtensions(blackfriday.CommonExtensions), blackfriday.WithExtensions(blackfriday.Footnotes))

	fWr, err := os.Create("./docs/index.html")
	if err != nil {
		panic(err)
	}
	defer fWr.Close()

	tmpl.ExecuteTemplate(fWr, "index.gohtml",
		template.HTML(
			string(
				mdHtml,
			),
		),
	)
}
