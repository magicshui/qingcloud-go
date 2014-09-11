package client

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 *
 *
 *
action   每个 API 都有自己的 action，用来标识所请求指令。例如 RunInstances。  完整的指令名称列表请参见 API指令列表。
time_stamp  请求串生成时间，格式为 YYYY-MM-DDThh:mm:ssZ，例如”2013-08-27T13:58:35Z”，具体格式可以参见 ISO8601. 这个时间为 UTC 时间，假设您的本地时间为北京时间 UTC+8 ，您需要将其转化为 UTC+0 的时间。
signature   请求消息的签名，请参见 签名方法
*/
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

	/*
	 c.Params = append(c.Params, &Param{
	        "time_stamp", "2013-08-27T14:30:10Z",
	    })
	    time.Now()

	*/

	c.Params = append(c.Params, &Param{
		"time_stamp", time.Now().Add(-time.Hour * 8).Format("2006-01-02T15:04:05Z"),
	})
	// 2013-08-27T14%3A30%3A10Z time.Now().Format("2006-01-02T15:04:05Z")

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

func (c *Client) Get(action string, params Params) string {
	var _p []*Param
	c.Params = _p
	for i, _ := range params {
		c.Params = append(c.Params, params[i])
	}
	c.addAction(action)
	c.addTimeStamp()
	_url, _sig := c.getUrl("GET")
	url := fmt.Sprintf("https://api.qingcloud.com/iaas/?%v&signature=%v", _url, _sig)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//fmt.Println("读取错误：", err)
		return ""
	}
	result_string := string(result)
	return result_string
}
