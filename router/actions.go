package router

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeRouters(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeRouters", params)
}
func (c *Client) CreateRouters(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateRouters", params)
}
func (c *Client) DeleteRouters(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteRouters", params)
}
func (c *Client) UpdateRouters(params qingcloud.Params) ([]byte, error) {
	return c.Get("UpdateRouters", params)
}
func (c *Client) PowerOffRouters(params qingcloud.Params) ([]byte, error) {
	return c.Get("PowerOffRouters", params)
}

func (c *Client) PowerOnRouters(params qingcloud.Params) ([]byte, error) {
	return c.Get("PowerOnRouters", params)
}
func (c *Client) JoinRouter(params qingcloud.Params) ([]byte, error) {
	return c.Get("JoinRouter", params)
}
func (c *Client) LeaveRouter(params qingcloud.Params) ([]byte, error) {
	return c.Get("LeaveRouter", params)
}
func (c *Client) ModifyRouterAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyRouterAttributes", params)
}
func (c *Client) DescribeRouterStatics(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeRouterStatics", params)
}
func (c *Client) AddRouterStatics(params qingcloud.Params) ([]byte, error) {
	return c.Get("AddRouterStatics", params)
}
func (c *Client) ModifyRouterStaticAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyRouterStaticAttributes", params)
}
func (c *Client) DeleteRouterStatics(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteRouterStatics", params)
}
func (c *Client) DescribeRouterVxnets(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeRouterVxnets", params)
}
func (c *Client) AddRouterStaticEntries(params qingcloud.Params) ([]byte, error) {
	return c.Get("AddRouterStaticEntries", params)
}

func (c *Client) DeleteRouterStaticEntries(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteRouterStaticEntries", params)
}
func (c *Client) ModifyRouterStaticEntryAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyRouterStaticEntryAttributes", params)
}
func (c *Client) DescribeRouterStaticEntries(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeRouterStaticEntries", params)
}
