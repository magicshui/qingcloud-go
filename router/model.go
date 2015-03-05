package router

import (
	"time"
)

type Router struct {
	RouterId    string `json:"router_id"`
	Status      string `json:"status"`
	IsApplied   int    `json:"is_applied"`
	Description string `json:"description"`
	Eip         struct {
		EipName string `json:"eip_name"`
		EipId   string `json:"eip_id"`
		EipAddr string `json:"eip_addr"`
	} `json:"eip"`
	SubCode          int       `json:"sub_code"`
	TransitionStatus string    `json:"transition_status"`
	SecurityGroupId  string    `json:"security_group_id"`
	CreateTime       time.Time `json:"create_time"`
	PrivateIp        string    `json:"private_ip"`
	RouterType       int       `json:"router_type"`
	Vxnets           []struct {
		NicId   string `json:"nic_id"`
		VxnetId string `json:"vxnet_id"`
	} `json:"vxnets"`
	RouterName string `json:"router_name"`
}
