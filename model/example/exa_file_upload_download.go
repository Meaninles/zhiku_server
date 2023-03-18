package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ExaFileUploadAndDownload struct {
	global.GVA_MODEL
	Name        string `json:"name" gorm:"comment:文件名"` // 文件名
	Url         string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag         string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key         string `json:"key" gorm:"comment:编号"`   // 编号
	Description string `json:"description" gorm:"comment:内容描述"`
	Indexed     uint   `json:"indexed" gorm:"comment:已索引"`
	KBID        string `json:"kbid" gorm:"index;comment:知识库id"`
}

func (ExaFileUploadAndDownload) TableName() string {
	return "exa_file_upload_and_downloads"
}
