// @Title
// @Description
// @Author  Wangwengang  2021/8/25 上午9:20
// @Update  Wangwengang  2021/8/25 上午9:20
package rpcx

import (
	"context"

	"github.com/wwengg/arsenal/sdk/rpcx"
	"github.com/wwengg/proto/common"
	"github.com/wwengg/proto/pigeon"

	"github.com/wwengg/pigeon/internal"
)

func Run() {
	// new rpcx server
	s := rpcx.NewRpcxServer()

	s.RegisterName(internal.DefKey, new(PigeonServiceImpl), "")
	s.RegisterName("pigeon", new(PigeonServiceImpl), "")

	s.Serve()
}

type PigeonService interface {
	// PigeonService can be used for interface verification.

	// PushMsgByUser is server rpc method as defined
	PushMsgByUser(ctx context.Context, args *pigeon.PushMsgByUserArgs, reply *common.Err) (err error)

	// PushMsgByRoom is server rpc method as defined
	PushMsgByRoom(ctx context.Context, args *pigeon.PushMsgByRoomArgs, reply *common.Err) (err error)

	// PushMsgByClient is server rpc method as defined
	PushMsgByClient(ctx context.Context, args *pigeon.PushMsgByClientArgs, reply *common.Err) (err error)

	// Broadcast is server rpc method as defined
	Broadcast(ctx context.Context, args *pigeon.BroadcastArgs, reply *common.Err) (err error)
}

type PigeonServiceImpl struct {
}

// PushMsgByUser is server rpc method as defined
func (s *PigeonServiceImpl) PushMsgByUser(ctx context.Context, args *pigeon.PushMsgByUserArgs, reply *common.Err) (err error) {

	// TODO: add business logics

	// TODO: setting return values
	*reply = common.Err{}

	return nil
}

// PushMsgByRoom is server rpc method as defined
func (s *PigeonServiceImpl) PushMsgByRoom(ctx context.Context, args *pigeon.PushMsgByRoomArgs, reply *common.Err) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = common.Err{}

	return nil
}

// PushMsgByClient is server rpc method as defined
func (s *PigeonServiceImpl) PushMsgByClient(ctx context.Context, args *pigeon.PushMsgByClientArgs, reply *common.Err) (err error) {
	if c := internal.GlobalMgrObj.GetClientByCID(args.CId);c != nil{
		c.SendMsg(args.Op, args.Proto)
	}
	*reply = common.Err{}

	return nil
}

// Broadcast is server rpc method as defined
func (s *PigeonServiceImpl) Broadcast(ctx context.Context, args *pigeon.BroadcastArgs, reply *common.Err) (err error) {
	internal.GlobalMgrObj.Broadcast(args.Op, args.Proto)
	*reply = common.Err{}

	return nil
}
