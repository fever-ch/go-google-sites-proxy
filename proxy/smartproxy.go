// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package proxy

import (
	"strconv"
	"go-google-sites-proxy/config"
	"net/http"
)

type Context struct {
	configuration config.Configuration
	sites         map[string]*func(responseWriter http.ResponseWriter, request *http.Request)
}

func NewCheapProxy(configuration config.Configuration) *SmartProxy {

	buildContext := func(configuration config.Configuration) Context {
		context := Context{
			configuration,
			make(map[string]*func(responseWriter http.ResponseWriter, request *http.Request))}

		for _, e := range configuration.Sites {
			context.sites[e.Host] = GetSiteHandler(e)
			for _, f := range e.Redirects {
				addRedirect(f, e.Host)
			}
		}
		return context
	}

	context := buildContext(configuration)

	handler := func(responseWriter http.ResponseWriter, request *http.Request) {
		siteHandler := context.sites[request.Host]
		if siteHandler != nil {
			(*siteHandler)(responseWriter, request)
		} else if context.configuration.Index {
			(*getIndex(context.configuration))(responseWriter, request)
		}
	}

	return &SmartProxy{
		Start: func() error {
			http.HandleFunc("/", handler)
			err := http.ListenAndServe(":"+strconv.Itoa(context.configuration.Port), nil)
			if err != nil {
				return err
			}
			select {} // wait forever
		},
	}
}

type SmartProxy struct {
	Start func() error
}

type SmartProxyConfig struct {
	site      string
	port      int
	gz        bool
	keepLinks bool
	language  string
}
