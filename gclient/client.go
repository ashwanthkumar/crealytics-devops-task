package gclient

import (
	"net/http"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// cache the client instance for the entire lifetime of the application to speed things up
var client *http.Client

// CreateOrGetClient returns an instance of *http.Client that is wrapped around Google's OAuth2.
func CreateOrGetClient(scopes []string) (*http.Client, error) {
	if nil == client {
		ctx := context.Background()
		token, err := google.DefaultTokenSource(ctx, strings.Join(scopes, " "))
		if err != nil {
			return nil, err
		}
		client = oauth2.NewClient(ctx, token)
	}
	return client, nil
}
