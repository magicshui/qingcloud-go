package cache

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeCaches(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeCaches", params)
}
func (c *Client) CreateCache(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateCache", params)
}
func (c *Client) StopCaches(params qingcloud.Params) ([]byte, error) {
	return c.Get("StopCaches", params)
}

func (c *Client) StartCaches(params qingcloud.Params) ([]byte, error) {
	return c.Get("StartCaches", params)
}
func (c *Client) RestartCaches(params qingcloud.Params) ([]byte, error) {
	return c.Get("RestartCaches", params)
}
func (c *Client) DeleteCaches(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteCaches", params)
}
func (c *Client) ResizeCaches(params qingcloud.Params) ([]byte, error) {
	return c.Get("ResizeCaches", params)
}

func (c *Client) UpdateCache(params qingcloud.Params) ([]byte, error) {
	return c.Get("UpdateCache", params)
}

func (c *Client) ChangeCacheVxnet(params qingcloud.Params) ([]byte, error) {
	return c.Get("ChangeCacheVxnet", params)
}

func (c *Client) ModifyCacheAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyCacheAttributes", params)
}

func (c *Client) DescribeCacheNodes(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeCacheNodes", params)
}

func (c *Client) AddCacheNodes(params qingcloud.Params) ([]byte, error) {
	return c.Get("AddCacheNodes", params)
}

func (c *Client) DeleteCacheNodes(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteCacheNodes", params)
}

func (c *Client) RestartCacheNodes(params qingcloud.Params) ([]byte, error) {
	return c.Get("RestartCacheNodes", params)
}

func (c *Client) ModifyCacheNodeAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyCacheNodeAttributes", params)
}
func (c *Client) CreateCacheFromSnapshot(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateCacheFromSnapshot", params)
}

func (c *Client) DescribeCacheParameterGroups(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeCacheParameterGroups", params)
}

func (c *Client) CreateCacheParameterGroup(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateCacheParameterGroup", params)
}

func (c *Client) ApplyCacheParameterGroup(params qingcloud.Params) ([]byte, error) {
	return c.Get("ApplyCacheParameterGroup", params)
}

func (c *Client) DeleteCacheParameterGroups(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteCacheParameterGroups", params)
}

func (c *Client) ModifyCacheParameterGroupAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifyCacheParameterGroupAttributes", params)
}

func (c *Client) DescribeCacheParameters(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeCacheParameters", params)
}

func (c *Client) UpdateCacheParameters(params qingcloud.Params) ([]byte, error) {
	return c.Get("UpdateCacheParameters", params)
}

func (c *Client) ResetCacheParameters(params qingcloud.Params) ([]byte, error) {
	return c.Get("ResetCacheParameters", params)
}
