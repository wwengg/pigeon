// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午11:44
// @Update  Wangwengang  2021/8/19 下午11:44
package internal

import (
	"sync"
)

type GlobalManager struct {
	gLock sync.RWMutex
	Clients sync.Map
	Rooms sync.Map

}

var GlobalMgrObj *GlobalManager

func init(){
	GlobalMgrObj = nil
}
func (gm *GlobalManager) AddClient(client *Client){
	gm.Clients.Store(client.CID,client)
}

func (gm *GlobalManager) RemoveClientByCID(cID uint64){
	gm.Clients.Delete(cID)
}

func (gm *GlobalManager) GetClientByCID(cID uint64) *Client{
	if c,ok := gm.Clients.Load(cID);ok{
		return c.(*Client)
	}
	return nil
}

