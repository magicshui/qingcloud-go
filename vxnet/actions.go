package vxnet

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeVxnets(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeVxnets", params)
}

func (c *Client) CreateVxnets(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateVxnets", params)
}

func (c *Client) DeleteVxnets(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteVxnets", params)
}

func (c *Client) JoinVxnet(params qingcloud.Params) ([]byte, error) {
	return c.Get("JoinVxnet", params)
}

func (c *Client) LeaveVxnet(params qingcloud.Params) ([]byte, error) {
	return c.Get("LeaveVxnet", params)
}

func (c *Client) ModifyVxnetAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyVxnetAttributes", params)
}

func (c *Client) DescribeVxnetInstances(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeVxnetInstances", params)
}
