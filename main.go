// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"log"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"go-google-sites-proxy/proxy"
	"go-google-sites-proxy/config"
)

// Load configuration stored in filename (yaml format)
func loadConfig(filename string) (config.Configuration, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return config.Configuration{}, err
	}

	c := config.Configuration{}
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return config.Configuration{}, err
	}

	return c, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Cmd: " + os.Args[0] + " config-file")
	}

	cfg, err := loadConfig(os.Args[1])
	proxy := proxy.NewCheapProxy(cfg)

	if err != nil {
		log.Fatal(err)
	}


	err = proxy.Start()
	if err != nil {
		log.Fatal(err)
	}
}
