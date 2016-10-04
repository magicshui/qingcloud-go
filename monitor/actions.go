package monitor

import (
	"github.com/magicshui/qingcloud-go"
)

type MONITOR struct {
	*qingcloud.Client
}

func NewClient(clt *qingcloud.Client) *MONITOR {
	var a = &MONITOR{
		Client: clt,
	}
	return a
}

type GetMonitorRequest struct {
	Resource  qingcloud.String
	MetersN   qingcloud.NumberedString
	Step      qingcloud.String
	StartTime qingcloud.String
	EndTime   qingcloud.String
}
type GetMonitorResponse struct {
	MeterSet   []Meter `json:"meter_set"`
	ResourceId string  `json:"resource_id"`
	qingcloud.CommonResponse
}

// GetMonitor
// 获取资源监控数据。支持的资源包括主机、公网 IP 和路由器， 选定不同类型资源，可获取的监控项不同。
// 主机的监控项包括: CPU 使用率，内存使用率，系统盘数据(吞吐量， IOPS 和使用率)， 主机连接私有网络的网卡流量，与主机绑定的各磁盘数据（吞吐量， IOPS 和使用率）。
// 注解 其中内存使用率和磁盘使用率暂不支持 Windows ，以及 kernel 版本过低的 Linux
// 公网 IP 资源可得到公网 “进/出” 的流量数据。
// 如果资源为路由器，可得到路由器在基础网络的流量数据，以及与路由器连接的私有网络的流量数据。
// 为减少数据传输，在保持数据结构清晰的前提下，我们对监控数据做了压缩， 在解析返回数据时要留意。详细说明参见 监控数据压缩说明 。
func (c *MONITOR) GetMonitor(params GetMonitorRequest) (GetMonitorResponse, error) {
	var result GetMonitorResponse
	err := c.Get("GetMonitor", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type GetLoadBalancerMonitorRequest struct {
	Resource  qingcloud.String
	MetersN   qingcloud.NumberedString
	Step      qingcloud.String
	StartTime qingcloud.String
	EndTime   qingcloud.String
}
type GetLoadBalancerMonitorResponse struct {
	MeterSet   []Meter `json:"meter_set"`
	ResourceId string  `json:"resource_id"`
	qingcloud.CommonResponse
}

func (c *MONITOR) GetLoadBalancerMonitor(params GetLoadBalancerMonitorRequest) (GetLoadBalancerMonitorResponse, error) {
	var result GetLoadBalancerMonitorResponse
	err := c.Get("GetLoadBalancerMonitor", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type GetRDBMonitorRequest struct {
	Resource  qingcloud.String
	MetersN   qingcloud.NumberedString
	Step      qingcloud.String
	StartTime qingcloud.String
	EndTime   qingcloud.String
}
type GetRDBMonitorResponse struct {
	MeterSet   []Meter `json:"meter_set"`
	ResourceId string  `json:"resource_id"`
	qingcloud.CommonResponse
}

// GetRDBMonitor
// 获取指定数据库实例的监控信息。
func (c *MONITOR) GetRDBMonitor(params GetRDBMonitorRequest) (GetRDBMonitorResponse, error) {
	var result GetRDBMonitorResponse
	err := c.Get("GetRDBMonitor", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type GetCacheMonitorRequest struct {
	Resource  qingcloud.String
	MetersN   qingcloud.NumberedString
	Step      qingcloud.String
	StartTime qingcloud.String
	EndTime   qingcloud.String
}
type GetCacheMonitorResponse struct {
	qingcloud.CommonResponse
	MeterSet   []Data `json:"meter_set"`
	ResourceId string `json:"resource_id"`
}

// GetCacheMonitor
// 通过此 API 可获得缓存服务每个节点的流量监控，以及与缓存相关的监控项目。
// 因为缓存服务包含多个缓存节点，每个缓存节点都有独立的监控数据。
// 不同的缓存类型，监控数据含义不同，具体可见下面”监控数据集说明”
// 为减少数据传输，在保持数据结构清晰的前提下，我们对监控数据做了压缩， 在解析返回数据时要留意。详细说明参见 监控数据压缩说明 。
func (c *MONITOR) GetCacheMonitor(params GetCacheMonitorRequest) (GetCacheMonitorResponse, error) {
	var result GetCacheMonitorResponse
	err := c.Get("GetCacheMonitor", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type GetZooKeeperMonitorRequest struct {
	Resource  qingcloud.String
	MetersN   qingcloud.NumberedString
	Step      qingcloud.String
	StartTime qingcloud.String
	EndTime   qingcloud.String
}
type GetZooKeeperMonitorResponse struct {
	qingcloud.CommonResponse
	MeterSet   []Meter `json:"meter_set"`
	ResourceId string  `json:"resource_id"`
}

// GetZooKeeperMonitor 通过此 API 可获得 ZooKeeper 服务每个节点的监控指标。
// 因为 ZooKeeper 服务包含多个节点，每个 ZooKeeper 节点都有独立的监控数据。
// 为减少数据传输，在保持数据结构清晰的前提下，我们对监控数据做了压缩， 在解析返回数据时要留意。详细说明参见 监控数据压缩说明 。
func (c *MONITOR) GetZooKeeperMonitor(params GetZooKeeperMonitorRequest) (GetZooKeeperMonitorResponse, error) {
	var result GetZooKeeperMonitorResponse
	err := c.Get("GetZooKeeperMonitor", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type GetQueueMonitorRequest struct {
	Resource  qingcloud.String
	MetersN   qingcloud.NumberedString
	Step      qingcloud.String
	StartTime qingcloud.String
	EndTime   qingcloud.String
}
type GetQueueMonitorResponse struct {
	qingcloud.CommonResponse
	MeterSet   []Meter `json:"meter_set"`
	ResourceId string  `json:"resource_id"`
}

// GetQueueMonitor 通过此 API 可获得消息队列服务每个节点的流量监控，以及与消息队列相关的监控项目。
// 因为消息队列服务包含多个节点，每个消息队列节点都有独立的监控数据。
// 为减少数据传输，在保持数据结构清晰的前提下，我们对监控数据做了压缩， 在解析返回数据时要留意。详细说明参见 监控数据压缩说明 。
func (c *MONITOR) GetQueueMonitor(params GetQueueMonitorRequest) (GetQueueMonitorResponse, error) {
	var result GetQueueMonitorResponse
	err := c.Get("GetQueueMonitor", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}
