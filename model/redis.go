package model

import (
	"china-russia/global"
	"errors"
	"time"
)

const (
	//redis hash key
	HashKeyLangConfig    string = "lang_config"
	HashKeyKfConfig      string = "kf_config"
	HashKeyBankConfig    string = "bank_config"
	HashKeyUsdtConfig    string = "usdt_config"
	HashKeyAlipayConfig  string = "alipay_config"
	HashKeyRollNotice    string = "roll_notice"
	HashKeyPopNotice     string = "pop_notice"
	HashKeyProduct       string = "product"
	HashKeyBanner        string = "banner"
	HashKeyPayment       string = "payment"
	HashKeyPayChannel    string = "pay_channel"
	StringKeyBaseConfig  string = "base_config"
	StringKeyFundsConfig string = "fund_config"
)

var (
	//redis string key
	StringKeyAdmin             string = "admin:%d"
	StringKeyAgent             string = "agent:%d"
	StringKeyOrder             string = "order:%d"
	StringKeyWithdraw          string = "withdraw:%d"
	StringKeyRecharge          string = "recharge:%d"
	StringKeyMember            string = "member:%d"
	StringKeyInviteCode        string = "invite_code:%s"
	StringKeyProductPrice      string = "product_price:%s"
	StringKeyProduct1MinKline  string = "product_1min_kline:%s"
	StringKeyProduct5MinKline  string = "product_5min_kline:%s"
	StringKeyProduct15MinKline string = "product_15min_kline:%s"
	StringKeyProduct30MinKline string = "product_30min_kline:%s"
	StringKeyProduct1HourKline string = "product_1hour_kline:%s"
	StringKeyProduct1DayKline  string = "product_1day_kline:%s"
	StringKeyWhiteIP           string = "white_ip"
	StringKeyRisk              string = "risk"
	//redis setnx
	LockKeyAdmin           string = "admin_lock:%d"
	LockKeyAgent           string = "agent_lock:%d"
	LockKeyOrder           string = "order_lock:%d"
	LockKeyWithdraw        string = "withdraw_lock:%d"
	LockKeyMember          string = "member_lock:%d"
	LockKeyInviteCode      string = "invite_code:%d"
	LockKeyProduct         string = "product_lock:%d"
	LockKeyPayChannel      string = "pay_channel_lock:%d"
	LockKeyPayment         string = "payment_lock:%d"
	LockKeyUpgrade         string = "upgrade_lock:%d"
	LockKeyRole            string = "role_lock:%d"
	LockKeyRecharge        string = "recharge_lock:%d"
	LockKeyBank            string = "bank_lock:%d"
	LockKeyNews            string = "news_lock:%d"
	LockKeyNotice          string = "notice_lock:%d"
	LockKeyMessage         string = "message_lock:%d"
	LockKeyMemberBank      string = "member_bank_lock:%d"
	LockKeyBanner          string = "banner_lock:%d"
	LockKeyBaseConfig      string = "base_config_lock:%d"
	LockKeyFundsConfig     string = "funds_config_lock:%d"
	LockKeyBankConfig      string = "bank_config_lock:%d"
	LockKeyAlipayConfig    string = "alipay_config_lock:%d"
	LockKeyUsdtConfig      string = "usdt_config_lock:%d"
	LockKeyKfConfig        string = "kf_config_lock:%d"
	LockKeyLangConfig      string = "lang_config_lock:%d"
	LockKeyLevel           string = "level_lock:%d"
	LockKeyMemberReport    string = "member_report_lock:%d"
	LockKeyAgentReport     string = "agent_report_lock:%d"
	LockKeyReport          string = "report_lock:%d"
	LockKeyRechargeMethod  string = "recharge_method_lock:%d"
	LockKeyWithdrawMethod  string = "withdraw_method_lock:%d"
	LockKeyProductCategory string = "product_category_lock:%d"
	LockKeyRisk            string = "risk_lock:%d"
	LockKeyCountry         string = "country_lock:%d"
	LockKeyHelp            string = "help_lock:%d"
	LockKeyMemberOptional  string = "member_optional_lock:%d"
	LockKeyMemberVerified  string = "member_verified_lock:%d"
	LockKeyWhiteIP         string = "white_ip_lock:%d"
	LockKeyPermission      string = "permission_lock:%d"
)

type Redis struct {
}

func (Redis) Lock(key string) error {
	ch := time.After(time.Second * 5)
	for {
		select {
		case <-ch:
			return errors.New("请求超时,请重试")
		default:
			if global.REDIS.SetNX(key, 1, time.Second*5).Val() {
				return nil
			}
		}
	}
}
func (Redis) Unlock(key string) {
	global.REDIS.Del(key)
}
