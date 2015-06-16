package monitor

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) GetMonitor(params qingcloud.Params) ([]byte, error) {
	return c.Get("GetMonitor", params)
}

func (c *Client) GetLoadBalancerMonitor(params qingcloud.Params) ([]byte, error) {
	return c.Get("GetLoadBalancerMonitor", params)
}
func (c *Client) GetRDBMonitor(params qingcloud.Params) ([]byte, error) {
	return c.Get("GetRDBMonitor", params)
}
func (c *Client) GetCacheMonitor(params qingcloud.Params) ([]byte, error) {
	return c.Get("GetCacheMonitor", params)
}
