package config

// Configuration represents the setup of the whole GGSP
type Configuration interface {
	Sites() []Site
	Index() bool
	Port() uint16
}

// Site represents the settings of a specific site
type Site interface {
	Ref() string
	Host() string
	Description() string
	Redirects() []string
	Language() string
	KeepLinks() bool
	FaviconPath() string
	GRef() string

	IPHeader() string
}

// ConfigLoader is a function that can load a configuration from
// specific manner, i.e. NewYamlConfigLoader
type ConfigLoader func() (Configuration, error)
