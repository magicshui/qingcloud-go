package resource

import (
	"github.com/magicshui/qingcloud-go"
)

type RESOURCE struct {
	*qingcloud.Client
}

func NewClient(clt *qingcloud.Client) *RESOURCE {
	var a = &RESOURCE{
		Client: clt,
	}
	return a
}

type DescribeSharedResourceGroupsRequest struct {
	ResourceGroupsN qingcloud.NumberedString
}
type DescribeSharedResourceGroupsResponse struct {
	qingcloud.CommonResponse
	TotalCount             int                   `json:"total_count"`
	SharedResourceGroupSet []SharedResourceGroup `json:"shared_resource_group_set"`
}

// DescribeSharedResourceGroups 查询共享给自己的资源组。
func (c *RESOURCE) DescribeSharedResourceGroups(params DescribeSharedResourceGroupsRequest) (DescribeSharedResourceGroupsResponse, error) {
	var result DescribeSharedResourceGroupsResponse
	err := c.Get("DescribeSharedResourceGroups", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DescribeResourceGroupsRequest struct {
	ResourceGroupsN qingcloud.NumberedString
	SearchWord      qingcloud.String
	Limit           qingcloud.Integer
	Offset          qingcloud.Integer
	Verbose         qingcloud.Integer
	SortKey         qingcloud.String
	Reverse         qingcloud.Integer
}
type DescribeResourceGroupsResponse struct {
	qingcloud.CommonResponse
	TotalCount       int             `json:"total_count"`
	ResourceGroupSet []ResourceGroup `json:"resource_group_set"`
}

// DescribeResourceGroups 查询资源组信息。
func (c *RESOURCE) DescribeResourceGroups(params DescribeResourceGroupsRequest) (DescribeResourceGroupsResponse, error) {
	var result DescribeResourceGroupsResponse
	err := c.Get("DescribeResourceGroups", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type CreateResourceGroupsRequest struct {
	ResourceGroupName qingcloud.String
	Description       qingcloud.String
	Count             qingcloud.Integer
}
type CreateResourceGroupsResponse struct {
	qingcloud.CommonResponse
	ResourceGroupIDs []string `json:"resource_group_ids"`
}

// CreateResourceGroups 创建资源组。
func (c *RESOURCE) CreateResourceGroups(params CreateResourceGroupsRequest) (CreateResourceGroupsResponse, error) {
	var result CreateResourceGroupsResponse
	err := c.Get("CreateResourceGroups", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ModifyResourceGroupAttributesRequest struct {
	ResourceGroupName qingcloud.String
	ResourceGroup     qingcloud.String
}
type ModifyResourceGroupAttributesResponse struct {
	qingcloud.CommonResponse
	ResourceGroupID string `json:"resource_group_id"`
}

// ModifyResourceGroupAttributes 修改资源组属性。
func (c *RESOURCE) ModifyResourceGroupAttributes(params ModifyResourceGroupAttributesRequest) (ModifyResourceGroupAttributesResponse, error) {
	var result ModifyResourceGroupAttributesResponse
	err := c.Get("ModifyResourceGroupAttributes", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

// TODO: 文档不全 https://docs.qingcloud.com/api/resource_acl/DeleteResourceGroups.html
type DeleteResourceGroupsRequest struct {
	ResourceGroupsN qingcloud.NumberedString
}
type DeleteResourceGroupsResponse struct {
	qingcloud.CommonResponse
	ResourceGroups []string `json:"resource_groups"`
}

// DeleteResourceGroups 删除资源组。
func (c *RESOURCE) DeleteResourceGroups(params DeleteResourceGroupsRequest) (DeleteResourceGroupsResponse, error) {
	var result DeleteResourceGroupsResponse
	err := c.Get("DeleteResourceGroups", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DescribeResourceGroupItemsRequest struct {
	ResourceGroupsN qingcloud.NumberedString
	ResourcesN      qingcloud.NumberedString
	SearchWord      qingcloud.String
	Limit           qingcloud.Integer
	Offset          qingcloud.Integer
	Verbose         qingcloud.Integer
	SortKey         qingcloud.String
	Reverse         qingcloud.Integer
}
type DescribeResourceGroupItemsResponse struct {
	qingcloud.CommonResponse
	// TODO: 这个字段待定
	ItemSet              []interface{}       `json:"item_set"`
	ResourceGroupItemSet []ResourceGroupItem `json:"resource_group_item_set"`
	TotalCount           int                 `json:"total_count"`
}

// DescribeResourceGroupItems
func (c *RESOURCE) DescribeResourceGroupItems(params DescribeResourceGroupItemsRequest) (DescribeResourceGroupItemsResponse, error) {
	var result DescribeResourceGroupItemsResponse
	err := c.Get("DescribeResourceGroupItems", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type AddResourceGroupItemsRequest struct {
	ResourceGroup qingcloud.String
	ResourcesN    qingcloud.NumberedString
}
type AddResourceGroupItemsResponse struct {
	qingcloud.CommonResponse
	ResourceIDs []string `json:"resource_ids"`
}

// AddResourceGroupItems 将资源加入到资源组中。
func (c *RESOURCE) AddResourceGroupItems(params AddResourceGroupItemsRequest) (AddResourceGroupItemsResponse, error) {
	var result AddResourceGroupItemsResponse
	err := c.Get("AddResourceGroupItems", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DeleteResourceGroupItemsRequest struct {
	ResourceGroup qingcloud.String
	ResourcesN    qingcloud.NumberedString
}
type DeleteResourceGroupItemsResponse struct {
	qingcloud.CommonResponse
	ResourceGroupID string `json:"resource_group_id"`
}

// DeleteResourceGroupItems 将资源从资源组里删除。
func (c *RESOURCE) DeleteResourceGroupItems(params DeleteResourceGroupItemsRequest) (DeleteResourceGroupItemsResponse, error) {
	var result DeleteResourceGroupItemsResponse
	err := c.Get("DeleteResourceGroupItems", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}
