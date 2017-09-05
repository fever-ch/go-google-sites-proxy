// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package proxy

import (
	"strconv"
	"net/http"
)

func addRedirect(redirectedHost string, destHost string) *func(responseWriter http.ResponseWriter, request *http.Request) {
	redirectHandler := func(response http.ResponseWriter, req *http.Request) {
		response.WriteHeader(http.StatusMovedPermanently)
		response.Write([]byte(strconv.Itoa(http.StatusMovedPermanently) + " Moved permanently"))
	}
	return &redirectHandler
}
