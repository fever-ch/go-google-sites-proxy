# Go Google Sites Proxy
[![Go Report Card](https://goreportcard.com/badge/github.com/rbarazzutti/go-google-sites-proxy)](https://goreportcard.com/report/github.com/rbarazzutti/go-google-sites-proxy)


Project homepage: [https://ggsp.fever.ch](https://ggsp.fever.ch)

## Intro

A simple proxy for the [New Google Sites](https://sites.google.com/new).

Fall 2016, Google introduced New Google Sites, it's a very nice product, it allows the creation of nice and responsive websites that suits well either for mobile and desktop use.

Sadly, Google [at least at the time this text is being written, [August 24th, 2017] doesn't support the usage of a custom domain name yet [[link]](https://productforums.google.com/forum/#!topic/sites/44_WTQ44MJk).

## Features
- Works with any domain!
- Support for custom favicon

## Try it!

This software is already used *in production*. The [website of this project](https://ggsp.fever.ch/) also using it!

## Config file format

```yaml
port: 9080
sites:
  - ref: go-gsites-proxy                # the name of the website on Google Sites
    host: ggsp.fever.ch                 # the vhost linked with the current site
    language: en-US, en;q=0.9           # a language hint, to get a properly localized page
    description: Go Google Sites Proxy  # a short description of this site
``` 

## To-Do

- add cache for the retrieved content


## How to install 

### Build from sources

Fetch sources and build it:

    go get -v -u github.com/fever-ch/go-google-sites-proxy


Once the `get` completes, you should find your new `go-google-sites-proxy` (or `go-google-sites-proxy.exe`) executable sitting inside `$GOPATH/bin/`.


### Execute it

Create your own configuration file (formatted in yaml). This project contains an example named `config-example.yaml`([link](https://github.com/fever-ch/go-google-sites-proxy/blob/master/config-example.yaml)). 

    nohup $GOPATH/bin/go-google-sites-proxy my-custom-config.yaml &



## Use it with Docker :

Let's assume that the configuration file (```config.yaml```) is in ```my-config-folder```

    docker pull feverch/go-google-sites-proxy
    docker run -d --name containerName -v /full-path/to/my-config-folder:/etc/ggsp/ -p 80:9080 feverch/go-google-sites-proxy



#### Build your own image with Rocker
    rocker build .

This project migrated to [Rocker](https://github.com/grammarly/rocker) for builds. It allows the creation of much slimmer images (no layered images with all compilation/build steps). Thanks to Rocker the image is now smaller than 10mb!

This change has no impact for users of the image (no need to install Rocker). 

## HTTPS, HTTP/2.0 support

`Go Google Sites Proxy` doesn't support yet, and might never, HTTPS and HTTP/2.0 protocols. Since these protocols are nowadays *a must*, but on the other hand, they require fine tuning, a clever approach would be to use [NGINX](https://www.nginx.org) [(docker image)](https://hub.docker.com/_/nginx/) or to use a well-known service such that [CloudFlare](https://www.cloudflare.com). 


2017 fever.ch - Raphaël P. Barazzutti 
