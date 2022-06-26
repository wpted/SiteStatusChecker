package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var Site string

// getDomainName returns the Host name of the domain
func getDomainName(siteUrl string) string {
	resp, err := url.Parse(siteUrl)

	if err != nil {
		log.Panic(err)
	}

	segments := strings.Split(resp.Host, ".")
	domainName := segments[1] + "." + segments[2]
	return domainName
}

// getSiteStatus returns the current status description of the site we pass in
func getSiteStatus(siteUrl string) string {
	// resp is a type Response struct which represents the response from an HTTP request
	resp, err := http.Get(siteUrl)

	if err != nil {
		log.Panic(err)
	}
	return resp.Status
}

func main() {
	fmt.Println("Input the site url for status checking:")
	fmt.Scanln(&Site)
	fmt.Printf("The current host your asking for is < %s >\n", getDomainName(Site))

	fmt.Printf("The requested site status is < %s >", getSiteStatus(Site))
}
