// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type ConfigurationYaml struct {
	PortField  uint16      `yaml:"port"`
	SitesField []*SiteYaml `yaml:"sites"`
	IndexField bool        `yaml:"index"`
}

type SiteYaml struct {
	RefField         string          `yaml:"ref"`
	HostField        string          `yaml:"host"`
	DescriptionField string          `yaml:"description""`
	RedirectsField   []string        `yaml:"redirects"`
	LanguageField    string          `yaml:"language"`
	KeepLinksField   bool            `yaml:"keeplinks"`
	FaviconPathField string          `yaml:"faviconpath"`
	FrontProxyField  *FrontProxyYaml `yaml:"frontproxy"`
}

type FrontProxyYaml struct {
	ForceSSL bool   `yaml:"forcessl"`
	IPHeader string `yaml:"ipheader"`
}

func (config *ConfigurationYaml) Port() uint16 { return config.PortField }
func (config *ConfigurationYaml) Sites() []Site {
	sites := make([]Site, len(config.SitesField))

	for i, _ := range config.SitesField {
		sites[i] = config.SitesField[i]
	}

	return sites
}
func (config *ConfigurationYaml) Index() bool { return config.IndexField }

func (site *SiteYaml) Ref() string                 { return site.RefField }
func (site *SiteYaml) Host() string                { return site.HostField }
func (site *SiteYaml) Description() string         { return site.DescriptionField }
func (site *SiteYaml) Redirects() []string         { return site.RedirectsField }
func (site *SiteYaml) Language() string            { return site.LanguageField }
func (site *SiteYaml) KeepLinks() bool             { return site.KeepLinksField }
func (site *SiteYaml) FaviconPath() string         { return site.FaviconPathField }
func (site *SiteYaml) FrontProxy() *FrontProxyYaml { return site.FrontProxyField }
func (s *SiteYaml) GRef() string {
	if !strings.Contains(s.Ref(), "/") {
		return "view/" + s.Ref()
	} else {
		return s.Ref()
	}
}

func (site *SiteYaml) IPHeader() string {
	if site.FrontProxyField != nil {
		return site.FrontProxyField.IPHeader
	} else {
		return ""
	}
}

func LoadConfig(filename string) func() (Configuration, error) {
	return func() (Configuration, error) {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		c := &ConfigurationYaml{}
		err = yaml.Unmarshal(bytes, c)
		if err != nil {
			return nil, err
		}
		return c, nil
	}
}
