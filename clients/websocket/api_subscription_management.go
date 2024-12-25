package websocket

import (
	"github.com/BestNathan/deribit-api/pkg/models"
)

func (c *DeribitWSClient) PublicSubscribe(params *models.SubscribeParams) (result models.SubscribeResponse, err error) {
	err = c.Call("public/subscribe", params, &result)
	return
}

func (c *DeribitWSClient) PublicUnsubscribe(params *models.UnsubscribeParams) (result models.UnsubscribeResponse, err error) {
	err = c.Call("public/unsubscribe", params, &result)
	return
}

func (c *DeribitWSClient) PrivateSubscribe(params *models.SubscribeParams) (result models.SubscribeResponse, err error) {
	err = c.Call("private/subscribe", params, &result)
	return
}

func (c *DeribitWSClient) PrivateUnsubscribe(params *models.UnsubscribeParams) (result models.UnsubscribeResponse, err error) {
	err = c.Call("private/unsubscribe", params, &result)
	return
}
