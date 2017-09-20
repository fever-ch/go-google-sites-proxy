// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package proxy

import (
	"strconv"
	"net/http"
	"sync/atomic"
	"unsafe"
	"github.com/fever-ch/go-google-sites-proxy/common/config"
)

type Context struct {
	configuration config.Configuration
	sites         map[string]*func(responseWriter http.ResponseWriter, request *http.Request)
}

func NewCheapProxy(port uint16) *SmartProxy {

	buildContext := func(configuration config.Configuration) *Context {

		context := Context{
			configuration,
			make(map[string]*func(responseWriter http.ResponseWriter, request *http.Request))}

		for _, e := range configuration.Sites() {

			addRedirect := func(redirectedHost string, destHost string) *func(responseWriter http.ResponseWriter, request *http.Request) {
				prot := "http"
				if e.FrontProxy().ForceSSL {
					prot = "https"
				}

				redirectHandler := func(response http.ResponseWriter, req *http.Request) {

					response.WriteHeader(http.StatusMovedPermanently)
					response.Header().Add("Location", prot+"://"+destHost+"/"+req.URL.Path)

					response.Write([]byte(strconv.Itoa(http.StatusMovedPermanently) + " Moved permanently"))
				}
				return &redirectHandler
			}

			context.sites[e.Host()] = GetSiteHandler(e)
			for _, f := range e.Redirects() {
				context.sites[f] = addRedirect(f, e.Host())
			}
		}
		return &context
	}

	var context unsafe.Pointer

	handler := func(responseWriter http.ResponseWriter, request *http.Request) {
		ctx := (*Context)(atomic.LoadPointer(&context))
		siteHandler := ctx.sites[request.Host]
		if siteHandler != nil {
			(*siteHandler)(responseWriter, request)
		} else if ctx.configuration.Index() {
			(*getIndex(ctx.configuration))(responseWriter, request)
		}
	}

	return &SmartProxy{
		SetConfiguration: func(configuration config.Configuration) {
			atomic.StorePointer(&context, unsafe.Pointer(buildContext(configuration)))
		},
		Port: func() uint16 { return port },
		Start: func() error {
			http.HandleFunc("/", handler)
			return http.ListenAndServe(":"+strconv.Itoa(int(port)), nil)
		},
	}
}

type SmartProxy struct {
	Start            func() error
	SetConfiguration func(config.Configuration)
	Port             func() uint16
}

type SmartProxyConfig struct {
	site      string
	port      int
	gz        bool
	keepLinks bool
	language  string
}
