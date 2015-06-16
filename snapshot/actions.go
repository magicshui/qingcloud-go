package snapshot

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeSnapshots(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeSnapshots", params)
}

func (c *Client) CreateSnapshots(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateSnapshots", params)
}
func (c *Client) DeleteSnapshots(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteSnapshots", params)
}
func (c *Client) ApplySnapshots(params qingcloud.Params) ([]byte, error) {
	return c.Get("ApplySnapshots", params)
}

func (c *Client) ModifySnapshotAttributes(params qingcloud.Params) ([]byte, error) {
	return c.Get("ModifySnapshotAttributes", params)
}
func (c *Client) CaptureInstanceFromSnapshot(params qingcloud.Params) ([]byte, error) {
	return c.Get("CaptureInstanceFromSnapshot", params)
}
func (c *Client) CreateVolumeFromSnapshot(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateVolumeFromSnapshot", params)
}
