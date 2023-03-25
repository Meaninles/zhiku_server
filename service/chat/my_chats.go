package chat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/chat"
	chatReq "github.com/flipped-aurora/gin-vue-admin/server/model/chat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/knowledgeBase"
	"gorm.io/gorm"
)

type MyChatsService struct {
}

// CreateMyChats 创建MyChats记录
// Author [piexlmax](https://github.com/piexlmax)
func (myChatsService *MyChatsService) CreateMyChats(myChats chat.MyChats) (err error) {
	err = global.GVA_DB.Create(&myChats).Error
	return err
}

// DeleteMyChats 删除MyChats记录
// Author [piexlmax](https://github.com/piexlmax)
func (myChatsService *MyChatsService) DeleteMyChats(myChats chat.MyChats) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&chat.MyChats{}).Where("id = ?", myChats.ID).Update("deleted_by", myChats.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&myChats).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMyChatsByIds 批量删除MyChats记录
// Author [piexlmax](https://github.com/piexlmax)
func (myChatsService *MyChatsService) DeleteMyChatsByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&chat.MyChats{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&chat.MyChats{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMyChats 更新MyChats记录
// Author [piexlmax](https://github.com/piexlmax)
func (myChatsService *MyChatsService) UpdateMyChats(myChats chat.MyChats) (err error) {
	err = global.GVA_DB.Save(&myChats).Error
	return err
}

// GetMyChats 根据id获取MyChats记录
// Author [piexlmax](https://github.com/piexlmax)
func (myChatsService *MyChatsService) GetMyChats(id uint) (myChats chat.MyChats, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&myChats).Error
	return
}

// GetMyChatsInfoList 分页获取MyChats记录
// Author [piexlmax](https://github.com/piexlmax)
func (myChatsService *MyChatsService) GetMyChatsInfoList(info chatReq.MyChatsSearch) (list []chat.MyChats, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&chat.MyChats{})
	var myChatss []chat.MyChats
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&myChatss).Error
	return myChatss, total, err
}

func (myChatsService *MyChatsService) GetKBIds(userId uint) (kbList []knowledgeBase.UserKnowledgeBase, err error) {
	err = global.GVA_DB.Model(&knowledgeBase.UserKnowledgeBase{}).Where("user_id = ?", userId).Find(&kbList).Error
	return
}
