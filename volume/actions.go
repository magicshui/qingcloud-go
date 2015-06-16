package instance

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeVolumes(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeVolumes", params)
}

func (c *Client) CreateVolumes(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateVolumes", params)
}

func (c *Client) DeleteVolumes(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteVolumes", params)
}

func (c *Client) AttachVolumes(params qingcloud.Params) ([]byte, error) {
	return c.Get("AttachVolumes", params)
}

func (c *Client) DetachVolumes(params qingcloud.Params) ([]byte, error) {
	return c.Get("DetachVolumes", params)
}

func (c *Client) ResizeVolumes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ResizeVolumes", params)
}

func (c *Client) ModifyVolumeAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyVolumeAttributes", params)
}
