package githubjobs

import (
	"github.com/microcosm-cc/bluemonday"
)

//stripHTML removes html tags from string
func stripHTML(html string) string {
	p := bluemonday.StripTagsPolicy()
	return p.Sanitize(html)

}
