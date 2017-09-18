// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package proxy

import (
	"bytes"
	"regexp"
	"github.com/fever-ch/go-google-sites-proxy/blob"
	"strconv"
	"github.com/fever-ch/go-google-sites-proxy/utils"
	"github.com/fever-ch/go-google-sites-proxy/common"
)

func patchLinks(input [] byte, site *common.Site) []byte {
	return bytes.Replace(input, []byte( "\"/view/"+site.Ref), []byte( "\""), -1)
}

func newPatcher(site *common.Site, context *SiteContext) func(*Page) *Page {
	var htmlRx, _ = regexp.Compile("text/html($|;.*)")
	patchLinks := func(input *Page) *Page {
		if !site.KeepLinks && htmlRx.MatchString(input.Headers["Content-Type"]) {
			return &Page{input.Code,
				input.Headers,
				blob.NewRawBlob(bytes.Replace(input.Blob.Raw(), []byte( "\"/view/"+site.Ref), []byte( "\""), -1)),
				input.OriginallyGziped}
		} else {
			return input
		}
	}

	ep := int(utils.Epoch())

	patchFavicon := func(input *Page) *Page {
		if context.Favicon != nil && htmlRx.MatchString(input.Headers["Content-Type"]) {
			return &Page{input.Code,
				input.Headers,
				blob.NewRawBlob(bytes.Replace(input.Blob.Raw(),
					[]byte("<link rel=\"icon\" href=\"//ssl.gstatic.com/atari/images/favicon_2.ico\"/>"),
					[]byte("<link rel=\"icon\" href=\"/favicon.ico?v="+strconv.Itoa(ep)+"\"/>"),
					-1)),

				input.OriginallyGziped}
		} else {
			return input
		}
	}

	return func(input *Page) *Page {
		return patchFavicon(patchLinks(input))
	}
}
