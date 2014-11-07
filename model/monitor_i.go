package model

type ModelMonitorI struct {
	Action     string                      `json:"action"`
	RetCode    int                         `json:"ret_code"`
	ResourceId string                      `json:"resource_id"`
	MeterSet   []ModelMonitorIMeterSetItem `json:"meter_set"`
}

type ModelMonitorIMeterSetItem struct {
	Data     []interface{} `json:"data,omitempty"`
	MeterId  string        `json:"meter_id"`
	Sequence string        `json:"sequence,omitempty"`
	VxnetId  string        `json:"vxnet_id,omitempty"`
}

// [128,1]
type ModelMonitorIMeterSetItemSingle struct {
}
type ModelMonitorIMeterSetItemDouble struct {
}

/*
func (m *ModelMonitorIMeterSetItem) Unzip(step float64) {

	for i, v := range m.Data {
		if i == 0 {

		} else {
			if fmt.Sprintf("%T", v) == "float64" {
				var first []float64
				first = m.Data[0][0]
				var result []float64
				result[0] = first[0] + step*i.(float64)
				result[1] = v
				m.SingleData[i] = result
				fmt.Printf(m.Data[i])
			}
		}
	}
}
*/
