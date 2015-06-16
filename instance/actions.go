package instance

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeInstances(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeInstances", params)
}

func (c *Client) RunInstances(params qingcloud.Params) ([]byte, error) {
	return c.Get("RunInstances", params)
}

func (c *Client) TerminateInstances(params qingcloud.Params) ([]byte, error) {
	return c.Get("TerminateInstances", params)
}

func (c *Client) StartInstances(params qingcloud.Params) ([]byte, error) {
	return c.Get("StartInstances", params)
}

func (c *Client) StopInstances(params qingcloud.Params) ([]byte, error) {
	return c.Get("StopInstances", params)
}

func (c *Client) RestartInstances(params qingcloud.Params) ([]byte, error) {
	return c.Get("RestartInstances", params)
}

func (c *Client) ResetInstances(params qingcloud.Params) ([]byte, error) {
	return c.Get("ResetInstances", params)
}

func (c *Client) ResizeInstances(params qingcloud.Params) ([]byte, error) {
	return c.Get("ResizeInstances", params)
}

func (c *Client) ModifyInstanceAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyInstanceAttributes", params)
}
