// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午11:46
// @Update  Wangwengang  2021/8/19 下午11:46
package internal

import (
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/wwengg/arsenal/anet"
)

type Client struct {
	CID  uint64
	Conn anet.Connection

	cLock sync.RWMutex
}

// new Client
func NewClient(conn anet.Connection, cID uint64) *Client {
	return &Client{
		CID:  cID,
		Conn: conn,
	}
}

// get Conn
func (c *Client) GetConn() anet.Connection {
	c.cLock.Lock()
	defer c.cLock.Unlock()

	return c.Conn
}

func (c *Client) SendMsg(msgID uint32, data proto.Message) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("出了错：", r)
		}
	}()
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err: ", err)
		return
	}
	if c == nil {
		fmt.Println("c is nil")
		return
	}
	//fmt.Printf("after Marshal data = %+v\n", msg)
	if c.Conn == nil {
		fmt.Println("connection in client is nil")
		return
	}

	if err := c.Conn.SendMsg(msgID, msg); err != nil {
		fmt.Println("Clinet SendMsg error !")
		fmt.Println(err)
		return
	}
	return
}


func (c *Client) ConnectionLost(){
	// 删除redis中的在线

}

func (c *Client) ConnectionAdd(){

}

