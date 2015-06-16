package zone

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) CreateRDB(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateRDB", params)
}
func (c *Client) DescribeRDBs(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeRDBs", params)
}

func (c *Client) DeleteRDBs(params qingcloud.Params) ([]byte, error) {
	return c.Get("DeleteRDBs", params)
}
func (c *Client) StartRDBs(params qingcloud.Params) ([]byte, error) {
	return c.Get("StartRDBs", params)
}
func (c *Client) StopRDBs(params qingcloud.Params) ([]byte, error) {
	return c.Get("StopRDBs", params)
}
func (c *Client) ResizeRDBs(params qingcloud.Params) ([]byte, error) {
	return c.Get("ResizeRDBs", params)
}

func (c *Client) RDBsLeaveVxnet(params qingcloud.Params) ([]byte, error) {
	return c.Get("RDBsLeaveVxnet", params)
}
func (c *Client) RDBsJoinVxnet(params qingcloud.Params) ([]byte, error) {
	return c.Get("RDBsJoinVxnet", params)
}
func (c *Client) CreateRDBFromSnapshot(params qingcloud.Params) ([]byte, error) {
	return c.Get("CreateRDBFromSnapshot", params)
}
func (c *Client) GetRDBInstanceFiles(params qingcloud.Params) ([]byte, error) {
	return c.Get("GetRDBInstanceFiles", params)
}

func (c *Client) CopyRDBInstanceFilesToFTP(params qingcloud.Params) ([]byte, error) {
	return c.Get("CopyRDBInstanceFilesToFTP", params)
}
func (c *Client) GetRDBMonitor(params qingcloud.Params) ([]byte, error) {
	return c.Get("GetRDBMonitor", params)
}
