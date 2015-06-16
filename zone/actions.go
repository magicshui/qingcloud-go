package zone

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeZones(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeZones", params)
}
