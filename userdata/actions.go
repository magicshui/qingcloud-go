package userdata

import (
	"github.com/magicshui/qingcloud-go"
)

type DescribeDNSAliasesRequest struct {
}
type DescribeDNSAliasesResponse qingcloud.CommonResponse

func DescribeDNSAliases(c *qingcloud.Client, params DescribeDNSAliasesRequest) (DescribeDNSAliasesResponse, error) {
	var result DescribeDNSAliasesResponse
	err := c.Get("DescribeDNSAliases", qingcloud.TransfomerRequestToParams(&params), &result)
	return result, err
}
