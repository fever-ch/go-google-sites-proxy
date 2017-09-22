package main

import (
	"fmt"
	"github.com/fever-ch/go-google-sites-proxy/common"
	log "github.com/sirupsen/logrus"
	"os"
)

var gitVersion = ""

var buildDate = ""

func main() {
	if gitVersion != "" {
		common.ProgramInfo.Git = gitVersion
	}
	if buildDate != "" {
		common.ProgramInfo.BuildDate = buildDate
	}
	log.Info(fmt.Sprintf("GGSP %s %s", common.ProgramInfo.Git, common.ProgramInfo.BuildDate))

	if len(os.Args) != 2 {
		log.Fatal("Cmd: " + os.Args[0] + " config-file")
	}
	confFile := os.Args[1]

	startDaemonFromFile(confFile)
}
