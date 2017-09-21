// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/fever-ch/go-google-sites-proxy/common/config"
	"github.com/fever-ch/go-google-sites-proxy/proxy"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func startDaemonFromFile(confFile string) {
	os.Chdir(filepath.Dir(confFile))
	startDaemon(config.LoadConfig(confFile))
}

func startDaemon(cl config.ConfigLoader) {

	if cfg, err := cl(); err != nil {
		log.WithError(err).Fatal("Unable to load configuration")
	} else {
		proxy := proxy.NewCheapProxy(cfg.Port())
		proxy.SetConfiguration(cfg)

		startUp := func() {
			if err = proxy.Start(); err != nil {
				log.WithError(err).Fatal("Unable to start proxy")
			}
		}

		go startUp()

		for true {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGUSR1)
			<-c
			if cfg, err := cl(); err != nil {
				log.WithError(err).Warn("Unable to load config")
			} else if cfg.Port() != proxy.Port() {
				log.Warning(fmt.Sprintf("Server currently running on port %d but config specifies %d. This change will "+
					"not be taken in account. Please restart daemon.", proxy.Port(), cfg.Port()))
			} else {
				proxy.SetConfiguration(cfg)
				log.Info("Configuration reloaded")
			}
		}
	}
}
