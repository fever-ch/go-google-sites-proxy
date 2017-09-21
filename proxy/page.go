package proxy

import "github.com/fever-ch/go-google-sites-proxy/blob"

type Page struct {
	Code             int
	Headers          map[string](string)
	Blob             blob.HybridRawGzipBlob
	OriginallyGziped bool
}
