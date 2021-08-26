// @Title
// @Description
// @Author  Wangwengang  2021/8/20 上午10:10
// @Update  Wangwengang  2021/8/20 上午10:10
package router

import (
	"time"

	"github.com/wwengg/arsenal/anet"
	"github.com/wwengg/arsenal/anet/impl"
	"github.com/wwengg/arsenal/atimer"
	"github.com/wwengg/arsenal/logger"
	"go.uber.org/zap"

	"github.com/wwengg/pigeon/internal/timer"
)

type KeepAlive struct {
	impl.BaseRouter
}

func (r *KeepAlive) Handle(request anet.Request) {
	tID, err := request.GetConnection().GetProperty("tID")
	if err != nil {
		logger.ZapLog.Debug("Keepalive GetProperty tId error", zap.Any("err", err))
		return
	}
	timer.AutoTimerSchedulerObj.CancelTimer(tID.(uint32))

	tId, _ := timer.AutoTimerSchedulerObj.CreateTimerAfter(atimer.NewDelayFunc(timer.StopConnect, []interface{}{request.GetConnection()}), 30*time.Second)
	request.GetConnection().RemoveProperty("tID")
	request.GetConnection().SetProperty("tID", tId)
}
