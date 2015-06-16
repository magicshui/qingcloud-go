package loadbalancer

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) AddLoadBalancerBackends(params qingcloud.Params) ([]byte, error) {
	return c.Get("AddLoadBalancerBackends", params)
}

func (c *Client) DescribeLoadBalancers(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeLoadBalancers", params)
}

func (c *Client) DeleteLoadBalancers(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteLoadBalancers", params)
}

func (c *Client) ModifyLoadBalancerAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyLoadBalancerAttributes", params)
}

func (c *Client) StartLoadBalancers(params qingcloud.Params) ([]byte, error) {
	return c.Get("StartLoadBalancers", params)
}

func (c *Client) StopLoadBalancers(params qingcloud.Params) ([]byte, error) {
	return c.Get("StopLoadBalancers", params)
}

func (c *Client) UpdateLoadBalancers(params qingcloud.Params) ([]byte, error) {
	return c.Get("UpdateLoadBalancers", params)
}

func (c *Client) ResizeLoadBalancers(params qingcloud.Params) ([]byte, error) {
	return c.Get("ResizeLoadBalancers", params)
}

func (c *Client) AssociateEipsToLoadBalancer(params qingcloud.Params) ([]byte, error) {
	return c.Get("AssociateEipsToLoadBalancer", params)
}

func (c *Client) DissociateEipsFromLoadBalancer(params qingcloud.Params) ([]byte, error) {
	return c.Get("DissociateEipsFromLoadBalancer", params)
}

func (c *Client) AddLoadBalancerListeners(params qingcloud.Params) ([]byte, error) {
	return c.Get("AddLoadBalancerListeners", params)
}

func (c *Client) DescribeLoadBalancerListeners(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeLoadBalancerListeners", params)
}

func (c *Client) DeleteLoadBalancerListeners(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteLoadBalancerListeners", params)
}

func (c *Client) ModifyLoadBalancerListenerAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyLoadBalancerListenerAttributes", params)
}

func (c *Client) AddLoadBalancerBackends(params qingcloud.Params) ([]byte, error) {
	return c.Get("AddLoadBalancerBackends", params)
}

func (c *Client) DescribeLoadBalancerBackends(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeLoadBalancerBackends", params)
}

func (c *Client) DeleteLoadBalancerBackends(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteLoadBalancerBackends", params)
}

func (c *Client) ModifyLoadBalancerBackendAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyLoadBalancerBackendAttributes", params)
}

func (c *Client) CreateLoadBalancerPolicy(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateLoadBalancerPolicy", params)
}

func (c *Client) DescribeLoadBalancerPolicies(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeLoadBalancerPolicies", params)
}

func (c *Client) ModifyLoadBalancerPolicyAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyLoadBalancerPolicyAttributes", params)
}

func (c *Client) ApplyLoadBalancerPolicy(params qingcloud.Params) ([]byte, error) {
	return c.Get("ApplyLoadBalancerPolicy", params)
}

func (c *Client) DeleteLoadBalancerPolicies(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteLoadBalancerPolicies", params)
}

func (c *Client) AddLoadBalancerPolicyRules(params qingcloud.Params) ([]byte, error) {
	return c.Get("AddLoadBalancerPolicyRules", params)
}

func (c *Client) DescribeLoadBalancerPolicyRules(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeLoadBalancerPolicyRules", params)
}

func (c *Client) ModifyLoadBalancerPolicyRuleAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyLoadBalancerPolicyRuleAttributes", params)
}

func (c *Client) DeleteLoadBalancerPolicyRules(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteLoadBalancerPolicyRules", params)
}
