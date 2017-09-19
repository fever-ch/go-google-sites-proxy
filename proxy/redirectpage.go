// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package proxy

import (
	"html/template"
	"net/http"
	"github.com/fever-ch/go-google-sites-proxy/common"
	"bytes"
	"github.com/fever-ch/go-google-sites-proxy/blob"
)

const movedPageTmpl = `<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
<html><head>
<Title>Moved</Title>
</head><body>
<h1>Moved</h1>
<p>This page has moved to <a href="{{.Dest}}"">{{.Dest}}</a>.</p>
<hr>
</body></html>`

func movedPage(code int, destRoot string) func(request *http.Request) *Page {
	type MovedPage struct {
		Dest        string
		ProgramInfo common.ProgramInfoStruct
	}

	t := template.New("")
	tt, _ := t.Parse(movedPageTmpl)

	return func(request *http.Request) *Page {
		fullDest := destRoot + request.URL.Path
		hdrs := make(map[string]((string)))
		hdrs["Location"] = fullDest

		var doc bytes.Buffer

		tt.Execute(&doc, MovedPage{fullDest, common.ProgramInfo})

		return &Page{code, hdrs, blob.NewRawBlob([]byte (doc.String())), true}
	}
}
