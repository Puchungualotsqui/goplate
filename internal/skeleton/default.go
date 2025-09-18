package skeleton

import (
	"goplate/internal"
)

var DefaultSkeleton = []internal.FileTemplate{
	{Path: "cmd", IsDir: true},
	{Path: "internal", IsDir: true},
	{Path: "web/templates", IsDir: true},
	{Path: "web/static/css", IsDir: true},
	{Path: "web/static/js", IsDir: true},
	{Path: "web/templates/home.templ", Content: `package templates

templ Home() {
  <html>
    <head><title>GoPlate App</title></head>
    <body><h1>Hello from GoPlate 🚀</h1></body>
  </html>
}`, IsDir: false},
}
