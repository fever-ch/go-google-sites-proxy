// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package proxy

import "go-google-sites-proxy/blob"

type Page struct {
	Code             int
	Headers          map[string](string)
	Blob             blob.HybridRawGzipBlob
	OriginallyGziped bool
}
