package skeleton

import (
	"github.com/Puchungualotsqui/goplate/internal"
)

var TailwindSkeleton = []internal.FileTemplate{
	{Path: "web/static/css", IsDir: true},
	{Path: "web/static/css/input.css", Content: `
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
	`, IsDir: false},
	{Path: "web/static/css/tailwind.config.js", Content: `
	/** @type {import('tailwindcss').Config} */
	module.exports = {
	  content: [
	    "../../**/*.templ",
	    "./static/**/*.html",
	  ],
	  theme: {
	    extend: {},
	  },
	};
	`, IsDir: false},
	{Path: "web/templates/hello.templ", Content: `
	package templates

	templ Hello() {
	<html>
		<head>
			<title>Hello World</title>
			<link href="/static/css/output.css" rel="stylesheet">
		</head>
		<body>
			<h1 class="text-3xl font-bold text-blue-600">
				Hello from Goplate and Tailwind 🚀
			</h1>
		</body>
	</html>
	}
	`, IsDir: false},
}
