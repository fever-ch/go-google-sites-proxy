// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package proxy

import (
	"github.com/fever-ch/go-google-sites-proxy/common"
	"html/template"
	"net/http"
)

const errorPageTmpl = `<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
<html><head>
<Title>{{.Title}}</Title>
</head><body>
<h1>{{.Title}}</h1>
<p>{{.Message}}</p>
<hr>
<address>{{.ProgramInfo.Name}} - {{.ProgramInfo.Version}}</address>
</body></html>`

func errorPage(code int, title string, message string) func(responseWriter http.ResponseWriter, request *http.Request) {
	type ErrorPage struct {
		Title       string
		Message     string
		ProgramInfo common.ProgramInfoStruct
	}

	t := template.New("")
	tt, _ := t.Parse(errorPageTmpl)

	e := ErrorPage{title, message, common.ProgramInfo}

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(code)
		tt.Execute(responseWriter, e)
	}
}
