package job

import (
	"github.com/magicshui/qingcloud-go"
)

func (c *Client) DescribeJobs(params qingcloud.Params) ([]byte, error) {
	return c.Get("DescribeJobs", params)
}
