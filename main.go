package main

import (
	"SiteStatusChecker/util"
	"fmt"
)

var Site string

func main() {
	fmt.Println("Input the site url for status checking:")
	// Scan the user input to the Global variable's address &Site
	fmt.Scanln(&Site)
	fmt.Printf("The current host your asking for is < %s >\n", util.GetDomainName(Site))
	fmt.Printf("The requested site status is < %s >", util.GetSiteStatus(Site))
}
