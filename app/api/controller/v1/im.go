package v1

import (
	"china-russia/app/api/controller"
	"china-russia/global"
	"china-russia/lang"
	"china-russia/model"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type IMController struct {
	controller.AuthController
}

func (this IMController) Link(c *gin.Context) {
	user := this.MemberInfo(c)
	err, s := LoginImUser(*user)
	if err != nil {
		//用户不存在
		if s == "account_not_found" {
			err = RegisterIMUser(*user)
			if err != nil {
				this.Json(c, 10001, err.Error(), nil)
				return
			} else {
				err, s = LoginImUser(*user)
				if err != nil {
					this.Json(c, 10001, err.Error(), nil)
					return
				}
				this.Json(c, 10001, err.Error(), nil)
				return
			}
		}
		this.Json(c, 10001, err.Error(), nil)
		return
	}
	this.Json(c, 0, "ok", s)
	return
}

func RegisterIMUser(user model.Member) error {
	resp, err := http.PostForm(global.CONFIG.IM.Api, url.Values{
		"api_secret_key": {global.CONFIG.IM.Key},
		"add":            {"site_users"},
		"full_name":      {user.Username},
		"username":       {user.Username},
		"email_address":  {user.Username + "@email.com"},
		"password":       {user.Password},
	})
	//resp, err := client.Get(urlPath)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println("注册用户返回：", string(body))
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	m := map[string]interface{}{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	if _, ok := m["success"]; ok {
		if m["success"].(bool) {
			return nil
		}
		return errors.New(m["error_message"].(string))
	}
	return errors.New(m["error_message"].(string))
}
func LoginImUser(user model.Member) (error, string) {
	resp, err := http.PostForm(global.CONFIG.IM.Api, url.Values{
		"api_secret_key": {global.CONFIG.IM.Key},
		"add":            {"login_session"},
		"username":       {user.Username},
		"email_address":  {user.Username + "@email.com"},
	})
	//resp, err := client.Get(urlPath)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return err, ""
	}
	log.Println("登录用户返回：", string(body))
	defer resp.Body.Close()
	m := map[string]interface{}{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		logrus.Error(err)
		return err, ""
	}
	if v, ok := m["auto_login_url"]; ok {
		return nil, v.(string)
	}
	if _, ok := m["error_key"]; ok {
		return errors.New(lang.Lang("System error, please contact the administrator")), m["error_key"].(string)
	}
	return errors.New(lang.Lang("System error, please contact the administrator")), ""
}
