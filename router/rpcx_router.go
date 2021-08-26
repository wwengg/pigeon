// @Title
// @Description
// @Author  Wangwengang  2021/8/26 上午8:45
// @Update  Wangwengang  2021/8/26 上午8:45
package router

import (
	"context"
	"strconv"

	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/share"
	"github.com/wwengg/arsenal/anet"
	"github.com/wwengg/arsenal/anet/impl"
	"github.com/wwengg/arsenal/config"
	"github.com/wwengg/arsenal/logger"
	"github.com/wwengg/arsenal/sdk/rpcx"
	"github.com/wwengg/proto/identity"
)

type RpcxRouter struct {
	impl.RpcxRouter
}

func (r *RpcxRouter) Handle(request anet.Request) {
	cID, err := request.GetConnection().GetProperty("cID")
	if err != nil {
		logger.ZapLog.Debug("no cID")
		return
	}
	uID, err := request.GetConnection().GetProperty("uID")
	if err != nil {
		logger.ZapLog.Debug("no uID")
		return
	}
	op := request.GetMsgID()
	rpcxRouterMap := config.ConfigHub.RpcxRouterMap
	router := rpcxRouterMap[op]

	xclient, err := rpcx.RpcxClientsObj.GetXClient(router.ServicePath)
	if err != nil {
		logger.ZapLog.Error(err.Error())
		return
	}

	req := protocol.NewMessage()
	req.SetMessageType(protocol.Request)

	req.ServicePath = router.ServicePath // servicePath
	req.ServiceMethod = router.ServiceMethod

	// 获取雪花id
	xclient2, err := rpcx.RpcxClientsObj.GetXClient("Identity")
	if err != nil {
		logger.ZapLog.Error("Identity service not found")
		return
	}

	identityClient := identity.NewIdentityClient(xclient2)
	reply, err := identityClient.GetId(context.Background(), nil)
	if err != nil {
		logger.ZapLog.Error(err.Error())
		return
	}
	req.SetSeq(uint64(reply.Id)) // seq

	req.SetOneway(router.Oneway)
	req.Payload = request.GetData()
	req.SetSerializeType(protocol.ProtoBuffer) //  0 raw bytes, 1 JSON, 2 protobuf, 3 msgpack
	var data = make(map[string]string)
	data["CID"] = strconv.Itoa(int(cID.(uint64)))
	data["UID"] = strconv.Itoa(int(uID.(uint64)))
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, data)
	_, payload, _ := xclient.SendRaw(ctx, req)
	if err := request.GetConnection().SendMsg(router.BackOp, payload); err != nil {
		logger.ZapLog.Error(err.Error())
		return
	}
}
