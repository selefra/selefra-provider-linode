package linode_client

import (
	"context"
	"github.com/linode/linodego"
	"golang.org/x/oauth2"
	"net/http"
)

func Connect(_ context.Context, config *Config) (linodego.Client, error) {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.Token})

	oauth2Client := &http.Client{
		Transport: &oauth2.Transport{
			Source: tokenSource,
		},
	}

	conn := linodego.NewClient(oauth2Client)

	return conn, nil
}
