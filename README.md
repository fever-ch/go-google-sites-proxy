# Go Google Sites Proxy
Project homepage: [https://ggsp.fever.ch](https://ggsp.fever.ch)

## Intro

A simple proxy for the [New Google Sites](https://sites.google.com/new).

Fall 2016, Google introduced New Google Sites, it's a very nice product, it allows the creation of nice and responsive websites that suits well either for mobile and desktop use.

Sadly, Google [at least at the time this text is being written, [August 24th, 2017] doesn't support the usage of a custom domain name yet [[link]](https://productforums.google.com/forum/#!topic/sites/44_WTQ44MJk).

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

### Run it on any architecture supported by Golang

Build a Linux static binary from Linux, OSX, Microsoft Windows, ... (cross-compilation)

``` bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
```
Then run it:
``` bash
nohup go-google-sites-proxy my-custom-config.yaml &
``` 

### Use it with Docker :

Let's assume that the configuration file (```config.yaml```) is in ```my-config-folder```

```bash
docker pull feverch/go-google-sites-proxy
docker run -d --name containerName -v /full-path/to/my-config-folder:/etc/ggsp/ -p 80:9080 feverch/go-google-sites-proxy
```


#### Build your own image with Rocker
```bash
rocker build .
```
This project migrated to [Rocker](https://github.com/grammarly/rocker) for builds. It allows the creation of much slimmer images (no layered images with all compilation/build steps).

This change has no impact for users of the image (no need to install Rocker). 

### HTTPS, HTTP/2.0 support

`Go Google Sites Proxy` doesn't support yet, and might never, HTTPS and HTTP/2.0 protocols. Since these protocols are nowadays *a must*, but on the other hand, they require fine tuning, a clever approach would be to use [NGINX](https://www.nginx.org) [(docker image)](https://hub.docker.com/_/nginx/) or to use a well-known service such that [CloudFlare](https://www.cloudflare.com). 


2017 fever.ch - Raphaël P. Barazzutti 
