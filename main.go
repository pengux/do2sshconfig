package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func main() {
	ipv6 := flag.Bool("ipv6", false, "Use IPv6 instead of IPv4 as hostnames where possible")
	user := flag.String("user", "root", "Use a different user instead of 'root'")
	idFile := flag.String("idFile", "", "Path of private SSH key to use")
	flag.Parse()

	if len(flag.Args()) < 1 || flag.Args()[0] == "" {
		log.Fatal("no Personal Access Token provided, please pass it in as first argument")
	}

	// Authentication with oAuth2
	tokenSource := &TokenSource{
		AccessToken: flag.Args()[0],
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := godo.NewClient(oauthClient)

	fmt.Println("# --- DigitalOcean hosts - Start ---")

	// Iterate over droplets
	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := client.Droplets.List(opt)
		if err != nil {
			log.Fatal(err)
		}

		// append the current page's droplets to our list
		for _, d := range droplets {
			var ip string
			// Find public IP
			if *ipv6 {
				for _, v6 := range d.Networks.V6 {
					if v6.Type == "public" {
						ip = v6.IPAddress
						break
					}
				}
			}

			if ip == "" {
				for _, v4 := range d.Networks.V4 {
					if v4.Type == "public" {
						ip = v4.IPAddress
						break
					}
				}
			}

			fmt.Println(`Host`, d.Name)
			fmt.Println(`	HostName`, ip)
			fmt.Println(`	User`, *user)
			if *idFile != "" {
				fmt.Println(`	IdentityFile`, *idFile)
			}
			fmt.Println(``)
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			log.Fatal(err)
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}

	fmt.Println("# --- DigitalOcean hosts - End ---")
}
