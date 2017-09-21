package proxy

import (
	"github.com/fever-ch/go-google-sites-proxy/common"
	"github.com/fever-ch/go-google-sites-proxy/common/config"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
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
    {{range .Sites}}<li><a href="http://{{.HostField}}">{{if .DescriptionField}}{{.DescriptionField}}{{else}}{{.HostField}}{{end}}</a></li>{{end}}
    </ul>
    <br>
    <br>
    2017 Fever.ch - <a href="{{.ProgramInfo.ProjectURL}}">{{.ProgramInfo.Fullname}}</a>
  </body>
</html>`

type IndexStruct struct {
	ProgramInfo common.ProgramInfoStruct
	Sites       []config.Site
}

func getIndex(configuration config.Configuration) *func(responseWriter http.ResponseWriter, request *http.Request) {
	f := func(responseWriter http.ResponseWriter, req *http.Request) {
		t := template.New("")
		tt, _ := t.Parse(tmpl)
		err := tt.Execute(responseWriter, IndexStruct{common.ProgramInfo, configuration.Sites()})
		if err != nil {
			log.Warning("Problem rendering the template", err)
		}
	}
	return &f
}
