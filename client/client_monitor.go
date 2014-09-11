package client

func (c *Client) GetMonitor(params Params) string {
	result := c.Get("GetMonitor", params)
	return result
}
