// @Title
// @Description
// @Author  Wangwengang  2021/8/20 上午10:10
// @Update  Wangwengang  2021/8/20 上午10:10
package router

import (
	"time"

	"github.com/wwengg/arsenal/anet"
	"github.com/wwengg/arsenal/atimer"
	"github.com/wwengg/arsenal/logger"
	"go.uber.org/zap"

	"github.com/wwengg/pigeon/internal/timer"
)

type KeepAlive struct {
	anet.Router
}

func (r *KeepAlive) Handle(request anet.Request) {
	cID,err := request.GetConnection().GetProperty("cID")
	if err != nil {
		logger.ZapLog.Debug("Keepalive cID error",zap.Any("err",err))
		return
	}
	//	获取当前连接的用户id
	uID,err := request.GetConnection().GetProperty("uID")
	if err != nil {
		logger.ZapLog.Debug("Keepalive GetProperty uID error",zap.Any("err",err),zap.Any("cID",cID))
		//request.GetConnection().Stop()
		return
	}

	tID,err := request.GetConnection().GetProperty("tID")
	if err != nil {
		logger.ZapLog.Debug("Keepalive GetProperty tId error",zap.Any("err",err),zap.Any("uId",uID))
		return
	}
	timer.AutoTimerSchedulerObj.CancelTimer(tID.(uint32))

	tId,_ := timer.AutoTimerSchedulerObj.CreateTimerAfter(atimer.NewDelayFunc(timer.StopConnect, []interface{}{request.GetConnection()}),30*time.Second)
	request.GetConnection().RemoveProperty("tID")
	request.GetConnection().SetProperty("tID",tId)

}
