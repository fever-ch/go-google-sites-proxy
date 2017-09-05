// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package proxy

import (
	"text/template"
	"net/http"
	"go-google-sites-proxy/config"
)

const tmpl = `<!DOCTYPE html>
<html lang="en">
  <head>
  	<title>Go Google Sites Proxy</title>
    <meta charset="utf-8">
 </head>
  <body>
    <h1>Go Google Sites Proxy </h1>
	<ul>
    {{range .}}<li><a href="http://{{.Host}}">{{if .Description}}{{.Description}}{{else}}{{.Host}}{{end}}</a></li>{{end}}
    </ul>
    <br>
    <br>
    2017 Fever.ch - <a href="https://www.github.com/fever-ch/go-google-sites-proxy">Go Google Sites Proxy</a>
  </body>
</html>`

func getIndex(configuration config.Configuration) *func(responseWriter http.ResponseWriter, request *http.Request) {
	f := func(responseWriter http.ResponseWriter, req *http.Request) {
		t := template.New("")
		tt, _ := t.Parse(tmpl)
		tt.Execute(responseWriter, configuration.Sites)
	}
	return &f
}
