package keypair

import (
	"github.com/magicshui/qingcloud-go"
)

type DescribeKeyPairsRequest struct {
}
type DescribeKeyPairsResponse qingcloud.CommonResponse

func DescribeKeyPairs(c *qingcloud.Client, params DescribeKeyPairsRequest) (DescribeKeyPairsResponse, error) {
	var result DescribeKeyPairsResponse
	err := c.Get("DescribeKeyPairs", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}
