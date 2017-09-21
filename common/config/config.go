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
	GRef() string

	IPHeader() string
}

type ConfigLoader func() (Configuration, error)
