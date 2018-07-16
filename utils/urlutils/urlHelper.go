package urlutils

import (
	"net/url"
	"path"
)

// Join joins a domain with following paths
func Join(domain string, paths ...string) string {
	u, _ := url.Parse(domain)
	for _, pathItem := range paths {
		u.Path = path.Join(u.Path, pathItem)
	}
	return u.String()
}
