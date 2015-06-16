package eip

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeEips(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeEips", params)
}

func (c *Client) AllocateEips(params qingcloud.Params) ([]byte, error) {
	return c.Get("AllocateEips", params)
}

func (c *Client) ReleaseEips(params qingcloud.Params) ([]byte, error) {
	return c.Get("ReleaseEips", params)
}

func (c *Client) AssociateEip(params qingcloud.Params) ([]byte, error) {
	return c.Get("AssociateEip", params)
}

func (c *Client) DissociateEips(params qingcloud.Params) ([]byte, error) {
	return c.Get("DissociateEips", params)
}
func (c *Client) ChangeEipsBandwidth(params qingcloud.Params) ([]byte, error) {
	return c.Get("ChangeEipsBandwidth", params)
}
func (c *Client) ChangeEipsBillingMode(params qingcloud.Params) ([]byte, error) {
	return c.Get("ChangeEipsBillingMode", params)
}
func (c *Client) ModifyEipAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyEipAttributes", params)
}
