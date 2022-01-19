package godis

type Client struct {
	token  string
}


func New() *Client {
	c := &Client{}
	return c
}

func Login(token string) {
	//logic
}