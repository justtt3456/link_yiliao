package extends

import (
	"china-russia/common"
	"china-russia/global"
	"errors"
	"fmt"
	"time"
)

type SmsBao struct {
}

func (SmsBao) Send(mobile string) error {
	if !common.IsMobile(mobile, global.Language) {
		return errors.New("手机号必传")
	}
	code := common.RandIntRunes(4)
	if global.REDIS.Get(mobile).Val() != "" {
		return errors.New("验证码已发送，请间隔5分钟再尝试")
	}
	global.REDIS.Set(mobile, code, 300*time.Second)
	msg := fmt.Sprintf(global.CONFIG.Sms.Phone.Sign, code)
	param := map[string]string{
		"u": global.CONFIG.Sms.Phone.Username,
		"p": common.Md5String(global.CONFIG.Sms.Phone.Password),
		"m": mobile,
		"c": msg,
	}
	b, err := common.GetParam(global.CONFIG.Sms.Phone.Url, param, nil, nil)
	if err != nil {
		return err
	}
	if string(b) != "0" {
		return errors.New("发送失败")
	}
	return nil
}
