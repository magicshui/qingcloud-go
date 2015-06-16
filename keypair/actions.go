package keypair

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeKeyPairs(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeKeyPairs", params)
}
func (c *Client) CreateKeyPair(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateKeyPair", params)
}
func (c *Client) DeleteKeyPairs(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteKeyPairs", params)
}
func (c *Client) AttachKeyPairs(params qingcloud.Params) ([]byte, error) {
	return c.Get("AttachKeyPairs", params)
}
func (c *Client) DetachKeyPairs(params qingcloud.Params) ([]byte, error) {
	return c.Get("DetachKeyPairs", params)
}
func (c *Client) ModifyKeyPairAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyKeyPairAttributes", params)
}
