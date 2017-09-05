FROM golang

ADD . /go/src/go-google-sites-proxy

RUN cd /go/src/go-google-sites-proxy; go get; go install go-google-sites-proxy

ENTRYPOINT /go/bin/go-google-sites-proxy /etc/ggsp/config.yaml

EXPOSE 9080
