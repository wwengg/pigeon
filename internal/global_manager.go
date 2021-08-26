// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午11:44
// @Update  Wangwengang  2021/8/19 下午11:44
package internal

import (
	"fmt"
	"os"
	"sync"

	"github.com/golang/protobuf/proto"
)

var (
	defHostname, _ = os.Hostname()
	DefKey         = fmt.Sprintf("pigeon.%s", defHostname)
)

type GlobalManager struct {
	gLock   sync.RWMutex
	Clients sync.Map
	Rooms   sync.Map
}

var GlobalMgrObj *GlobalManager

func init() {
	GlobalMgrObj = &GlobalManager{}
}
func (gm *GlobalManager) AddClient(client *Client) {
	gm.Clients.Store(client.CID, client)
}

func (gm *GlobalManager) RemoveClientByCID(cID uint64) {
	gm.Clients.Delete(cID)
}

func (gm *GlobalManager) GetClientByCID(cID uint64) *Client {
	if c, ok := gm.Clients.Load(cID); ok {
		return c.(*Client)
	}
	return nil
}

// 获取所有用户信息
func (gm *GlobalManager) GetAllClient() []*Client {
	// 创建返回的client集合的切片
	clients := make([]*Client, 0)
	// 遍历所有sync.Map中的键值对
	gm.Clients.Range(func(k, v interface{}) bool {
		clients = append(clients, v.(*Client))
		return true
	})
	return clients
}

func (gm *GlobalManager) Broadcast(op uint32, data proto.Message) {
	// 遍历所有sync.Map中的键值对
	gm.Clients.Range(func(k, v interface{}) bool {
		c := v.(*Client)
		c.SendMsg(op, data)
		return true
	})
}
