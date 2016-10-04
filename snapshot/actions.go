package snapshot

import (
	"github.com/magicshui/qingcloud-go"
)

type SNAPSHOT struct {
	*qingcloud.Client
}

func NewClient(clt *qingcloud.Client) *SNAPSHOT {
	var a = &SNAPSHOT{
		Client: clt,
	}
	return a
}

type DescribeSnapshotsRequest struct {
	SnapshotsN   qingcloud.NumberedString
	ResourceId   qingcloud.String
	SnapshotType qingcloud.Integer
	StatusN      qingcloud.NumberedString

	SearchWord qingcloud.String
	TagsN      qingcloud.NumberedString
	Verbose    qingcloud.Integer
	Offset     qingcloud.Integer
	Limit      qingcloud.Integer
}
type DescribeSnapshotsResponse struct {
	SnapshotSet []Snapshot `json:"snapshot_set"`
	qingcloud.CommonResponse
}

// DescribeSnapshots
// 获取指定资源的所有备份。
func (c *SNAPSHOT) DescribeSnapshots(params DescribeSnapshotsRequest) (DescribeSnapshotsResponse, error) {
	var result DescribeSnapshotsResponse
	err := c.Get("DescribeSnapshots", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type CreateSnapshotsRequest struct {
	ResourcesN   qingcloud.NumberedString
	SnapshotName qingcloud.String
	IsFull       qingcloud.Integer
}
type CreateSnapshotsResponse struct {
	Snapshots []string `json:"snapshots"`
	qingcloud.CommonResponse
}

// CreateSnapshots
// 为指定的资源创建备份。
func (c *SNAPSHOT) CreateSnapshots(params CreateSnapshotsRequest) (CreateSnapshotsResponse, error) {
	var result CreateSnapshotsResponse
	err := c.Get("CreateSnapshots", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type DeleteSnapshotsRequest struct {
	SnapshotsN qingcloud.NumberedString
}
type DeleteSnapshotsResponse qingcloud.CommonResponse

// DeleteSnapshots
// 删除备份。请注意，当删除一条备份链中某个增量备份点之后，该增量备份点后的所有备份点都会被自动删除， 并且该操作是不可逆的。
func (c *SNAPSHOT) DeleteSnapshots(params DeleteSnapshotsRequest) (DeleteSnapshotsResponse, error) {
	var result DeleteSnapshotsResponse
	err := c.Get("DeleteSnapshots", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ApplySnapshotsRequest struct {
	SnapshotsN qingcloud.NumberedString
}
type ApplySnapshotsResponse qingcloud.CommonResponse

// ApplySnapshots
// 回滚到指定备份点。请注意，为了保证回滚的安全性，当回滚的资源为运行的主机，或者为绑定在运行主机上的硬盘时，该操作会导致主机重启。
func (c *SNAPSHOT) ApplySnapshots(params ApplySnapshotsRequest) (ApplySnapshotsResponse, error) {
	var result ApplySnapshotsResponse
	err := c.Get("ApplySnapshots", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type ModifySnapshotAttributesRequest struct {
	Snapshot     qingcloud.String
	SnapshotName qingcloud.String
	Description  qingcloud.String
}
type ModifySnapshotAttributesResponse qingcloud.CommonResponse

// ModifySnapshotAttributes
// 修改指定备份的相关属性。
func (c *SNAPSHOT) ModifySnapshotAttributes(params ModifySnapshotAttributesRequest) (ModifySnapshotAttributesResponse, error) {
	var result ModifySnapshotAttributesResponse
	err := c.Get("ModifySnapshotAttributes", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type CaptureInstanceFromSnapshotRequest struct {
	Snapshot  qingcloud.String
	ImageName qingcloud.String
}
type CaptureInstanceFromSnapshotResponse struct {
	IamgeId string `json:"image_id"`
	qingcloud.CommonResponse
}

// CaptureInstanceFromSnapshot
// 将指定备份导出为映像。请注意，此备份点必须为主机的备份点才能导出为映像。
func (c *SNAPSHOT) CaptureInstanceFromSnapshot(params CaptureInstanceFromSnapshotRequest) (CaptureInstanceFromSnapshotResponse, error) {
	var result CaptureInstanceFromSnapshotResponse
	err := c.Get("CaptureInstanceFromSnapshot", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}

type CreateVolumeFromSnapshotRequest struct {
	Snapshot   qingcloud.String
	VolumeName qingcloud.String
}
type CreateVolumeFromSnapshotResponse struct {
	qingcloud.CommonResponse
	VolumeId string `json:"volume_id"`
}

// CreateVolumeFromSnapshot
// 将指定备份导出为硬盘。请注意，此备份点必须为硬盘的备份点才能导出为硬盘，而且通过备份创建的硬盘类型和备份的来源硬盘类型是一致的。
func (c *SNAPSHOT) CreateVolumeFromSnapshot(params CreateVolumeFromSnapshotRequest) (CreateVolumeFromSnapshotResponse, error) {
	var result CreateVolumeFromSnapshotResponse
	err := c.Get("CreateVolumeFromSnapshot", qingcloud.TransfomRequestToParams(&params), &result)
	return result, err
}
