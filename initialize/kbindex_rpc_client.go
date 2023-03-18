package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	pb "github.com/flipped-aurora/gin-vue-admin/server/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	address = ":3743"
)

func KBINDEX_RPC_CLIENT() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		global.GVA_LOG.Fatal("无法连接到rpc", zap.Error(err))
	}
	global.GVA_KBINDEX_RPC_CLIENT = pb.NewKBIndexServiceClient(conn)
}
