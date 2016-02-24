package qingcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	zone            string `json:"zone"`
	accessKeyID     string `json:"access_key_id"`
	secretAccessKey string `json:"secret_access_key"`
	params          Params
	commonParams    Params
}

func NewClient() *Client {
	return &Client{
		params:       Params{},
		commonParams: Params{},
	}
}

func (c *Client) ConnectToZone(zone string, id string, secret string) {
	c.zone = zone
	c.accessKeyID = id
	c.secretAccessKey = secret
	c.commonParams = []*Param{
		{"version", 1},
		{"signature_method", "HmacSHA256"},
		{"signature_version", 1},
		{"access_key_id", id},
		{"zone", zone},
	}

}

func (c *Client) addAction(action string) {
	c.params = append(c.params, &Param{
		"action", action,
	})

}

func (c *Client) addTimeStamp() {
	c.params = append(c.params, &Param{
		"time_stamp", time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	})
}

func (c *Client) getUrl(httpMethod string) (string, string) {
	for _, v := range c.commonParams {
		c.params = append(c.params, v)
	}
	sortParamsByKey(c.params)
	// log.Printf("%v", c.params)
	urlEscapeParams(c.params)
	// log.Printf("%v", c.params)
	url := generateUrlByParams(c.params)
	url2 := genSignatureUrl(httpMethod, `/iaas/`, url)
	sig := genSignature(url2, c.secretAccessKey)
	// log.Printf("url:%s", url)
	// log.Printf("url2: %s", url2)
	// log.Printf("sig: %s", sig)
	return url, sig
}

func (c *Client) Get(action string, params Params, response interface{}) error {
	result, err := c.get(action, params)
	if err != nil {
		return err
	}
	// log.Printf("%s", string(result))
	err = json.Unmarshal(result, &response)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) get(action string, params Params) ([]byte, error) {
	var _p []*Param
	c.params = _p
	for i, _ := range params {
		c.params = append(c.params, params[i])
	}
	c.addAction(action)
	c.addTimeStamp()
	_url, _sig := c.getUrl("GET")
	url := fmt.Sprintf("https://api.qingcloud.com/iaas/?%v&signature=%v", _url, _sig)
	log.Printf("URL:\n %s", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}
