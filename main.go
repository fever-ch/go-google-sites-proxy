package main

import (
	"fmt"
	"github.com/fever-ch/go-google-sites-proxy/common"
	log "github.com/sirupsen/logrus"
	"os"
)

// GitVersion is a variable representing the version of the repo at build time (set with -X)
var GitVersion = ""

// BuildDate is the date at the time the software was compiled (set with -X)
var BuildDate = ""

func main() {
	if GitVersion != "" {
		common.ProgramInfo.Git = GitVersion
	}
	if BuildDate != "" {
		common.ProgramInfo.BuildDate = BuildDate
	}
	log.Info(fmt.Sprintf("GGSP %s %s", common.ProgramInfo.Git, common.ProgramInfo.BuildDate))

	if len(os.Args) != 2 {
		log.Fatal("Cmd: " + os.Args[0] + " config-file")
	}
	confFile := os.Args[1]

	startDaemonFromFile(confFile)
}
