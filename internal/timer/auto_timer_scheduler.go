// @Title
// @Description
// @Author  Wangwengang  2021/8/21 下午3:31
// @Update  Wangwengang  2021/8/21 下午3:31
package timer

import (
	"fmt"

	"github.com/wwengg/arsenal/anet"
	"github.com/wwengg/arsenal/atimer"
	"github.com/wwengg/arsenal/logger"
	"go.uber.org/zap"

	"github.com/wwengg/pigeon/internal"
)

var AutoTimerSchedulerObj *atimer.TimerScheduler

func init() {
	AutoTimerSchedulerObj = atimer.NewAutoExecTimerScheduler()
}

func StopConnect(v ...interface{}) {
	logger.ZapLog.Info("StopConnect func Start")
	conn := v[0].(anet.Connection)
	logger.ZapLog.Info("StopConnect func Start", zap.Any("conn", conn))
	//cID := conn.GetProperty("cID")
	cID, err := conn.GetProperty("cID")
	if err != nil {
		logger.ZapLog.Error("StopConnect cID error", zap.Any("err", err))
		conn.Stop()
		return
	}
	if c := internal.GlobalMgrObj.GetClientByCID(cID.(uint64)); c == nil {
		fmt.Println("心跳超时 30秒,客户端已移除，忽略")
		return
	}
	fmt.Println("心跳超时 30秒")
	conn.Stop()
}
