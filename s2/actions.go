package s2

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) CreateS2Server(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateS2Server", params)
}
func (c *Client) DescribeS2Servers(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeS2Servers", params)
}
func (c *Client) ModifyS2Server(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyS2Server", params)
}

func (c *Client) ResizeS2Servers(params qingcloud.Params) ([]byte, error) {
	return c.Get("ResizeS2Servers", params)
}
func (c *Client) DeleteS2Servers(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteS2Servers", params)
}
func (c *Client) PowerOnS2Servers(params qingcloud.Params) ([]byte, error) {
	return c.Get("PowerOnS2Servers", params)
}

func (c *Client) PowerOffS2Servers(params qingcloud.Params) ([]byte, error) {
	return c.Get("PowerOffS2Servers", params)
}

func (c *Client) UpdateS2Servers(params qingcloud.Params) ([]byte, error) {
	return c.Get("UpdateS2Servers", params)
}

func (c *Client) ChangeS2ServerVxnet(params qingcloud.Params) ([]byte, error) {
	return c.Get("ChangeS2ServerVxnet", params)
}

func (c *Client) CreateS2SharedTarget(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateS2SharedTarget", params)
}

func (c *Client) DescribeS2SharedTargets(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeS2SharedTargets", params)
}

func (c *Client) DeleteS2SharedTargets(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteS2SharedTargets", params)
}

func (c *Client) EnableS2SharedTargets(params qingcloud.Params) ([]byte, error) {
	return c.Get("EnableS2SharedTargets", params)
}

func (c *Client) DisableS2SharedTargets(params qingcloud.Params) ([]byte, error) {
	return c.Get("DisableS2SharedTargets", params)
}

func (c *Client) ModifyS2SharedTargets(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyS2SharedTargets", params)
}
func (c *Client) AttachToS2SharedTarget(params qingcloud.Params) ([]byte, error) {
	return c.Get("AttachToS2SharedTarget", params)
}

func (c *Client) DetachFromS2SharedTarget(params qingcloud.Params) ([]byte, error) {
	return c.Get("DetachFromS2SharedTarget", params)
}

func (c *Client) DescribeS2DefaultParameters(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeS2DefaultParameters", params)
}
