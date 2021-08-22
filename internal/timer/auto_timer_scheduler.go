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
)

var AutoTimerSchedulerObj *atimer.TimerScheduler

func init(){
	AutoTimerSchedulerObj = atimer.NewAutoExecTimerScheduler()
}

func StopConnect(v ...interface{}){
	conn := v[0].(anet.Connection)
	//cID := conn.GetProperty("cID")
	fmt.Println("心跳超时 30秒")
	_,err := conn.GetProperty("cID")
	if err != nil {
		logger.ZapLog.Error("StopConnect cID error",zap.Any("err",err))
		conn.Stop()
		return
	}
	_,err = conn.GetProperty("uID")
	if err != nil {
		logger.ZapLog.Error("StopConnect uID error",zap.Any("err",err))
		conn.Stop()
		return
	}
	conn.Stop()
}