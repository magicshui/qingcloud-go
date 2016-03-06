package qingcloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	zone            string `json:"zone"`
	accessKeyID     string `json:"access_key_id"`
	secretAccessKey string `json:"secret_access_key"`
	params          Params
	commonParams    Params
	l               sync.Mutex
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
	urlEscapeParams(c.params)
	url := generateUrlByParams(c.params)
	url2 := genSignatureUrl(httpMethod, `/iaas/`, url)
	sig := genSignature(url2, c.secretAccessKey)
	return url, sig
}

func (c *Client) Get(action string, params Params, response interface{}) error {
	c.l.Lock()
	defer c.l.Unlock()
	result, err := c.get(action, params)
	if err != nil {
		log.Printf("Get Error %s , %#v   %s", err, params, string(result))
		return err
	}
	var errCode CommonResponse
	err = json.Unmarshal(result, &errCode)
	if err != nil {
		log.Printf("Get Error Unmashl %s , %#v   %s", err, params, string(result))
		return err
	}

	if errCode.RetCode != 0 {
		return errors.New(errCode.Message)
	}

	log.Printf("%s", string(result))

	err = json.Unmarshal(result, &response)
	if err != nil {
		log.Printf("Get Error Unmashl to Response %s , %#v   %s", err, params, string(result))
		return err
	}
	return nil
}
func (c *Client) Post(action string, params Params, response interface{}) error {
	c.l.Lock()
	defer c.l.Unlock()
	result, err := c.post(action, params)
	if err != nil {
		return err
	}
	var errCode CommonResponse
	json.Unmarshal(result, &errCode)
	if errCode.RetCode != 0 {
		return errors.New(errCode.Message)
	}

	log.Printf("%s", string(result))

	err = json.Unmarshal(result, &response)
	if err != nil {
		return err
	}
	return nil
}

// TODO: fix this
func (c *Client) post(action string, params Params) ([]byte, error) {
	var _p []*Param
	c.params = _p
	for i, _ := range params {
		c.params = append(c.params, params[i])
	}
	c.addAction(action)
	c.addTimeStamp()
	_url, _sig := c.getUrl("Post")
	url := fmt.Sprintf("https://api.qingcloud.com/iaas/?%v&signature=%v", _url, _sig)
	log.Printf("%s", url)
	res, err := http.Post(url, "application/json;utf-8", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
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
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
