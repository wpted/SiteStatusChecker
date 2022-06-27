package util

import (
	"testing"
)

func TestGetDomainName(t *testing.T) {
	mockSite := "https://google.com"
	mockDomainName := GetDomainName(mockSite)
	if mockDomainName != "google.com" {
		t.Errorf("wrong domain name, expecting 'google.com', got %s", mockDomainName)
	}
}

func TestGetSiteStatus(t *testing.T) {
	mockSite := "https://google.com"
	mockStatus := GetSiteStatus(mockSite)
	if mockStatus != "200 OK" {
		t.Errorf("wrong status description, expecting '200 OK', got %s", mockStatus)
	}
}
