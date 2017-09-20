// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package main

import (
	"os"
	log "github.com/sirupsen/logrus"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Cmd: " + os.Args[0] + " config-file")
	}
	confFile := os.Args[1]

	startDaemonFromFile(confFile)
}
