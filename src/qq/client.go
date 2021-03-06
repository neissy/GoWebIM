package qq

import (
	"code.google.com/p/go.net/websocket"

//	"html"
//	"fmt"
)

type UserInfo struct {
	Id     int
	Name   string
	Avatar string
}

type Client struct {
	Id       int
	Addr     string
	Info     UserInfo
	Conn     *websocket.Conn
	CurGroup *Group
}

func (c *Client) Write(content string) {
	websocket.Message.Send(c.Conn, content)
}

func (c *Client) SetName(name, avatar string) {
	c.Info.Name = name
	c.Info.Avatar = avatar
	c.Info.Id = c.Id
	c.On_SetName()
}
