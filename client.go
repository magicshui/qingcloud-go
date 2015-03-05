package qingcloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/magicshui/goutils/requests"
	"time"
)

type Client struct {
	Zone            string `json:"zone"`
	AccessKeyId     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	Params          Params
	CommonParams    Params
}

func (c *Client) ConnectToZone(zone string, id string, secret string) {
	c.Zone = zone
	c.AccessKeyId = id
	c.SecretAccessKey = secret
	c.CommonParams = []*Param{
		{"version", 1},
		{"signature_method", "HmacSHA256"},
		{"signature_version", 1},
		{"access_key_id", id},
		{"zone", zone},
	}

}

func (c *Client) addAction(action string) {
	c.Params = append(c.Params, &Param{
		"action", action,
	})

}

func (c *Client) addTimeStamp() {

	c.Params = append(c.Params, &Param{
		"time_stamp", time.Now().Add(-time.Hour * 8).Format("2006-01-02T15:04:05Z"),
	})

}

func (c *Client) getUrl(httpMethod string) (string, string) {
	for _, v := range c.CommonParams {
		c.Params = append(c.Params, v)
	}
	sortParamsByKey(c.Params)
	urlEscapeParams(c.Params)
	url := generateUrlByParams(c.Params)
	url2 := genSignatureUrl(httpMethod, "/iaas/", url)
	sig := genSignature(url2, c.SecretAccessKey)
	return url, sig
}

func (c *Client) Get(action string, params Params) ([]byte, error) {
	var _p []*Param
	c.Params = _p
	for i, _ := range params {
		c.Params = append(c.Params, params[i])
	}
	c.addAction(action)
	c.addTimeStamp()
	_url, _sig := c.getUrl("GET")
	url := fmt.Sprintf("https://api.qingcloud.com/iaas/?%v&signature=%v", _url, _sig)

	result, err := requests.GetHttpsRequest(url)

	if err != nil {
		return nil, err
	}

	var hasError QCError
	err = json.Unmarshal(result, &hasError)
	if err != nil {
		return nil, err
	}
	if hasError.RetCode > 0 {
		return nil, errors.New(hasError.Message)
	}
	return result, nil
}
