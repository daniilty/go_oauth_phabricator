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
	phabricator "github.com/Megaputer/go_oauth_phabricator"

	"github.com/astaxie/beego"
)

var client *phabricator.Config

func init() {
	phabricatorURL := beego.AppConfig.DefaultString("PhabricatorURL", "https://phabricator.megaputer.ru")
	redirectURL := beego.AppConfig.DefaultString("RedirectURL", "http://metrics.megaputer.ru/auth")
	oauthPHID := beego.AppConfig.String("OAuthPHID")
	oauthSecret := beego.AppConfig.String("OAuthSecret")

	client = phabricator.ClientConfig(oauthPHID, oauthSecret, redirectURL, phabricatorURL)
}

//Auth OAuth Phabricator
func Auth(code string) (User, error) {
	user, err := client.Authenticate(code)
	return User(user), err
}

// URL return url from OAuth
func URL() string {
	return client.AuthCodeURL("")
}

```