package skeleton

import (
	"github.com/Puchungualotsqui/goplate/internal"
)

var DefaultSkeleton = []internal.FileTemplate{
	{Path: "cmd", IsDir: true},
	{Path: "internal", IsDir: true},
	{Path: "web/templates", IsDir: true},
	{Path: "web/static/assets", IsDir: true},
	{Path: "web/templates/hello.templ", Content: `
	package templates

	templ Hello() {
	  <html>
	    <head>
	      <title>Hello World</title>
	    </head>
	    <body>
	      <h1>Hello from Templ 🚀</h1>
	    </body>
	  </html>
	}
	`, IsDir: false},
	{Path: "main.go", Content: `
	package main

	import (
			"log"
			"{{MODULE}}/web/templates"
			"net/http"
	)

	func main() {
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				err := templates.Hello().Render(r.Context(), w)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			})

			http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

			log.Println("🚀 Server running at http://localhost:3000")
			log.Fatal(http.ListenAndServe(":3000", nil))
	}
	`, IsDir: false},
}
