// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package config

type Site struct {
	Ref         string `yaml:"ref"`
	Host        string `yaml:"host"`
	Description string `yaml:description`
	Redirects   []string `yaml:"redirects"`
	Language    string `yaml:"language"`
	KeepLinks   bool `yaml:"keeplinks"`
	IPHeader string `yaml:ipheader`
	FaviconPath string `yaml:faviconpath`
}

type Configuration struct {
	Port  uint16 `yaml:"port"`
	Sites []*Site `yaml:"sites"`
	Index bool `yaml:index`
}
