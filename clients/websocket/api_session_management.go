package websocket

import (
	"github.com/BestNathan/deribit-api/pkg/models"
)

func (c *DeribitWSClient) SetHeartbeat(params *models.SetHeartbeatParams) (result string, err error) {
	err = c.Call("public/set_heartbeat", params, &result)
	return
}

func (c *DeribitWSClient) DisableHeartbeat() (result string, err error) {
	err = c.Call("public/disable_heartbeat", nil, &result)
	return
}

func (c *DeribitWSClient) EnableCancelOnDisconnect() (result string, err error) {
	err = c.Call("private/enable_cancel_on_disconnect", nil, &result)
	return
}

func (c *DeribitWSClient) DisableCancelOnDisconnect() (result string, err error) {
	err = c.Call("private/disable_cancel_on_disconnect", nil, &result)
	return
}
