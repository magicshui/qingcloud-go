package rdb

import (
	"github.com/magicshui/qingcloud-go"
)

type RDB struct {
	*qingcloud.Client
}

func NewClient(clt *qingcloud.Client) *RDB {
	var a = &RDB{
		Client: clt,
	}
	return a
}

type DescribeRDBsRequest struct {
	RdbsN     qingcloud.NumberedString
	RdbEngine qingcloud.String
	StatusN   qingcloud.NumberedString
	RdbName   qingcloud.String

	TagsN   qingcloud.NumberedString
	Verbose qingcloud.Integer
	Offset  qingcloud.Integer
	Limit   qingcloud.Integer
}
type DescribeRDBsResponse struct {
	TotalCount int   `json:"total_count"`
	RdbSet     []Rdb `json:"rdb_set"`
	qingcloud.CommonResponse
}

// DescribeRDBs 获取一个或多个数据库集群信息。
// 可根据数据库集群 ID，状态，数据库集群名称作过滤条件，来获取数据库集群列表。 如果不指定任何过滤条件，默认返回你所拥有的所有数据库集群。 如果指定不存在的路由器ID，或非法状态值，则会返回错误信息。
func (c *RDB) DescribeRDBs(params DescribeRDBsRequest) (DescribeRDBsResponse, error) {
	var result DescribeRDBsResponse
	err := c.Get("DescribeRDBs", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type CreateRDBRequest struct {
	Vxnet          qingcloud.String
	RdbEngine      qingcloud.String
	EngineVersion  qingcloud.String
	RdbUsername    qingcloud.String
	RdbPassword    qingcloud.String
	RbdType        qingcloud.Integer
	StorageSize    qingcloud.Integer
	DdbName        qingcloud.String
	PrivateIps     qingcloud.Dict
	Description    qingcloud.String
	AutoBackupTime qingcloud.Integer
}

type CreateRDBResponse struct {
	Rdb string `json:"rdb"`
	qingcloud.CommonResponse
}

// CreateRDB
// 创建一个数据库集群。
func (c *RDB) CreateRDB(params CreateRDBRequest) (CreateRDBResponse, error) {
	var result CreateRDBResponse
	err := c.Get("CreateRDB", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DeleteRDBsRequest struct {
	RdbsN qingcloud.NumberedString
}
type DeleteRDBsResponse qingcloud.CommonResponse

// DeleteRDBs
// 删除一个数据库集群。
func (c *RDB) DeleteRDBs(params DeleteRDBsRequest) (DeleteRDBsResponse, error) {
	var result DeleteRDBsResponse
	err := c.Get("DeleteRDBs", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type StartRDBsRequest struct {
	RdbsN qingcloud.NumberedString
}
type StartRDBsResponse struct {
	qingcloud.CommonResponse
	Rdbs []string `json:"rdbs"`
}

// StartRDBs
// 启动指定的数据库集群。
func (c *RDB) StartRDBs(params StartRDBsRequest) (StartRDBsResponse, error) {
	var result StartRDBsResponse
	err := c.Get("StartRDBs", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type StopRDBsRequest struct {
	RdbsN qingcloud.NumberedString
}
type StopRDBsResponse struct {
	qingcloud.CommonResponse
	Rdbs []string `json:"rdbs"`
}

// StopRDBs
// 关闭指定的数据库集群。
func (c *RDB) StopRDBs(params StopRDBsRequest) (StopRDBsResponse, error) {
	var result StopRDBsResponse
	err := c.Get("StopRDBs", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ResizeRDBsRequest struct {
	RdbsN       qingcloud.NumberedString
	RdbType     qingcloud.String
	StorageSize qingcloud.Integer
}
type ResizeRDBsResponse struct {
	qingcloud.CommonResponse
	Rdbs []string `json:"rdbs"`
}

// ResizeRDBs
// 扩容指定的数据库集群。
func (c *RDB) ResizeRDBs(params ResizeRDBsRequest) (ResizeRDBsResponse, error) {
	var result ResizeRDBsResponse
	err := c.Get("ResizeRDBs", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type RDBsLeaveVxnetRequest struct {
	RdbsN qingcloud.NumberedString
	Vxnet qingcloud.String
}
type RDBsLeaveVxnetResponse struct {
	qingcloud.CommonResponse
	Rdbs []string `json:"rdbs"`
}

// RDBsLeaveVxnet¶
// 将指定的数据库集群从私有网络中脱离。
func (c *RDB) RDBsLeaveVxnet(params RDBsLeaveVxnetRequest) (RDBsLeaveVxnetResponse, error) {
	var result RDBsLeaveVxnetResponse
	err := c.Get("RDBsLeaveVxnet", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type RDBsJoinVxnetRequest struct {
	RdbsN qingcloud.NumberedString
	Vxnet qingcloud.String
}
type RDBsJoinVxnetResponse struct {
	qingcloud.CommonResponse
	Rdbs []string `json:"rdbs"`
}

func (c *RDB) RDBsJoinVxnet(params RDBsJoinVxnetRequest) (RDBsJoinVxnetResponse, error) {
	var result RDBsJoinVxnetResponse
	err := c.Get("RDBsJoinVxnet", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type CreateRDBFromSnapshotRequest struct {
	Snapshot       qingcloud.String
	Vxnet          qingcloud.String
	RdbType        qingcloud.Integer
	RdbUsername    qingcloud.String
	RdbPassword    qingcloud.String
	RdbName        qingcloud.String
	PrivateIps     qingcloud.Dict
	Description    qingcloud.String
	AutoBackupTime qingcloud.Integer
}
type CreateRDBFromSnapshotResponse struct {
	qingcloud.CommonResponse
	Rdb string `json:"rdb"`
}

// CreateRDBFromSnapshot
// 从指定备份创建出一个全新的数据库集群。
func (c *RDB) CreateRDBFromSnapshot(params CreateRDBFromSnapshotRequest) (CreateRDBFromSnapshotResponse, error) {
	var result CreateRDBFromSnapshotResponse
	err := c.Get("CreateRDBFromSnapshot", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type CreateTempRDBInstanceFromSnapshotRequest struct {
	Rdb      qingcloud.String
	Snapshot qingcloud.String
}
type CreateTempRDBInstanceFromSnapshotResponse struct {
	qingcloud.CommonResponse
	Rdb string `json:"rdb"`
}

// CreateTempRDBInstanceFromSnapshot
// 从备份创建一个临时性数据库实例，并将之添加到指定的数据库集群。
func (c *RDB) CreateTempRDBInstanceFromSnapshot(params CreateTempRDBInstanceFromSnapshotRequest) (CreateTempRDBInstanceFromSnapshotResponse, error) {
	var result CreateTempRDBInstanceFromSnapshotResponse
	err := c.Get("CreateTempRDBInstanceFromSnapshot", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type GetRDBInstanceFilesRequest struct {
	RdbInstance qingcloud.String
}
type GetRDBInstanceFilesResponse struct {
	Files File `json:"files"`
	qingcloud.CommonResponse
}

func (c *RDB) GetRDBInstanceFiles(params GetRDBInstanceFilesRequest) (GetRDBInstanceFilesResponse, error) {
	var result GetRDBInstanceFilesResponse
	err := c.Get("GetRDBInstanceFiles", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type CopyRDBInstanceFilesToFTPRequest struct {
	RdbInstance qingcloud.String
	FilesN      qingcloud.NumberedString
}
type CopyRDBInstanceFilesToFTPResponse struct {
	qingcloud.CommonResponse
	RdbInstance string `json:"rdb_instance"`
}

// CopyRDBInstanceFilesToFTP 将指定的日志文件拷贝到 FTP 目录。
func (c *RDB) CopyRDBInstanceFilesToFTP(params CopyRDBInstanceFilesToFTPRequest) (CopyRDBInstanceFilesToFTPResponse, error) {
	var result CopyRDBInstanceFilesToFTPResponse
	err := c.Get("CopyRDBInstanceFilesToFTP", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type CeaseRDBInstanceRequest struct {
	Rdb         qingcloud.String
	RdbInstance qingcloud.String
}
type CeaseRDBInstanceResponse struct {
	qingcloud.CommonResponse
	Rdb         string `json:"rdb"`
	RdbInstance string `json:"rdb_instance"`
}

// CeaseRDBInstance  销毁指定的数据库实例，目前只能销毁临时性数据库实例。
func (c *RDB) CeaseRDBInstance(params CeaseRDBInstanceRequest) (CeaseRDBInstanceResponse, error) {
	var result CeaseRDBInstanceResponse
	err := c.Get("CeaseRDBInstance", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ModifyRDBParametersRequest struct {
	Rdb        qingcloud.String
	Parameters qingcloud.Dict
}
type ModifyRDBParametersResponse struct {
	qingcloud.CommonResponse
	Rdb string `json:"rdb" bson:"rdb"`
}

// ModifyRDBParameters 修改指定 RDB 的配置
func (c *RDB) ModifyRDBParameters(params ModifyRDBParametersRequest) (ModifyRDBParametersResponse, error) {
	var result ModifyRDBParametersResponse
	err := c.Get("ModifyRDBParameters", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ApplyRDBParameterGroupRequest struct {
	Rdb qingcloud.String
}
type ApplyRDBParameterGroupResponse struct {
	qingcloud.CommonResponse
	Rdb string `json:"rdb"`
}

// ApplyRDBParameterGroup 应用数据库配置。
func (c *RDB) ApplyRDBParameterGroup(params ApplyRDBParameterGroupRequest) (ApplyRDBParameterGroupResponse, error) {
	var result ApplyRDBParameterGroupResponse
	err := c.Get("ApplyRDBParameterGroup", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DescribeRDBParametersRequest struct {
	Rdb    qingcloud.String
	Offset qingcloud.Integer
	Limit  qingcloud.Integer
}
type DescribeRDBParametersResponse struct {
	qingcloud.CommonResponse
	TotalCount   int                `json:"total_count"`
	ParameterSet []RDBParameterItem `json:"parameter_set"`
}

// DescribeRDBParameters 获取指定 RDB 的配置信息。
func (c *RDB) DescribeRDBParameters(params DescribeRDBParametersRequest) (DescribeRDBParametersResponse, error) {
	var result DescribeRDBParametersResponse
	err := c.Get("DescribeRDBParameters", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}
