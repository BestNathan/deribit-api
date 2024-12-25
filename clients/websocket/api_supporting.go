package websocket

import (
	"github.com/BestNathan/deribit-api/pkg/models"
)

func (c *DeribitWSClient) GetTime() (result int64, err error) {
	err = c.Call("public/get_time", nil, &result)
	return
}

func (c *DeribitWSClient) Hello(params *models.HelloParams) (result models.HelloResponse, err error) {
	err = c.Call("public/hello", params, &result)
	return
}

func (c *DeribitWSClient) Test() (result models.TestResponse, err error) {
	err = c.Call("public/test", nil, &result)
	return
}
