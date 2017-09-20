// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package config


type Configuration interface {
	Sites() []Site
	Index() bool
	Port() uint16
}


type Site interface {
	Ref() string
	Host() string
	Description() string
	Redirects() []string
	Language() string
	KeepLinks() bool
	FaviconPath() string
	FrontProxy() *FrontProxyYaml
	GRef() string
}




