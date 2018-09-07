# go_oauth_phabricator
Phabricator API with OAuth2 in Golang

Installation and Usage
=============


Install
---------------
    go get -v github.com/Megaputer/go_oauth_phabricator

Usage
---------------

    

```go
package main

import (
	"fmt"
	"log"

	phabricator "github.com/Megaputer/go_oauth_phabricator"
)

var client *phabricator.Config

// initialize the client in the init () function
func init() {
	// Get oauthPHID and oauthSecret from
	// https://example.phabricator.com/oauthserver/query/all/
	oauthPHID := "OAuthPHID"
	oauthSecret := "OAuthSecret"

	// redirectURL is the URL to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	redirectURL := "https://my.com/auth"

	//phabricatorURL the url of the phabricator server
	// that is the source of OAuth
	phabricatorURL := "https://phabricator.exapmle.ru"

	client = phabricator.ClientConfig(oauthPHID, oauthSecret, redirectURL, phabricatorURL)
}

func main() {
	// AuthCodeURL return url from OAuth with CSRF token
	url := client.AuthCodeURL("CSRF token")
	fmt.Println(url)

	// code will be in the *http.Request.FormValue("code")
	// https://secure.phabricator.com/book/phabcontrib/article/using_oauthserver/
	code := ""

	user, err := client.Authenticate(code)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(user.UserName, user.RealName)
}

```