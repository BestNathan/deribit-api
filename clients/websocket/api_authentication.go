package websocket

import (
	websocketmodels "github.com/BestNathan/deribit-api/clients/websocket/models"
	"github.com/BestNathan/deribit-api/pkg/models"
)

func (c *DeribitWSClient) Auth(apiKey string, secretKey string) (err error) {
	params := models.ClientCredentialsParams{
		GrantType:    "client_credentials",
		ClientID:     apiKey,
		ClientSecret: secretKey,
	}
	var result models.AuthResponse
	err = c.Call("public/auth", params, &result)
	if err != nil {
		return
	}

	c.authentication = &websocketmodels.Authentication{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}

	return
}

func (c *DeribitWSClient) Logout() (err error) {
	var result = struct {
	}{}
	err = c.Call("public/auth", nil, &result)
	return
}
