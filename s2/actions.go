package s2

import (
	"github.com/magicshui/qingcloud-go"
)

// ROUTER 路由
type S2 struct {
	*qingcloud.Client
}

// NewClient 创建新的路由控制器
func NewClient(clt *qingcloud.Client) *S2 {
	return &S2{
		Client: clt,
	}
}

type CreateS2ServerRequest struct {
	VxnetID      qingcloud.String
	ServiceType  qingcloud.String
	Name         qingcloud.String
	S2ServerType qingcloud.String
	PrivateIP    qingcloud.String
	Description  qingcloud.String
}
type CreateS2ServerResponse struct {
	qingcloud.CommonResponse
	S2ServerID string `json:"s2_server_id"`
}

// CreateS2Server 创建新的共享存储服务器。
// 青云目前仅支持VSAN类型的共享存储服务器， 及基于iSCSI协议的存储服务。
func (c *S2) CreateS2Server(params CreateS2ServerRequest) (CreateS2ServerResponse, error) {
	var result CreateS2ServerResponse
	err := c.Get("CreateS2Server", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DescribeS2ServersRequest struct {
	S2ServersN qingcloud.NumberedString
	StatusN    qingcloud.NumberedString
	SearchWord qingcloud.String
	TagsN      qingcloud.NumberedString
	Verbose    qingcloud.Integer
	Offset     qingcloud.Integer
	Limit      qingcloud.Integer
}
type DescribeS2ServersResponse struct {
	qingcloud.CommonResponse
	TotalCount  int        `json:"total_count"`
	S2ServerSet []S2Server `json:"s2_server_set"`
}

// DescribeS2Servers
func (c *S2) DescribeS2Servers(params DescribeS2ServersRequest) (DescribeS2ServersResponse, error) {
	var result DescribeS2ServersResponse
	err := c.Get("DescribeS2Servers", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ModifyS2ServerRequest struct {
	S2Server    qingcloud.String
	Name        qingcloud.String
	Description qingcloud.String
}
type ModifyS2ServerResponse struct {
	qingcloud.CommonResponse
}

// ModifyS2Server 修改共享存储服务器的属性。
func (c *S2) ModifyS2Server(params ModifyS2ServerRequest) (ModifyS2ServerResponse, error) {
	var result ModifyS2ServerResponse
	err := c.Get("ModifyS2Server", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ResizeS2ServersRequest struct {
	S2Server     qingcloud.String
	S2ServerType qingcloud.Integer
}
type ResizeS2ServersResponse struct {
	qingcloud.CommonResponse
	S2Servers []string `json:"s2_servers"`
}

// ResizeS2Servers 修改共享存储服务器的类型。
func (c *S2) ResizeS2Servers(params ResizeS2ServersRequest) (ResizeS2ServersResponse, error) {
	var result ResizeS2ServersResponse
	err := c.Get("ResizeS2Servers", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DeleteS2ServersRequest struct {
	S2ServersN qingcloud.NumberedString
}
type DeleteS2ServersResponse struct {
	qingcloud.CommonResponse
	S2Servers []string `json:"s2_servers"`
}

// DeleteS2Servers 删除一台或多台共享存储服务器。
// 销毁共享存储服务器的前提，是此共享存储服务器已建立租用信息（租用信息是在创建共享存储服务器成功后， 几秒钟内系统自动建立的）。所以正在创建的共享存储服务器（状态为 pending ）， 以及刚刚创建成功但还没有建立租用信息的共享存储服务器，是不能被销毁的。
func (c *S2) DeleteS2Servers(params DeleteS2ServersRequest) (DeleteS2ServersResponse, error) {
	var result DeleteS2ServersResponse
	err := c.Get("DeleteS2Servers", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type PowerOnS2ServersRequest struct {
	S2ServersN qingcloud.NumberedString
}
type PowerOnS2ServersResponse struct {
	qingcloud.CommonResponse
	S2Servers []string `json:"s2_servers"`
}

// PowerOnS2Servers   启动一台或多台共享存储服务器。
// 共享存储服务器只有在关闭 poweroffed 状态才能被启动，如果处于非关闭状态，则返回错误信息。
func (c *S2) PowerOnS2Servers(params PowerOnS2ServersRequest) (PowerOnS2ServersResponse, error) {
	var result PowerOnS2ServersResponse
	err := c.Get("PowerOnS2Servers", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type PowerOffS2ServersRequest struct {
	S2ServersN qingcloud.NumberedString
}
type PowerOffS2ServersResponse struct {
	qingcloud.CommonResponse
	S2Servers []string `json:"s2_servers"`
}

// PowerOffS2Servers 关闭一台或多台共享存储服务器。
// 共享存储服务器只有在活动 active 状态才能被关闭，如果处于非活动状态，则返回错误信息。
func (c *S2) PowerOffS2Servers(params PowerOffS2ServersRequest) (PowerOffS2ServersResponse, error) {
	var result PowerOffS2ServersResponse
	err := c.Get("PowerOffS2Servers", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type UpdateS2ServersRequest struct {
	S2ServersN qingcloud.NumberedString
}
type UpdateS2ServersResponse struct {
	qingcloud.CommonResponse
}

// UpdateS2Servers 更新一台或多台共享存储服务器的配置信息。当配置发生变更之后，需要执行本操作使配置生效。
// 只有在处于 active 状态的共享存储服务器才能支持此操作，如果处于非活跃状态，则返回错误信息。
func (c *S2) UpdateS2Servers(params UpdateS2ServersRequest) (UpdateS2ServersResponse, error) {
	var result UpdateS2ServersResponse
	err := c.Get("UpdateS2Servers", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ChangeS2ServerVxnetRequest struct {
	S2Server  qingcloud.String
	Vxnet     qingcloud.String
	PrivateIP qingcloud.String
}
type ChangeS2ServerVxnetResponse struct {
	qingcloud.CommonResponse
}

// ChangeS2ServerVxnet  修改共享存储服务器的属性。
func (c *S2) ChangeS2ServerVxnet(params ChangeS2ServerVxnetRequest) (ChangeS2ServerVxnetResponse, error) {
	var result ChangeS2ServerVxnetResponse
	err := c.Get("ChangeS2ServerVxnet", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type CreateS2SharedTargetRequest struct {
	S2Server        qingcloud.String
	ExportName      qingcloud.String
	TargetType      qingcloud.String
	Description     qingcloud.String
	VolumesN        qingcloud.NumberedString
	InitiatorNamesN qingcloud.NumberedString
}
type CreateS2SharedTargetResponse struct {
	qingcloud.CommonResponse
	S2SharedTarget string `json:"s2_shared_target"`
}

// CreateS2SharedTarget 新建共享存储目标，在新建时可以直接添加硬盘作为backstore，也可以以后添加。
// 此操作完成后需要调用 UpdateS2Servers 以应用到共享存储服务器上。
func (c *S2) CreateS2SharedTarget(params CreateS2SharedTargetRequest) (CreateS2SharedTargetResponse, error) {
	var result CreateS2SharedTargetResponse
	err := c.Get("CreateS2SharedTarget", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DescribeS2SharedTargetsRequest struct {
	SharedTargetsN qingcloud.NumberedString
	S2ServerID     qingcloud.String
	SearchWord     qingcloud.String
	Verbose        qingcloud.String
	Offset         qingcloud.Integer
	Limit          qingcloud.Integer
}
type DescribeS2SharedTargetsResponse struct {
	qingcloud.CommonResponse
	TotalCount     int `json:"total_count"`
	S2SharedTarget []struct {
	} `json:"s2_shared_target_set"`
}

// DescribeS2SharedTargets 获取一个或多个共享存储目标
// 可根据共享存储目标ID，共享存储服务器ID，状态，名称作过滤条件，来获取共享存储目标列表。 如果不指定任何过滤条件，默认返回你所拥有的所有共享存储目标。 如果指定不存在的共享存储目标ID，或非法状态值，则会返回错误信息。
func (c *S2) DescribeS2SharedTargets(params DescribeS2SharedTargetsRequest) (DescribeS2SharedTargetsResponse, error) {
	var result DescribeS2SharedTargetsResponse
	err := c.Get("DescribeS2SharedTargets", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DeleteS2SharedTargetsRequest struct {
	SharedTargetsN qingcloud.NumberedString
}
type DeleteS2SharedTargetsResponse struct {
	qingcloud.CommonResponse
	SharedTargets []string `json:"shared_targets"`
}

// DeleteS2SharedTargets 删除一个或多个共享存储目标。
// 此操作完成后需要调用 UpdateS2Servers 以应用到共享存储服务器上。
func (c *S2) DeleteS2SharedTargets(params DeleteS2SharedTargetsRequest) (DeleteS2SharedTargetsResponse, error) {
	var result DeleteS2SharedTargetsResponse
	err := c.Get("DeleteS2SharedTargets", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type EnableS2SharedTargetsRequest struct {
	SharedTargetsN qingcloud.NumberedString
}
type EnableS2SharedTargetsResponse struct {
	qingcloud.CommonResponse
	SharedTargets []string `json:"shared_targets"`
}

// EnableS2SharedTargets 启用一个或多个共享存储目标，此操作可在共享存储目标被禁用后使用。
// 此操作完成后需要调用 UpdateS2Servers 以应用到共享存储服务器上。
func (c *S2) EnableS2SharedTargets(params EnableS2SharedTargetsRequest) (EnableS2SharedTargetsResponse, error) {
	var result EnableS2SharedTargetsResponse
	err := c.Get("EnableS2SharedTargets", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DisableS2SharedTargetsRequest struct {
	SharedTargetsN qingcloud.NumberedString
}
type DisableS2SharedTargetsResponse struct {
	qingcloud.CommonResponse
	SharedTargets []string `json:"shared_targets"`
}

// DisableS2SharedTargets   禁用一个或多个共享存储目标。
// 此操作完成后需要调用 UpdateS2Servers 以应用到共享存储服务器上。
func (c *S2) DisableS2SharedTargets(params DisableS2SharedTargetsRequest) (DisableS2SharedTargetsResponse, error) {
	var result DisableS2SharedTargetsResponse
	err := c.Get("DisableS2SharedTargets", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

// TODO: 文档错误
// https://docs.qingcloud.com/api/vsan/modify_s2_shared_target.html
type ModifyS2SharedTargetsRequest struct {
	SharedTargetsN  qingcloud.NumberedString
	Operation       qingcloud.String
	InitiatorNamesN qingcloud.NumberedString
}
type ModifyS2SharedTargetsResponse struct {
	qingcloud.CommonResponse
	SharedTargets []string `json:"shared_targets"`
}

// ModifyS2SharedTargets
func (c *S2) ModifyS2SharedTargets(params ModifyS2SharedTargetsRequest) (ModifyS2SharedTargetsResponse, error) {
	var result ModifyS2SharedTargetsResponse
	err := c.Get("ModifyS2SharedTargets", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type AttachToS2SharedTargetRequest struct {
	SharedTarget qingcloud.String
}
type AttachToS2SharedTargetResponse struct {
	qingcloud.CommonResponse
	SharedTargets []string `json:"shared_targets"`
}

// TODO: 文档错误
// https://docs.qingcloud.com/api/vsan/attach_to_s2_shared_target.html
// AttachToS2SharedTarget 为共享存储目标添加硬盘。
// 此操作完成后需要调用 UpdateS2Servers 以应用到共享存储服务器上。
func (c *S2) AttachToS2SharedTarget(params AttachToS2SharedTargetRequest) (AttachToS2SharedTargetResponse, error) {
	var result AttachToS2SharedTargetResponse
	err := c.Get("AttachToS2SharedTarget", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

// TODO: 文档错误
// https://docs.qingcloud.com/api/vsan/detach_from_s2_shared_target.html
type DetachFromS2SharedTargetRequest struct {
	SharedTarget qingcloud.String
	VolumesN     qingcloud.NumberedString
}
type DetachFromS2SharedTargetResponse struct {
	qingcloud.CommonResponse
	SharedTargets []string `json:"shared_targets"`
}

// DetachFromS2SharedTarget 从共享存储目标卸载硬盘。
// 此操作完成后需要调用 UpdateS2Servers 以应用到共享存储服务器上。
func (c *S2) DetachFromS2SharedTarget(params DetachFromS2SharedTargetRequest) (DetachFromS2SharedTargetResponse, error) {
	var result DetachFromS2SharedTargetResponse
	err := c.Get("DetachFromS2SharedTarget", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DescribeS2DefaultParametersRequest struct {
	ServiceType qingcloud.String
	TargetType  qingcloud.String
	Offset      qingcloud.Integer
	Limit       qingcloud.Integer
}
type DescribeS2DefaultParametersResponse struct {
	qingcloud.CommonResponse
	S2DefaultParametersSet []string `json:"s2_default_parameters_set"`
}

// DescribeS2DefaultParameters
func (c *S2) DescribeS2DefaultParameters(params DescribeS2DefaultParametersRequest) (DescribeS2DefaultParametersResponse, error) {
	var result DescribeS2DefaultParametersResponse
	err := c.Get("DescribeS2DefaultParameters", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}
