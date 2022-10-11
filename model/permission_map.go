package model

import (
	"errors"
	"finance/global"
	"github.com/sirupsen/logrus"
)

type PermissionMap struct {
	RoleID       int        `gorm:"column:role_id"`       //
	PermissionID int        `gorm:"column:permission_id"` //
	Permission   Permission `gorm:"foreignKey:PermissionID"`
}

// TableName sets the insert table name for this struct type
func (p *PermissionMap) TableName() string {
	return "c_permission_map"
}

//角色权限
func (this PermissionMap) RolePermission() ([]Permission, error) {
	res := make([]Permission, 0)
	if this.RoleID == 0 {
		return nil, errors.New("参数错误")
	}
	tx := global.DB.Table("qd_permission_map a").Select("id").Joins("left join qd_permission b on a.permission_id = b.id").Where("a.role_id = ?", this.RoleID).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}
func (this PermissionMap) ValidPermission(route string) bool {
	if this.RoleID == 0 {
		return false
	}
	if this.RoleID == 1 {
		return true
	}
	tx := global.DB.Model(this).Joins("Permission").Where("Permission.backend = ? and "+this.TableName()+".role_id = ?", route, this.RoleID).First(&this)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		return false
	}
	return true
}

//角色权限
func (this PermissionMap) AdminPermission() ([]PermissionMap, error) {
	res := make([]PermissionMap, 0)
	if this.RoleID == 0 {
		return nil, errors.New("参数错误")
	}
	if this.RoleID == 1 { //超级管理员
		tx := global.DB.Model(this).Order("permission_id asc").Find(&res)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := global.DB.Model(this).Where("role_id = ?", this.RoleID).Order("permission_id asc").Find(&res)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return res, nil
}
func (this *PermissionMap) Remove() error {
	res := global.DB.Where("role_id = ?", this.RoleID).Delete(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (this *PermissionMap) InsertAll(data []PermissionMap) error {
	res := global.DB.Create(data)
	if res.Error != nil {
		logrus.Error(res.Error)
		return res.Error
	}
	return nil
}
