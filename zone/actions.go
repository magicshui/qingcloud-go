package zone

import (
	"github.com/magicshui/qingcloud-go"
)

type ZONE struct {
	*qingcloud.Client
}

func NewClient(clt *qingcloud.Client) *ZONE {
	var a = &ZONE{
		Client: clt,
	}
	return a
}

// DescribeZonesRequest 请求
type DescribeZonesRequest struct {
	ZonesN qingcloud.NumberedString
	StausN qingcloud.NumberedString
}

// DescribeZonesResponse 结果
type DescribeZonesResponse struct {
	ZoneSet    []Zone `json:"zone_set"`
	TotalCount int    `json:"total_count"`
	qingcloud.CommonResponse
}

// DescribeZones 获取可访问的区域列表。
func (c *ZONE) DescribeZones(params DescribeZonesRequest) (DescribeZonesResponse, error) {
	var result DescribeZonesResponse
	err := c.Get("DescribeZones", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}
