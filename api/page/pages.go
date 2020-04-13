package page

import (
	"html/template"
	"io"
)

func Write404(w io.Writer, url string) {
	tmpl, _ := template.New("404").Parse(template404)

	tmpl.Execute(w, struct {
		URL string
	}{
		URL: url,
	})
}

var template404 = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
  	<meta name=viewport content="initial-scale=1, minimum-scale=1, width=device-width">
    <title>404 Resource not found</title>
</head>
<body>
	<p><b>404.</b>
  	<p>The requested URL <code>{{.URL}}</code> was not found on this server.
</body>
</html>`
