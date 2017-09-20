// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/fever-ch/go-google-sites-proxy/proxy"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"os/signal"
	"syscall"
	"fmt"
	"github.com/fever-ch/go-google-sites-proxy/common"
)

// Load configuration stored in filename (yaml format)
func loadConfig(filename string) (common.Configuration, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return common.Configuration{}, err
	}

	c := common.Configuration{}
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return common.Configuration{}, err
	}

	return c, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Cmd: " + os.Args[0] + " config-file")
	}

	confFile := os.Args[1]
	os.Chdir(filepath.Dir(confFile))

	if cfg, err := loadConfig(confFile); err != nil {
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
			if cfg, err := loadConfig(confFile); err != nil {
				log.WithError(err).Warn("Unable to parse config")
			} else if cfg.Port() != proxy.Port() {
				log.Warning(fmt.Sprintf("Server currently running on port %d but config specifies %d. This change will "+
					"not be taken in account. Please restart daemon.", proxy.Port(), cfg.Port))
			} else {
				proxy.SetConfiguration(cfg)
				log.Info("Configuration reloaded")
			}
		}
	}
}
