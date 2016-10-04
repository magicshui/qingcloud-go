package notification

import (
	"time"
)

type NotificationCenterPost struct {
	NotificationCenterPostID string    `json:"notification_center_post_id"`
	CreateTime               time.Time `json:"create_time"`
	Status                   string    `json:"status"`
	PostType                 string    `json:"post_type"`
	Title                    string    `json:"title"`
	Content                  string    `json:"content"`
}
