package proxy

import (
	"bytes"
	"github.com/fever-ch/go-google-sites-proxy/blob"
	"github.com/fever-ch/go-google-sites-proxy/common/config"
	"github.com/fever-ch/go-google-sites-proxy/utils"
	"regexp"
	"strconv"
)

func patchLinks(input []byte, site config.Site) []byte {
	return bytes.Replace(input, []byte("\"/"+site.GRef()), []byte("\""), -1)
}

func newPatcher(site config.Site, context *siteContext) func(*Page) *Page {
	var htmlRx, _ = regexp.Compile("text/html($|;.*)")
	patchLinks := func(input *Page) *Page {
		if !site.KeepLinks() && htmlRx.MatchString(input.Headers["Content-Type"]) {
			return &Page{input.Code,
				input.Headers,
				blob.NewRawBlob(bytes.Replace(input.Blob.Raw(), []byte("\"/"+site.GRef()), []byte("\""), -1)),
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
