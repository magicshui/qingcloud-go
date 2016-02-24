package qingcloud

import (
	"fmt"
)

type CommonResponse struct {
	Action  string `json:"action"`
	JobId   string `json:"job_id"`
	RetCode int    `json:"ret_code"`
	Message string `json:"message"`
}

func (c *CommonResponse) Error() string {
	return fmt.Sprintf("rect_code: %d , message: %s", c.RetCode, c.Message)
}
