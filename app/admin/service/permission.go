package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/model"
)

type PermissionService struct {
	request.PermissionRequest
}

func (this PermissionService) List() []response.PermissionTree {
	m := model.Permission{}
	//所有权限
	list := m.List()
	if len(list) == 0 {
		return nil
	}
	return buildPermissionTree(list, 0)
}

//所有权限树
func buildPermissionTree(res []model.Permission, pid int) []response.PermissionTree {
	m := make([]response.PermissionTree, 0)
	for _, v := range res {
		s := response.PermissionTree{}
		if v.Pid == pid {
			children := buildPermissionTree(res, v.ID)
			if children != nil {
				s.ID = v.ID
				s.PID = v.Pid
				s.Label = v.Label
				s.Frontend = v.Frontend
				s.Backend = v.Backend
				s.IsBtn = v.IsBtn
				s.Sort = v.Sort
				s.Children = children
			}
			m = append(m, s)
		}
	}
	return m
}

type PermissionCreateService struct {
	request.PermissionCreateRequest
}

func (this PermissionCreateService) Create() error {
	if this.Label == "" {
		return errors.New("权限名称不能为空")
	}
	p := model.Permission{
		Frontend: this.Frontend,
		Backend:  this.Backend,
		Label:    this.Label,
		Pid:      this.PID,
		IsBtn:    this.IsBtn,
		Sort:     this.Sort,
	}
	return p.Insert()
}

type PermissionUpdateService struct {
	request.PermissionUpdateRequest
}

func (this PermissionUpdateService) Update() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	if this.Label == "" {
		return errors.New("权限名称不能为空")
	}
	p := model.Permission{
		ID: this.ID,
	}
	if !p.Get() {
		return errors.New("记录不存在")
	}
	p.Frontend = this.Frontend
	p.Backend = this.Backend
	p.Label = this.Label
	p.Pid = this.PID
	p.IsBtn = this.IsBtn
	p.Sort = this.Sort
	return p.Update("frontend", "backend", "label", "pid", "is_btn", "sort")
}

type PermissionRemoveService struct {
	request.PermissionRemoveRequest
}

func (this PermissionRemoveService) Remove() error {
	if this.ID == 0 {
		return errors.New("参数错误")
	}
	p := model.Permission{
		ID: this.ID,
	}
	return p.Remove()
}