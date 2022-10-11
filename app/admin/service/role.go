package service

import (
	"errors"
	"finance/app/admin/swag/request"
	"finance/app/admin/swag/response"
	"finance/common"
	"finance/model"
)

type RoleListService struct {
	request.RoleListRequest
}
type RoleCreateService struct {
	request.RoleCreateRequest
}
type RoleUpdateService struct {
	request.RoleUpdateRequest
}
type RoleRemoveService struct {
	request.RoleRemoveRequest
}

func (this RoleListService) PageList() (response.RoleData, error) {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Role{}
	list, page := m.PageList(this.Page, this.PageSize)
	res := make([]response.RoleInfo, 0)
	//所有权限
	p := model.Permission{}
	permissions := p.List()
	s := RolePermissionTree{}
	for _, v := range list {
		//用户权限
		item := response.RoleInfo{
			RoleID:      v.RoleID,
			RoleName:    v.RoleName,
			Status:      v.Status,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
			Permissions: s.Tree(permissions, v.RoleID),
		}
		res = append(res, item)
	}
	return response.RoleData{
		List: res,
		Page: FormatPage(page),
	}, nil
}
func (this RoleCreateService) Create() error {
	if this.RoleName == "" {
		return errors.New("角色名不能为空")
	}
	if len(this.Ids) == 0 {
		return errors.New("角色权限不能为空")
	}
	m := model.Role{
		RoleName: this.RoleName,
	}
	if m.Get() {
		return errors.New("角色名已存在")
	}
	if err := m.Insert(); err != nil {
		return err
	}
	//添加角色权限
	pm := model.PermissionMap{}
	data := make([]model.PermissionMap, 0)
	for _, v := range this.Ids {
		i := model.PermissionMap{
			RoleID:       m.RoleID,
			PermissionID: v,
		}
		data = append(data, i)
	}
	return pm.InsertAll(data)
}
func (this RoleUpdateService) Update() error {
	if this.RoleID == 0 {
		return errors.New("参数错误")
	}
	if this.RoleID == 1 {
		return errors.New("超级管理员不能修改")
	}
	if this.RoleName == "" {
		return errors.New("角色名不能为空")
	}
	if len(this.Ids) == 0 {
		return errors.New("角色权限不能为空")
	}
	m := model.Role{
		RoleName: this.RoleName,
	}
	if !m.Get() {
		return errors.New("角色名不存在")
	}
	m.RoleID = this.RoleID
	if err := m.Update("role_name"); err != nil {
		return err
	}
	//更新角色权限
	pm := model.PermissionMap{
		RoleID: this.RoleID,
	}
	if err := pm.Remove(); err != nil {
		return err
	}
	data := make([]model.PermissionMap, 0)
	for _, v := range this.Ids {
		i := model.PermissionMap{
			RoleID:       this.RoleID,
			PermissionID: v,
		}
		data = append(data, i)
	}
	return pm.InsertAll(data)

}
func (this RoleRemoveService) Remove() error {
	if this.RoleID == 0 {
		return errors.New("参数错误")
	}
	if this.RoleID == 1 {
		return errors.New("超级管理员不能删除")
	}
	m := model.Role{
		RoleID: this.RoleID,
	}
	if err := m.Remove(); err != nil {
		return err
	}
	pm := model.PermissionMap{
		RoleID: this.RoleID,
	}
	return pm.Remove()
}

type RolePermissionTree struct {
}

func (this RolePermissionTree) Tree(permissions []model.Permission, roleID int) []response.PermissionTree {
	m := model.PermissionMap{
		RoleID: roleID,
	}
	rolePermissions, _ := m.AdminPermission()
	return this.rolePermissionTree(permissions, rolePermissions, roleID, 0)
}

//所有权限树
func (this RolePermissionTree) rolePermissionTree(res []model.Permission, rolePermission []model.PermissionMap, roleID, pid int) []response.PermissionTree {
	m := make([]response.PermissionTree, 0)
	for _, v := range res {
		if v.Pid == pid {
			children := this.rolePermissionTree(res, rolePermission, roleID, v.ID)
			s := response.PermissionTree{
				PermissionInfo: response.PermissionInfo{
					ID:       v.ID,
					PID:      v.Pid,
					Label:    v.Label,
					Frontend: v.Frontend,
					Backend:  v.Backend,
					IsBtn:    v.IsBtn,
					Checked:  false,
					Sort:     v.Sort,
				},
				Children: children,
			}
			if roleID == 1 {
				s.Checked = true
			} else {
				for _, rv := range rolePermission {
					if rv.PermissionID == v.ID {
						s.Checked = true
					}
				}
			}
			m = append(m, s)
		}
	}
	return m
}
