package core

import (
	"strconv"

	"github.com/mileusna/useragent"
)

// browser-version
var supportedVersionByBrowser = map[string]int{
	useragent.Chrome: 103,
	useragent.Edge:   103,
}

// IsSupported indicates whether the browser supports EarlyHints.
func IsSupported(userAgent string) bool {
	p := useragent.Parse(userAgent)

	if p.Bot {
		return false
	}

	fromVersion, found := supportedVersionByBrowser[p.Name]
	if !found {
		return false
	}

	versionInt, err := strconv.Atoi(p.Version)
	if err != nil {
		return false
	}

	return versionInt > fromVersion
}
