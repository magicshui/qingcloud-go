package image

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeImages(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeImages", params)
}
func (c *Client) CaptureInstance(params qingcloud.Params) ([]byte, error) {
	return c.Get("CaptureInstance", params)
}
func (c *Client) DeleteImages(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteImages", params)
}
func (c *Client) ModifyImageAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyImageAttributes", params)
}
func (c *Client) GrantImageToUsers(params qingcloud.Params) ([]byte, error) {
	return c.Get("GrantImageToUsers", params)
}
func (c *Client) RevokeImageFromUsers(params qingcloud.Params) ([]byte, error) {
	return c.Get("RevokeImageFromUsers", params)
}
