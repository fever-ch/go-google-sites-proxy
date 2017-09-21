package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

// ConfigurationYaml is a structure that will be automatically populated with the value of the configuration
type ConfigurationYaml struct {
	PortField  uint16      `yaml:"port"`
	SitesField []*SiteYaml `yaml:"sites"`
	IndexField bool        `yaml:"index"`
}

// SiteYaml is a structure that will be automatically populated with the value of the configuration
type SiteYaml struct {
	RefField         string          `yaml:"ref"`
	HostField        string          `yaml:"host"`
	DescriptionField string          `yaml:"description"`
	RedirectsField   []string        `yaml:"redirects"`
	LanguageField    string          `yaml:"language"`
	KeepLinksField   bool            `yaml:"keeplinks"`
	FaviconPathField string          `yaml:"faviconpath"`
	FrontProxyField  *FrontProxyYaml `yaml:"frontproxy"`
}

// FrontProxyYaml is a structure that will be automatically populated with the value of the configuration
type FrontProxyYaml struct {
	ForceSSL bool   `yaml:"forcessl"`
	IPHeader string `yaml:"ipheader"`
}

// Port return the port to which the daemon will listen to
func (config *ConfigurationYaml) Port() uint16 { return config.PortField }

// Sites returns the list of sites
func (config *ConfigurationYaml) Sites() []Site {
	sites := make([]Site, len(config.SitesField))

	for i := range config.SitesField {
		sites[i] = config.SitesField[i]
	}

	return sites
}

// Index returns true if an index should be displayed when neither a host or a redirect was found
func (config *ConfigurationYaml) Index() bool { return config.IndexField }

// Ref returns the Google Sites reference of the websites
func (site *SiteYaml) Ref() string { return site.RefField }

// Host returns the host that is supposed to handle respond for this site
func (site *SiteYaml) Host() string { return site.HostField }

// Description returns a string describing the content of the site
func (site *SiteYaml) Description() string { return site.DescriptionField }

// Redirects returns a list of hostname that should redirect their traffic to the host
func (site *SiteYaml) Redirects() []string { return site.RedirectsField }

// Language returns the HTTP header used to query Google servers
func (site *SiteYaml) Language() string { return site.LanguageField }

// KeepLinks returns true if links should be kept
func (site *SiteYaml) KeepLinks() bool { return site.KeepLinksField }

// FaviconPath returns the path to the favicon file. No file is assumed if the string is empty
func (site *SiteYaml) FaviconPath() string { return site.FaviconPathField }

// GRef returns a consistent reference to the Google Sites instance
func (site *SiteYaml) GRef() string {
	if !strings.Contains(site.Ref(), "/") {
		return "view/" + site.Ref()
	}
	return site.Ref()
}

// IPHeader returns the HTTP header containing the ip of the requester, if empty this information will
// be taken from the connection directly
func (site *SiteYaml) IPHeader() string {
	if site.FrontProxyField != nil {
		return site.FrontProxyField.IPHeader
	}
	return ""
}

// LoadConfig returns a loader for a given YAML configuration file
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
