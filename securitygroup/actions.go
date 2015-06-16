package securitygroup

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeSecurityGroups(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeSecurityGroups", params)
}
func (c *Client) CreateSecurityGroup(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateSecurityGroup", params)
}
func (c *Client) DeleteSecurityGroups(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteSecurityGroups", params)
}
func (c *Client) ApplySecurityGroup(params qingcloud.Params) ([]byte, error) {
	return c.Get("ApplySecurityGroup", params)
}
func (c *Client) ModifySecurityGroupAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifySecurityGroupAttributes", params)
}
func (c *Client) DescribeSecurityGroupRules(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeSecurityGroupRules", params)
}
func (c *Client) AddSecurityGroupRules(params qingcloud.Params) ([]byte, error) {
	return c.Get("AddSecurityGroupRules", params)
}
func (c *Client) DeleteSecurityGroupRules(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteSecurityGroupRules", params)
}
func (c *Client) ModifySecurityGroupRuleAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifySecurityGroupRuleAttributes", params)
}
