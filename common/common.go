// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package common

import (
	"strings"
)

func (s *Site) GRef() string {
	if !strings.Contains(s.Ref, "/") {
		return "view/" + s.Ref
	} else {
		return s.Ref
	}
}

type Configuration struct {
	portField  uint16  `yaml:"port"`
	sitesField []*Site `yaml:"sites"`
	indexField bool    `yaml:index`
}

func (config *Configuration) Port() uint16 {
	return config.portField
}

func (config *Configuration) Sites() []*Site { return config.sitesField }

func (config *Configuration) Index() bool { return config.indexField }

type Site struct {
	Ref         string   `yaml:"ref"`
	Host        string   `yaml:"host"`
	Description string   `yaml:description`
	Redirects   []string `yaml:"redirects"`
	Language    string   `yaml:"language"`
	KeepLinks   bool     `yaml:"keeplinks"`
	FaviconPath string   `yaml:faviconpath`
	FrontProxy  *FrontProxy `yaml:frontproxy`
}

func (s *Configuration) PortA() {}

type FrontProxy struct {
	ForceSSL bool     `yaml:forcessl`
	IPHeader string   `yaml:ipheader`
}
