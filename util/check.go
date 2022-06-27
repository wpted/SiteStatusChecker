package util

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

// GetDomainName returns the Host name of the domain
func GetDomainName(siteUrl string) string {
	// resp is a type URL struct which represents the parsed URL
	resp, err := url.Parse(siteUrl)

	// panic if input isn't correct or something else fatal happened
	if err != nil {
		log.Panic(err)
	}

	// resp.Host is a string formatted something like ww w.example.com
	// split the string by '.' and get the domain name by the removing 'www.'
	segments := strings.Split(resp.Host, ".")
	domainName := segments[len(segments)-2] + "." + segments[len(segments)-1]
	return domainName
}

// GetSiteStatus returns the current status description of the site we pass in
func GetSiteStatus(siteUrl string) string {
	// resp is a type Response struct which represents the response from an HTTP request
	resp, err := http.Get(siteUrl)

	// panic if input isn't correct or something else fatal happened
	if err != nil {
		log.Panic(err)
	}

	return resp.Status
}
