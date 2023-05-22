package service

import (
	"china-russia/app/admin/swag/request"
	"china-russia/app/admin/swag/response"
	"china-russia/common"
	"china-russia/extends"
	"china-russia/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AdminService struct{}
type AdminListService struct {
	request.AdminListRequest
}
type AdminInsertService struct {
	request.AdminInsertRequest
}
type AdminUpdateService struct {
	request.AdminUpdateRequest
}
type AdminRemoveService struct {
	request.AdminRemoveRequest
}
type AdminPasswordService struct {
	request.AdminPasswordRequest
}

func (this AdminPasswordService) Password(admin model.Admin) error {
	if this.OldPassword == "" {
		return errors.New("原密码不能为空")
	}
	if this.NewPassword == "" {
		return errors.New("新密码不能为空")
	}
	if this.RePassword != this.NewPassword {
		return errors.New("两次密码不一致")
	}

	if admin.Password != common.Md5String(this.OldPassword+admin.Salt) {
		return errors.New("原密码不一致")
	}
	admin.Password = common.Md5String(this.NewPassword + admin.Salt)
	return admin.Update("password")
}
func (this AdminListService) PageList() (response.AdminListData, error) {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize > common.MaxPageSize || this.PageSize < common.MinPageSize {
		this.PageSize = common.DefaultPageSize
	}
	m := model.Admin{}
	where, args := this.getWhere()
	list, page := m.PageList(where, args, this.Page, this.PageSize)
	res := make([]response.AdminInfo, 0)
	for _, v := range list {
		//p := AdminPermissionService{}
		//tree, _ := p.Tree(v.Role)
		item := response.AdminInfo{
			AdminId:    v.Id,
			Username:   v.Username,
			Token:      v.Token,
			Role:       v.Role,
			RoleName:   v.RoleTab.RoleName,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			LoginIp:    v.LoginIp,
			RegisterIp: v.RegisterIp,
			Operator:   v.Operator,
			//Permission: tree,
		}
		res = append(res, item)
	}
	return response.AdminListData{
		List: res,
		Page: FormatPage(page),
	}, nil
}
func (this AdminListService) getWhere() (string, []interface{}) {
	where := map[string]interface{}{}
	if this.Name != "" {
		where["a.name"] = this.Name
	}
	build, vals, err := common.WhereBuild(where)
	if err != nil {
		logrus.Error(err)
	}
	return build, vals
}
func (this AdminInsertService) Insert(c *gin.Context, operator model.Admin) (*response.AdminGoogle, error) {
	if this.Username == "" {
		return nil, errors.New("账号不能为空")
	}
	if this.Password == "" {
		return nil, errors.New("密码不能为空")
	}
	if this.RePassword != this.Password {
		return nil, errors.New("两次密码不一致")
	}
	if this.Role == 0 {
		return nil, errors.New("角色不能为空")
	}
	//是否存在
	admin := model.Admin{
		Username: this.Username,
	}
	if admin.Get() {
		return nil, errors.New("账号已存在")
	}
	salt := common.RandStringRunes(6)
	google := extends.NewGoogleAuth()
	secret := google.Secret()
	m := model.Admin{
		Username:   this.Username,
		Password:   common.Md5String(this.Password + salt),
		Salt:       salt,
		Token:      common.RandStringRunes(32),
		Role:       this.Role,
		LoginIp:    c.ClientIP(),
		RegisterIp: c.ClientIP(),
		Operator:   operator.Id,
		GoogleAuth: secret,
	}
	err := m.Insert()
	if err != nil {
		return nil, err
	}
	return &response.AdminGoogle{
		Username: this.Username,
		Qrcode:   google.QrcodeUrl(this.Username, secret),
	}, nil
}
func (AdminService) Info(admin model.Admin) response.AdminInfo {
	jwtService := extends.JwtUtils{}
	token := jwtService.NewToken(admin.Id, admin.Token)
	//所有权限
	permission := model.Permission{}
	permissions := permission.List()
	p := RolePermissionTree{}
	return response.AdminInfo{
		AdminId:    admin.Id,
		Username:   admin.Username,
		Token:      token,
		Role:       admin.Role,
		CreateTime: admin.CreateTime,
		UpdateTime: admin.UpdateTime,
		LoginIp:    admin.LoginIp,
		RegisterIp: admin.RegisterIp,
		Operator:   admin.Operator,
		Permission: p.Tree(permissions, admin.Role),
	}
}

func (this AdminUpdateService) Update() error {
	if this.AdminId == 0 {
		return errors.New("参数错误")
	}
	admin := model.Admin{Id: this.AdminId}
	if !admin.Get() {
		return errors.New("管理员不存在")
	}
	if this.Password != "" {

		admin.Password = common.Md5String(this.Password + admin.Salt)
	}
	if this.Role != 0 {
		admin.Role = this.Role
	}
	return admin.Update("password", "role")
}
func (this AdminRemoveService) Remove() error {
	if this.AdminId == 0 {
		return errors.New("参数错误")
	}
	admin := model.Admin{Id: this.AdminId}
	if !admin.Get() {
		return errors.New("管理员不存在")
	}
	return admin.Remove()
}

type AdminGoogleService struct {
	request.AdminGoogleRequest
}

func (this AdminGoogleService) Google(admin model.Admin) (*response.AdminGoogle, error) {
	if this.AdminId == 0 {
		return nil, errors.New("参数错误")
	}
	a := model.Admin{Id: this.AdminId}
	if !a.Get() {
		return nil, errors.New("管理员不存在")
	}
	google := extends.NewGoogleAuth()
	b, err := google.VerifyCode(admin.GoogleAuth, this.GoogleCode)
	if err != nil {
		return nil, errors.New("验证码错误")
	}
	if !b {
		return nil, errors.New("验证码错误")
	}
	secret := google.Secret()
	a.GoogleAuth = secret
	err = a.Update("google_auth")
	if err != nil {
		return nil, err
	}
	return &response.AdminGoogle{
		Username: admin.Username,
		Qrcode:   google.QrcodeUrl("管理后台"+a.Username, secret),
	}, nil
}
