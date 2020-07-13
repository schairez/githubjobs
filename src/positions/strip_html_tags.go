package positions

import (
	"github.com/microcosm-cc/bluemonday"
)

func stripHTML(html string) string {
	p := bluemonday.StripTagsPolicy()
	return p.Sanitize(html)

}
