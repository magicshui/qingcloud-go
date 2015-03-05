package qingcloud

type QCError struct {
	RetCode int    `json:"ret_code"`
	Message string `json:"message"`
}
