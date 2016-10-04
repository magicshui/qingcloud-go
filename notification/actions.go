package notification

import (
	"github.com/magicshui/qingcloud-go"
)

type NOTIFICATION struct {
	*qingcloud.Client
}

func NewClient(clt *qingcloud.Client) *NOTIFICATION {
	var a = &NOTIFICATION{
		Client: clt,
	}
	return a
}

type DescribeNotificationCenterUserPostsRequest struct {
	PostTypeN qingcloud.NumberedString
	StatusN   qingcloud.NumberedString
	Offset    qingcloud.Integer
	Limit     qingcloud.Integer
}
type DescribeNotificationCenterUserPostsResponse struct {
	qingcloud.CommonResponse
	TotalCount                int                      `json:"total_count"`
	NotificationCenterPostSet []NotificationCenterPost `json:"notification_center_post_set"`
}

// DescribeNotificationCenterUserPosts
func (c *NOTIFICATION) DescribeNotificationCenterUserPosts(params DescribeNotificationCenterUserPostsRequest) (DescribeNotificationCenterUserPostsResponse, error) {
	var result DescribeNotificationCenterUserPostsResponse
	err := c.Get("DescribeNotificationCenterUserPosts", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}
