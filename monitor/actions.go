package monitor

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) GetMonitor(params qingcloud.Params) ([]byte, error) {
	return c.Get("GetMonitor", params)
}
