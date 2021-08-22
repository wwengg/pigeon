/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"errors"
	"time"

	"github.com/smallnest/rpcx/client"
	"github.com/wwengg/arsenal/anet"
	"github.com/wwengg/arsenal/anet/impl"
	"github.com/wwengg/arsenal/atimer"
	"github.com/wwengg/arsenal/config"
	"github.com/wwengg/arsenal/logger"
	"github.com/wwengg/arsenal/sdk/rpcx"
	"go.uber.org/zap"

	"github.com/wwengg/pigeon/internal"
	"github.com/wwengg/pigeon/internal/timer"
	"github.com/wwengg/pigeon/router"
)

func OnConnectionAdd(conn anet.Connection) {
	cID := conn.GetConnID()
	conn.SetProperty("cID", cID)

	tId, _ := timer.AutoTimerSchedulerObj.CreateTimerAfter(atimer.NewDelayFunc(timer.StopConnect, []interface{}{conn}), 30*time.Second)
	conn.SetProperty("tID", tId)

	logger.ZapLog.Info("=====> client arrived ====", zap.Any("cID", cID))
}

func OnConnectionLost(conn anet.Connection) {
	defer logger.ZapLog.Info("====> Client left ===== success")

	cID, err := conn.GetProperty("cID")
	if err != nil {
		logger.ZapLog.Debug("====> Client noCID left =====", zap.Error(errors.New("noCID")))
		return
	}
	logger.ZapLog.Debug("=====> Client left <=====", zap.Any("cID", cID))
	if c := internal.GlobalMgrObj.GetClientByCID(cID.(uint64)); c != nil {
		c.ConnectionLost()
	}
}

func main() {
	// Init config
	config.Viper()

	// Init logger
	logger.Setup()

	// new tcpServer
	ts := impl.NewServer()
	ts.AddRouter(0, new(router.KeepAlive))

	ts.SetOnConnStop(OnConnectionLost)
	ts.SetOnConnStart(OnConnectionAdd)
	ts.Serve()

	// new rpcx client
	rpcx.RpcxClientsObj.SetupServiceDiscovery()
	rpcx.RpcxClientsObj.SetFailMode(client.Failover)

	// new rpcx server
	s := rpcx.NewRpcxServer()

	s.RegisterName("pigeon", nil, "")

	s.Serve()

}
