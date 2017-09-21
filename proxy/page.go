package proxy

import "github.com/fever-ch/go-google-sites-proxy/blob"

// Page represents a retrieved page
type Page struct {
	Code             int
	Headers          map[string](string)
	Blob             blob.HybridRawGzipBlob
	OriginallyGziped bool
}
