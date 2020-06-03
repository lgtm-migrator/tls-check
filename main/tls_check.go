package main

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
)

// CheckResult for tls
type CheckResult struct {
	TLS10support bool
	TLS11support bool
	TLS12support bool
	TLS13support bool
	TLS30support bool
}

func createHTTPClient(tlsVersion uint16) *http.Client {
	return &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{
			MinVersion: tlsVersion,
			MaxVersion: tlsVersion,
		}},
	}
}

func checkVersionSupport(tlsVersion uint16, url string) bool {
	if _, err := createHTTPClient(tls.VersionTLS11).Get(url); err != nil {
		return false
	}
	return true
}

// CheckTLS for URL
func CheckTLS(u *url.URL) (r *CheckResult, e error) {
	if u == nil {
		return nil, errors.New("url must provided")
	}
	if u.Scheme != "https" {
		return nil, errors.New("only support https url")
	}
	r = &CheckResult{}
	r.TLS10support = checkVersionSupport(tls.VersionTLS10, u.String())
	r.TLS11support = checkVersionSupport(tls.VersionTLS11, u.String())
	r.TLS12support = checkVersionSupport(tls.VersionTLS12, u.String())
	r.TLS13support = checkVersionSupport(tls.VersionTLS13, u.String())
	return
}
