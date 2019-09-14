package unwarc

import (
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

// takes a URL path and returns directory, filename
func Sanitise(str string) (string, string) {
	u, err := url.Parse(str)
	if err == nil && u.Opaque == "" {
		str = path.Join(u.Host, u.Path)
		if u.RawQuery != "" {
			str = str + "_" + u.RawQuery
		}
		if u.Fragment != "" {
			str = str + "_" + u.Fragment
		}
	}
	str = filepath.FromSlash(str)
	return filepath.Split(str)
}

// takes a path like blackbooks.warc.gz and gives blackbooks
func Base(str string) string {
	str = filepath.Base(str)
	str = strings.TrimSuffix(str, filepath.Ext(str))
	return strings.TrimSuffix(str, filepath.Ext(str))
}
