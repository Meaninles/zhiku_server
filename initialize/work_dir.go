package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"os"
)

func InitCurrentWorkDir() {
	dir, err := os.Getwd()
	if err != nil || dir == "" {
		global.GVA_LOG.Fatal("无法获取当前目录", zap.Error(err))
	}
	global.CURRENT_WORK_DIR = dir
}
