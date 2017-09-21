package blob

import (
	"bytes"
	"compress/gzip"
)

// HybridRawGzipBlob is an hybrid blob that can gives its content in raw or gzipped form
type HybridRawGzipBlob struct {
	Raw     func() []byte
	Gzipped func() []byte
}

// NewRawBlob returns a new hybrid blob from a raw form
func NewRawBlob(raw []byte) HybridRawGzipBlob {
	var gzipped []byte
	return HybridRawGzipBlob{

		Raw: func() []byte { return raw },
		Gzipped: func() []byte {
			if gzipped == nil {
				var bbb bytes.Buffer
				gz := gzip.NewWriter(&bbb)
				_, e := gz.Write(raw)
				if e != nil || gz.Flush() != nil || gz.Close() != nil {
					return make([]byte, 0)
				}

				gzipped = bbb.Bytes()
			}
			return gzipped
		}}
}

// NewGzippedBlob returns a new hybrid blob from a gzipped form
func NewGzippedBlob(gzipped []byte) HybridRawGzipBlob {
	var raw []byte
	return HybridRawGzipBlob{

		Raw: func() []byte {
			if raw == nil {
				r, e := gzip.NewReader(bytes.NewBuffer(gzipped))
				if e != nil {
					return make([]byte, 0)
				}
				var resB bytes.Buffer
				_, e = resB.ReadFrom(r)
				if e != nil {
					return make([]byte, 0)
				}
				raw = resB.Bytes()
			}
			return raw
		},
		Gzipped: func() []byte { return gzipped }}
}
