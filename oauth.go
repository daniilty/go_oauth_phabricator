package phabricator

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

type oauthResponse struct {
	Result    User   `json:"result"`
	ErrorCode string `json:"error_code"`
	ErrorInfo string `json:"error_info"`
}

func (c *Config) token(code string) (*oauth2.Token, error) {
	token, err := c.oauth.Exchange(context.Background(), code)
	if err != nil {
		return token, fmt.Errorf("c.oauth.Exchange: %w", err)
	}

	if !token.Valid() {
		return token, fmt.Errorf("token is invalid: %w", err)
	}

	return token, nil
}

func (c *Config) body(token *oauth2.Token) ([]byte, error) {
	authClient := c.oauth.Client(context.Background(), token)

	getClientInfoURL := c.phabricatorURL + "/api/user.whoami?access_token=" + token.AccessToken

	authResponse, err := authClient.Get(getClientInfoURL)
	if err != nil {
		return []byte{}, fmt.Errorf("authClient.Get: %w", err)
	}
	defer authResponse.Body.Close()

	if authResponse.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("statusCode is not ok: %w", err)
	}

	return ioutil.ReadAll(authResponse.Body)
}

func (c *Config) unmarshal(body []byte) (User, error) {
	var resp oauthResponse

	if err := json.Unmarshal(body, &resp); err != nil {
		return resp.Result, fmt.Errorf("json.Unmarshal: %w", err)
	}

	if resp.ErrorCode != "" {
		return resp.Result, fmt.Errorf("can't find user info: %s", resp.ErrorInfo)
	}

	return resp.Result, nil
}
