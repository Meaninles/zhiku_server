// 自动生成模板ApiKeySetting
package setting

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ApiKeySetting 结构体
type ApiKeySetting struct {
	global.GVA_MODEL
	ServiceName string `json:"serviceName" form:"serviceName" gorm:"column:service_name;comment:服务提供商的名称;"`
	ApiKey      string `json:"apiKey" form:"apiKey" gorm:"column:api_key;comment:api key设置;"`
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
	Using       uint   `gorm:"index;column:using;comment:正在使用"`
}

// TableName ApiKeySetting 表名
func (ApiKeySetting) TableName() string {
	return "api_key_setting"
}
