package instance

import (
	"github.com/magicshui/qingcloud-go"
)

type INSTANCE struct {
	*qingcloud.Client
}

func NewClient(clt *qingcloud.Client) *INSTANCE {
	return &INSTANCE{
		Client: clt,
	}
}

type DescribeInstanceRequest struct {
	InstancesN    qingcloud.NumberedString
	ImageIdN      qingcloud.NumberedString
	InstanceTypeN qingcloud.NumberedString
	InstanceClass qingcloud.Integer

	StatusN    qingcloud.String
	SearchWord qingcloud.String
	TagsN      qingcloud.NumberedString
	Verbose    qingcloud.Integer
	Offset     qingcloud.Integer
	Limit      qingcloud.Integer
}

type DescribeInstanceResponse struct {
	InstanceSet []Instance `json:"instance_set"`
	TotalCount  int        `json:"total_count"`
	qingcloud.CommonResponse
}

// DescribeInstances 获取一个或多个主机
// 可根据主机ID, 状态, 主机名称, 映像ID 作过滤条件, 来获取主机列表。 如果不指定任何过滤条件, 默认返回你所拥有的所有主机。
func DescribeInstances(c *qingcloud.Client, params DescribeInstanceRequest) (DescribeInstanceResponse, error) {
	var result DescribeInstanceResponse
	// 主机性能类型: 性能型:0 ,超高性能型:1
	params.InstanceClass.Enum(0, 1)
	// 主机状态: pending, running, stopped, suspended, terminated, ceased
	params.InstanceTypeN.Enum("pending", "running", "stopped", "suspended", "terminated", "ceased")
	err := c.Get("DescribeInstances", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type RunInstancesRequest struct {
	ImageId       qingcloud.String
	InstanceType  qingcloud.String
	Cpu           qingcloud.Integer
	Memory        qingcloud.Integer
	Count         qingcloud.Integer
	InstanceName  qingcloud.String
	LoginMode     qingcloud.String
	LoginKeypair  qingcloud.String
	LoginPasswd   qingcloud.String
	VxnetsN       qingcloud.NumberedString
	SecurityGroup qingcloud.String
	VolumesN      qingcloud.NumberedString
	NeedNewsid    qingcloud.Integer
	NeedUserdata  qingcloud.Integer
	UserdataType  qingcloud.String
	UserdataValue qingcloud.String
	InstanceClass qingcloud.String
	UserdataPath  qingcloud.String
	UserdataFile  qingcloud.String
}

type RunInstancesResponse struct {
	Instances []string `json:"instances"`
	qingcloud.CommonResponse
}

// RunInstances 创建指定配置，指定数量的主机。
// 当你创建主机时，主机会先进入 pending 状态，直到创建完成后，变为 running 状态。 你可以使用 DescribeInstances 检查主机状态。
// 创建主机时，一旦参数 vxnets.n 包含基础网络（即： vxnet-0 ），则需要指定防火墙 security_group，如果没有指定，青云会自动使用缺省防火墙。
// 青云给主机定义了几种经典配置，可通过参数 instance_type 指定，配置列表请参考 Instance Types 。 如果经典配置不能满足你的需求，可通过参数 cpu, memory 自定义主机配置。
// 如果参数中既指定 instance_type ，又指定了 cpu 和 memory ， 则以指定的 cpu 和 memory 为准。
func RunInstances(c *qingcloud.Client, params RunInstancesRequest) (RunInstancesResponse, error) {
	var result RunInstancesResponse
	err := c.Get("RunInstances", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type TerminateInstancesRequest struct {
	InstanceN qingcloud.NumberedString
}
type TerminateInstancesResponse qingcloud.CommonResponse

// TerminateInstances 销毁一台或多台主机。
// 销毁主机的前提，是此主机已建立租用信息（租用信息是在创建主机成功后， 几秒钟内系统自动建立的）。所以正在创建的主机（状态为 pending ）， 以及刚刚创建成功但还没有建立租用信息的主机，是不能被销毁的。
func TerminateInstances(c *qingcloud.Client, params TerminateInstancesRequest) (TerminateInstancesResponse, error) {
	var result TerminateInstancesResponse
	err := c.Get("TerminateInstances", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type StartInstancesRequest struct {
	InstancesN qingcloud.NumberedString
}
type StartInstancesResponse qingcloud.CommonResponse

// StartInstances 启动一台或多台关闭状态的主机。
// 主机只有在关闭 stopped 状态才能被启动，如果处于非关闭状态，则返回错误信息。
func StartInstances(c *qingcloud.Client, params StartInstancesRequest) (StartInstancesResponse, error) {
	var result StartInstancesResponse
	err := c.Get("StartInstances", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type StopInstancesRequest struct {
	InstancesN qingcloud.NumberedString
	Force      qingcloud.Integer
}
type StopInstancesResponse qingcloud.CommonResponse

// StopInstances 关闭一台或多台运行状态的主机。
// 主机只有在运行 running 状态才能被关闭，如果处于非运行状态，则返回错误信息。
func StopInstances(c *qingcloud.Client, params StopInstancesRequest) (StopInstancesResponse, error) {
	var result StopInstancesResponse
	err := c.Get("StopInstances", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type RestartInstancesRequest struct {
	InstancesN qingcloud.NumberedString
}
type RestartInstancesResponse qingcloud.CommonResponse

// RestartInstances 重启一台或多台运行状态的主机。
// 主机只有在运行 running 状态才能被重启，如果处于非运行状态，则返回错误信息。
func RestartInstances(c *qingcloud.Client, params RestartInstancesRequest) (RestartInstancesResponse, error) {
	var result RestartInstancesResponse
	err := c.Get("RestartInstances", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ResetInstancesRequest struct {
	InstancesN   qingcloud.NumberedString
	LoginMode    qingcloud.String
	LoginKeypair qingcloud.String
	LoginPasswd  qingcloud.String
	NeedNewsid   qingcloud.String
}
type ResetInstancesResponse qingcloud.CommonResponse

// ResetInstances 将一台或多台主机的系统盘重置到初始状态。 被重置的主机必须处于运行（ running ）或关闭（ stopped ）状态。
// 重置只涉及系统盘数据，不包含主机所加载的硬盘。
func ResetInstances(c *qingcloud.Client, params ResetInstancesRequest) (ResetInstancesResponse, error) {
	var result ResetInstancesResponse
	err := c.Get("ResetInstances", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ResizeInstancesRequest struct {
	InstancesN   qingcloud.NumberedString
	InstanceType qingcloud.String
	Cpu          qingcloud.String
	Memory       qingcloud.String
	Zone         qingcloud.String
}

type ResizeInstancesResponse qingcloud.CommonResponse

// ResizeInstances 修改主机配置，包括 CPU 和内存。主机状态必须是关闭的 stopped ，不然会返回错误。
// 如果使用预设的 instance_type ，参数中就不需再指定 CPU 或内存，配置列表请参考 Instance Types 。
// 如果参数中没有指定 instance_type ，则必须指定 cpu 和 memory。
// 如果参数中既指定 instance_type ，又指定了 cpu 和 memory ， 则以指定的 cpu 和 memory 为准。
func ResizeInstances(c *qingcloud.Client, params ResizeInstancesRequest) (ResizeInstancesResponse, error) {
	var result ResizeInstancesResponse
	err := c.Get("ResizeInstances", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ModifyInstanceAttributesRequest struct {
	Instance     qingcloud.String
	InstanceName qingcloud.String
	Description  qingcloud.String
}

type ModifyInstanceAttributesResponse struct {
	Action string `json:"action"`
	qingcloud.CommonResponse
}

// ModifyInstanceAttributes 修改一台主机的名称和描述。
// 修改时不受主机状态限制。一次只能修改一台主机。
func ModifyInstanceAttributes(c *qingcloud.Client, params ModifyInstanceAttributesRequest) (ModifyInstanceAttributesResponse, error) {
	var result ModifyInstanceAttributesResponse
	err := c.Get("ModifyInstanceAttributes", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}
