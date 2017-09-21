package blob

import (
	"bytes"
	"compress/gzip"
)

// Hybrid blob that can gives its content in raw or gzipped form
type HybridRawGzipBlob struct {
	Raw     func() []byte
	Gzipped func() []byte
}

// Build a new hybrid blob from a raw form
func NewRawBlob(raw []byte) HybridRawGzipBlob {
	var gzipped []byte
	return HybridRawGzipBlob{

		Raw: func() []byte { return raw },
		Gzipped: func() []byte {
			if gzipped == nil {
				var bbb bytes.Buffer
				gz := gzip.NewWriter(&bbb)
				gz.Write(raw)
				gz.Flush()
				gz.Close()

				gzipped = bbb.Bytes()
			}
			return gzipped
		}}
}

// Build a new hybrid blob from a gzipped form
func NewGzippedBlob(gzipped []byte) HybridRawGzipBlob {
	var raw []byte
	return HybridRawGzipBlob{

		Raw: func() []byte {
			if raw == nil {
				r, _ := gzip.NewReader(bytes.NewBuffer(gzipped))
				var resB bytes.Buffer
				resB.ReadFrom(r)
				raw = resB.Bytes()
			}
			return raw
		},
		Gzipped: func() []byte { return gzipped }}
}
