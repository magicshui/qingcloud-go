package hadoop

import (
	"github.com/magicshui/qingcloud-go"
)

type HADOOP struct {
	*qingcloud.Client
}

func NewClient(clt *qingcloud.Client) *HADOOP {
	var a = &HADOOP{
		Client: clt,
	}
	return a
}

type AddHadoopNodesRequest struct {
	Hadoop                qingcloud.String
	NodeCount             qingcloud.Integer
	HadoopNodeName        qingcloud.String
	PrivateIPsNRole       qingcloud.NumberedString
	PrivateIPsNPrivateIPs qingcloud.NumberedString
}
type AddHadoopNodesResponse struct {
	qingcloud.CommonResponse
	HadoopID         string   `json:"hadoop_id"`
	HadoopNewNodeIDs []string `json:"hadoop_new_node_ids"`
}

// AddHadoopNodes 给 Hadoop 服务添加一个或多个 slave 节点。
func (c *HADOOP) AddHadoopNodes(params AddHadoopNodesRequest) (AddHadoopNodesResponse, error) {
	var result AddHadoopNodesResponse
	err := c.Get("AddHadoopNodes", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DeleteHadoopNodesRequest struct {
	Hadoop       qingcloud.String
	HadoopNodesN qingcloud.NumberedString
}
type DeleteHadoopNodesResponse struct {
	qingcloud.CommonResponse
	HadoopID string `json:"hadoop_id"`
}

// DeleteHadoopNodes 删除 Hadoop 服务 slave 节点。
func (c *HADOOP) DeleteHadoopNodes(params DeleteHadoopNodesRequest) (DeleteHadoopNodesResponse, error) {
	var result DeleteHadoopNodesResponse
	err := c.Get("DeleteHadoopNodes", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type StartHadoopsRequest struct {
	HadoopsN qingcloud.NumberedString
}
type StartHadoopsResponse struct {
	qingcloud.CommonResponse
}

// StartHadoops 启动一个或多个 Hadoop 服务。
func (c *HADOOP) StartHadoops(params StartHadoopsRequest) (StartHadoopsResponse, error) {
	var result StartHadoopsResponse
	err := c.Get("StartHadoops", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type StopHadoopsRequest struct {
	HadoopsN qingcloud.NumberedString
}
type StopHadoopsResponse struct {
	qingcloud.CommonResponse
}

// StopHadoops 关闭一个或多个 Hadoop 服务。该操作将关闭 Hadoop 服务的所有 Hadoop 节点。
func (c *HADOOP) StopHadoops(params StopHadoopsRequest) (StopHadoopsResponse, error) {
	var result StopHadoopsResponse
	err := c.Get("StopHadoops", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}
