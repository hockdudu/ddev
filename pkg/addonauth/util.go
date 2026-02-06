package addonauth

import (
	"errors"
	"net/url"
	"strings"
)

func normalizeRemote(remote string) (error, string) {
	remote = strings.ToLower(remote)

	if strings.HasPrefix(remote, "http://") || strings.HasPrefix(remote, "https://") {
		// TODO: What does Composer do? Can we safely validate?
		return errors.New("remote must not have a protocol prefix (e.g. http:// or https://)"), ""
	}

	// TODO: Can we safely trim instead?
	if strings.HasSuffix(remote, "/") {
		return errors.New("remote must not have a trailing slash"), ""
	}

	return nil, remote
}

// TODO: Should we add paths as well?
func doesUrlMatchRemote(requestUrl string, remote string) bool {
	u, err := url.Parse(requestUrl)
	if err != nil {
		return false
	}

	baseUrl := strings.ToLower(u.Host)

	// e.g. allows either `baseUrl="github.com"` or `baseUrl="www.github.com"` to match `remote="github.com"`
	return baseUrl == remote || strings.HasSuffix(baseUrl, "."+remote)
}
