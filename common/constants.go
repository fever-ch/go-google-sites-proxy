package common

// ProgramInfoStruct is a structure that define the program (incl. build informations)
type ProgramInfoStruct struct {
	Name       string
	Fullname   string
	Version    string
	Git        string
	BuildDate  string
	ProjectURL string
}

// ProgramInfo is a variable containing the informations about the current build
var ProgramInfo = ProgramInfoStruct{"GSSP",
	"Go Google Site Proxy",
	"0.0.1-SNAPSHOT",
	"undefined",
	"undefined",
	"https://ggsp.fever.ch"}
