package resource

import (
	"time"
)

type SharedResourceGroup struct {
	ResourceGroupID   string `json:"resource_group_id"`
	CreateTime        string `json:"create_time"`
	Description       string `json:"description"`
	ResourceGroupName string `json:"resource_group_name"`
}
type ResourceGroup struct {
	ResourceGroupID   string    `json:"resource_group_id"`
	CreateTime        time.Time `json:"create_time"`
	Description       string    `json:"description"`
	ResourceGroupName string    `json:"resource_group_name"`
}

type ResourceGroupItem struct {
	ResourceGroupID string `json:"resource_group_id"`
	CreateTime      string `json:"create_time"`
	ResourceType    string `json:"resource_type"`
	ResourceID      string `json:"resource_id"`
}
