package qingcloud

func (c *Client) GetMonitor(params Params) ([]byte, error) {
	return c.Get("GetMonitor", params)
}
