// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package proxy

import (
	"strconv"
	"go-google-sites-proxy/config"
	"net/http"
)

func NewCheapProxy(configuration config.Configuration) *SmartProxy {

	sites := make(map[string]*func(responseWriter http.ResponseWriter, request *http.Request))

	addSite := func(site config.Site) {
		sites[site.Host] = GetSiteHandler(&site)
	}

	for _, e := range configuration.Sites {
		addSite(e)
		for _, f := range e.Redirects {
			addRedirect(f, e.Host)
		}
	}

	handler := func(responseWriter http.ResponseWriter, request *http.Request) {
		siteHandler := sites[request.Host]
		if siteHandler != nil {
			(*siteHandler)(responseWriter, request)
		} else if configuration.Index {
			(*getIndex(configuration))(responseWriter, request)
		}
	}

	return &SmartProxy{
		Start: func() error {
			http.HandleFunc("/", handler)
			err := http.ListenAndServe(":"+strconv.Itoa(configuration.Port), nil)
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