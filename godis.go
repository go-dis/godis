package godis

type Client struct {
	token string
}

func New() *Client {
	c := &Client{}
	return c
}

func (c *Client) Login(token string) {
	c.token = token
        
}
