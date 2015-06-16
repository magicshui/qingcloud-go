package router

import (
	"encoding/json"
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeRouters(routers string, status string, search_word string, verbose int, offset int, limit int) ([]Router, error) {
	params := []*qingcloud.Param{
		{"routers.1", routers},
		{"status.1", status},
		{"search_word", search_word},
		{"verbose", verbose},
		{"offset", offset},
		{"limit", limit},
	}

	result, err := c.Get("DescribeRouters", params)
	if err != nil {
		return nil, err
	}
	var response struct {
		routers []Router `json:"router_statics"`
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	return response.routers, nil
}

func (c *Client) UpdateRouters(routers string) error {
	params := []*qingcloud.Param{
		{"routers.1", routers},
	}

	result, err := c.Get("UpdateRouters", params)
	if err != nil {
		return err
	}
	var response struct {
		RouterStatics []string `json:"router_statics"`
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return err
	}
	return nil

}
func (c *Client) AddRouterStatics(router string, router_static_name string, static_type int, val1, val2, val3, val4 string) (string, error) {
	params := []*qingcloud.Param{
		{"router", router},
		{"statics.1.router_static_name", router_static_name},
		{"statics.1.static_type", static_type},
		{"statics.1.val1", val1},
		{"statics.1.val2", val2},
		{"statics.1.val3", val3},
		{"statics.1.val4", val4},
	}

	result, err := c.Get("AddRouterStatics", params)
	if err != nil {
		return "", err
	}
	var response struct {
		RouterStatics []string `json:"router_statics"`
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return "", err
	}
	return response.RouterStatics[0], nil
}

func (c *Client) DeleteRouterStatics(router_statics string) (bool, error) {
	params := []*qingcloud.Param{
		{"router_statics.1", router_statics},
	}

	result, err := c.Get("DeleteRouterStatics", params)
	if err != nil {
		return false, err
	}
	var response struct {
		RouterStatics []string `json:"router_statics"`
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	return response.RouterStatics[0] == router_statics, nil

}
