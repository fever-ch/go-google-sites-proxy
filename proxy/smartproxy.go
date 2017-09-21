package proxy

import (
	"github.com/fever-ch/go-google-sites-proxy/common/config"
	"net/http"
	"strconv"
	"sync/atomic"
	"unsafe"
)

type context struct {
	configuration config.Configuration
	sites         map[string]*func(responseWriter http.ResponseWriter, request *http.Request)
}

// NewSmartProxy returns a smartProxy listening to a given TCP port
func NewSmartProxy(port uint16) *smartProxy {

	buildContext := func(configuration config.Configuration) *context {

		ctx := context{
			configuration,
			make(map[string]*func(responseWriter http.ResponseWriter, request *http.Request))}

		for _, e := range configuration.Sites() {

			addRedirect := func(redirectedHost string, destHost string) *func(responseWriter http.ResponseWriter, request *http.Request) {
				redirectHandler := func(response http.ResponseWriter, req *http.Request) {

					response.WriteHeader(http.StatusMovedPermanently)
					response.Header().Add("Location", req.Proto+"://"+destHost+"/"+req.URL.Path)

					response.Write([]byte(strconv.Itoa(http.StatusMovedPermanently) + " Moved permanently"))
				}
				return &redirectHandler
			}

			ctx.sites[e.Host()] = GetSiteHandler(e)
			for _, f := range e.Redirects() {
				ctx.sites[f] = addRedirect(f, e.Host())
			}
		}
		return &ctx
	}

	var ctx unsafe.Pointer

	handler := func(responseWriter http.ResponseWriter, request *http.Request) {
		ctx := (*context)(atomic.LoadPointer(&ctx))
		siteHandler := ctx.sites[request.Host]
		if siteHandler != nil {
			(*siteHandler)(responseWriter, request)
		} else if ctx.configuration.Index() {
			(*getIndex(ctx.configuration))(responseWriter, request)
		}
	}

	return &smartProxy{
		SetConfiguration: func(configuration config.Configuration) {
			atomic.StorePointer(&ctx, unsafe.Pointer(buildContext(configuration)))
		},
		Port: func() uint16 { return port },
		Start: func() error {
			http.HandleFunc("/", handler)
			return http.ListenAndServe(":"+strconv.Itoa(int(port)), nil)
		},
	}
}

type smartProxy struct {
	Start            func() error
	SetConfiguration func(config.Configuration)
	Port             func() uint16
}

