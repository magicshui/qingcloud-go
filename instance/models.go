package instance

import (
	"time"
)

type Instance struct {
	VcpusCurrent int      `json:"vcpus_current"`
	InstanceID   string   `json:"instance_id"`
	VolumeIds    []string `json:"volume_ids"`
	Vxnets       []struct {
		VxnetName string `json:"vxnet_name"`
		VxnetType int    `json:"vxnet_type"`
		VxnetID   string `json:"vxnet_id"`
		NicID     string `json:"nic_id"`
		PrivateIP string `json:"private_ip"`
	} `json:"vxnets"`
	Eip struct {
		EipID     string `json:"eip_id"`
		EipAddr   string `json:"eip_addr"`
		Bandwidth string `json:"bandwidth"`
	} `json:"eip"`
	MemoryCurrent    int         `json:"memory_current"`
	SubCode          int         `json:"sub_code"`
	TransitionStatus string      `json:"transition_status"`
	InstanceName     string      `json:"instance_name"`
	InstanceType     string      `json:"instance_type"`
	CreateTime       time.Time   `json:"create_time"`
	Status           string      `json:"status"`
	Description      interface{} `json:"description"`
	SecurityGroup    struct {
		IsDefault       int    `json:"is_default"`
		SecurityGroupID string `json:"security_group_id"`
	} `json:"security_group"`
	StatusTime time.Time `json:"status_time"`
	Image      struct {
		ProcessorType string `json:"processor_type"`
		Platform      string `json:"platform"`
		ImageSize     int    `json:"image_size"`
		ImageName     string `json:"image_name"`
		ImageID       string `json:"image_id"`
		OsFamily      string `json:"os_family"`
		Provider      string `json:"provider"`
	} `json:"image"`
	KeypairIds []string `json:"keypair_ids"`
}
