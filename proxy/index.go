// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package proxy

import (
	"text/template"
	"net/http"
	"go-google-sites-proxy/config"
	"go-google-sites-proxy/common"
	log "github.com/sirupsen/logrus"
)

const tmpl = `<!DOCTYPE html>
<html lang="en">
  <head>
  	<Title>{{.ProgramInfo.Name}}</Title>
    <meta charset="utf-8">
 </head>
  <body>
    <h1>{{.ProgramInfo.Fullname}}</h1>
	<ul>
    {{range .Sites}}<li><a href="http://{{.Host}}">{{if .Description}}{{.Description}}{{else}}{{.Host}}{{end}}</a></li>{{end}}
    </ul>
    <br>
    <br>
    2017 Fever.ch - <a href="{{.ProgramInfo.ProjectUrl}}">{{.ProgramInfo.Fullname}}</a>
  </body>
</html>`

type IndexStruct struct {
	ProgramInfo common.ProgramInfoStruct
	Sites             [] config.Site
}

func getIndex(configuration config.Configuration) *func(responseWriter http.ResponseWriter, request *http.Request) {
	f := func(responseWriter http.ResponseWriter, req *http.Request) {
		t := template.New("")
		tt, _ := t.Parse(tmpl)
		err := tt.Execute(responseWriter, IndexStruct{common.ProgramInfo, configuration.Sites})
		if err != nil {
			log.Warning("Problem rendering the template", err)
		}
	}
	return &f
}
