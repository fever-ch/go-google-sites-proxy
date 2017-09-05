// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package proxy

import (
	"strconv"
	"time"

	"io/ioutil"

	"net/http"
	"bytes"
	"regexp"
	"strings"
	"go-google-sites-proxy/blob"
	"go-google-sites-proxy/config"

)

// Get site handler for a given site
func GetSiteHandler(site *config.Site) *func(responseWriter http.ResponseWriter, request *http.Request) {
	var htmlRx, _ = regexp.Compile("text/html($|;.*)")

	googleSitePathRoot := "https://sites.google.com/view/" + site.Ref

	retrieve := func(url string) *http.Response {
		var netClient = &http.Client{
			Timeout: time.Second * 10,
		}
		req, _ := http.NewRequest("GET", googleSitePathRoot+url, nil)
		if site.Language != "" {
			req.Header.Set("Accept-Language", site.Language)
		}
		req.Header.Set("Accept-Encoding", "gzip")
		gsitesResponse, _ := netClient.Do(req)
		return gsitesResponse
	}

	// Render a page object
	renderPage := func(page* Page, responseWriter http.ResponseWriter, gzipSupport bool) {
		var buff []byte
		if page.OriginallyGziped && gzipSupport {
			responseWriter.Header().Set("Content-Encoding", "gzip")
			buff = page.Blob.Gzipped()
		} else {
			buff = page.Blob.Raw()
		}

		for key, value := range page.Headers {
			if value != "" {
				responseWriter.Header().Set(key, value)
			}
		}

		responseWriter.WriteHeader(page.Code)
		responseWriter.Header().Set("Content-Length", strconv.Itoa(len(buff)))
		responseWriter.Write(buff)
	}

	// Convert response to page object
	respToPage := func(resp *http.Response) *Page {
		headers := make(map[string](string))

		selectedHeaders := []string{"Content-Type", "Date", "Expires"}
		for _, key := range selectedHeaders {
			value := resp.Header.Get(key)
			if value != "" {
				headers[key] = value
			}
		}
		headers["Server"] = resp.Header.Get("Server") + " + GGSP"

		gzipped := resp.Header.Get("Content-Encoding") == "gzip"

		var body blob.HybridRawGzipBlob
		b, _ := ioutil.ReadAll(ioutil.NopCloser(resp.Body))
		if gzipped {
			body = blob.NewGzippedBlob(b)
		} else {
			body = blob.NewRawBlob(b)
		}

		if !site.KeepLinks && htmlRx.MatchString(resp.Header.Get("Content-Type")) {
			body = blob.NewRawBlob(bytes.Replace(body.Raw(), []byte( "\"/view/"+site.Ref), []byte( "\""), -1))
		}

		return &Page{resp.StatusCode, headers, body, gzipped}
	}

	// The actual handler that will get the request for this site
	handleRequest := func(responseWriter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			gsitesResponse := retrieve(request.URL.Path)
			page:= respToPage(gsitesResponse)
			renderPage(page, responseWriter, strings.Contains(request.Header.Get("Content-Encoding"), "gzip"))

		default:
		}
	}

	return &handleRequest
}
