package client

import (
	"encoding/json"
	"github.com/magicshui/qingcloud-go/model"
)

func (c *Client) GetMonitor(params Params) *model.ModelMonitorI {
	result := c.Get("GetMonitor", params)

	var res model.ModelMonitorI
	json.Unmarshal([]byte(result), &res)
	return &res
}
