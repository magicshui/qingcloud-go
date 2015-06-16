package rdb

import (
	"github.com/magicshui/qingcloud-go"
)

type Client struct {
	*qingcloud.Client
}

func NewClient(zone string, id string, secret string) *Client {
	var c qingcloud.Client
	c.ConnectToZone(zone, id, secret)
	var clt Client
	clt.Client = &c
	return &clt

}
