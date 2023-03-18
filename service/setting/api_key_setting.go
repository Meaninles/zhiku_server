package setting

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/setting"
	settingReq "github.com/flipped-aurora/gin-vue-admin/server/model/setting/request"
	"gorm.io/gorm"
)

type ApiKeySettingService struct {
}

// CreateApiKeySetting 创建ApiKeySetting记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiKeySettingService *ApiKeySettingService) CreateApiKeySetting(apiKeySetting setting.ApiKeySetting) (err error) {
	err = global.GVA_DB.Create(&apiKeySetting).Error
	return err
}

// DeleteApiKeySetting 删除ApiKeySetting记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiKeySettingService *ApiKeySettingService) DeleteApiKeySetting(apiKeySetting setting.ApiKeySetting) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&setting.ApiKeySetting{}).Where("id = ?", apiKeySetting.ID).Update("deleted_by", apiKeySetting.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&apiKeySetting).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteApiKeySettingByIds 批量删除ApiKeySetting记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiKeySettingService *ApiKeySettingService) DeleteApiKeySettingByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&setting.ApiKeySetting{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&setting.ApiKeySetting{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateApiKeySetting 更新ApiKeySetting记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiKeySettingService *ApiKeySettingService) UpdateApiKeySetting(apiKeySetting setting.ApiKeySetting) (err error) {
	err = global.GVA_DB.Save(&apiKeySetting).Error
	return err
}

// GetApiKeySetting 根据id获取ApiKeySetting记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiKeySettingService *ApiKeySettingService) GetApiKeySetting(id uint) (apiKeySetting setting.ApiKeySetting, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&apiKeySetting).Error
	return
}

func (apiKeySettingService *ApiKeySettingService) GetApiKeySettingInUsing() (apiKeySetting setting.ApiKeySetting, err error) {
	err = global.GVA_DB.Where("`using` = ?", 1).First(&apiKeySetting).Error
	return
}

// GetApiKeySettingInfoList 分页获取ApiKeySetting记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiKeySettingService *ApiKeySettingService) GetApiKeySettingInfoList(info settingReq.ApiKeySettingSearch) (list []setting.ApiKeySetting, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&setting.ApiKeySetting{})
	var apiKeySettings []setting.ApiKeySetting
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&apiKeySettings).Error
	return apiKeySettings, total, err
}
