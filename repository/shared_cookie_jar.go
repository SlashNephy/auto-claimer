package repository

import (
	"net/http"
	"net/url"

	"github.com/samber/lo"
)

type SharedCookieJar struct {
	cookieByName map[string]*http.Cookie
}

func NewSharedCookieJar(initial []*http.Cookie) http.CookieJar {
	return &SharedCookieJar{
		cookieByName: lo.KeyBy(initial, func(c *http.Cookie) string {
			return c.Name
		}),
	}
}

func (j *SharedCookieJar) Cookies(*url.URL) []*http.Cookie {
	return lo.Values(j.cookieByName)
}

func (j *SharedCookieJar) SetCookies(_ *url.URL, cookies []*http.Cookie) {
	for _, c := range cookies {
		j.cookieByName[c.Name] = c
	}
}
