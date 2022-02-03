package godis

import (
	"time"
	"github.com/gorilla/websocket"
)


type ws struct {
	socket    *websocket.Conn
	gateway   string
	heartbeat *time.Time
	listen    chan interface{}
	client    *Client
	sequence  int64
}

func newsocket(c *Client) *ws {
	return &ws{
		client: c,
		gateway: "wss://gateway.discord.gg",
	}
}


func (w *ws) Connect() {

}

func (w *ws) identifyData() {
	w.socket.WriteJSON(&identify {
		operator: 1,
		Data: &scopes{
			Token: w.client.token,
			Intents: 8,
			Properties: dataProperties{
                            "linux", "client/ 0.1" , "client/ 0.1",
			},
		},
	})
}