package service

import (
	"china-russia/app/api/swag/response"
	"china-russia/model"
	"time"
)

type Config struct {
}

func (this Config) Get() response.Config {
	res := response.Config{}
	base := model.SetBase{}
	if base.Get() {
		equityScore := model.StatusClose
		if time.Now().Unix() >= base.EquityStartDate {
			equityScore = model.StatusOk
		}
		res.Base = response.Base{
			AppName:       base.AppName,
			AppLogo:       base.AppLogo,
			VerifiedSend:  base.VerifiedSend,
			RegisterSend:  base.RegisterSend,
			OneSend:       base.OneSend,
			TwoSend:       base.TwoSend,
			ThreeSend:     base.ThreeSend,
			SendDesc:      base.SendDesc,
			RegisterDesc:  base.RegisterDesc,
			TeamDesc:      base.TeamDesc,
			IsEquityScore: equityScore,
			DownloadUrl:   base.DownloadUrl,
		}
	}
	funds := model.SetFunds{}
	if funds.Get() {
		res.Funds = response.Funds{
			RechargeStartTime:   funds.RechargeStartTime,
			RechargeEndTime:     funds.RechargeEndTime,
			RechargeMinAmount:   funds.RechargeMinAmount,
			RechargeMaxAmount:   funds.RechargeMaxAmount,
			RechargeQuickAmount: funds.RechargeQuickAmount,
			WithdrawStartTime:   funds.WithdrawStartTime,
			WithdrawEndTime:     funds.WithdrawEndTime,
			MustPassword:        funds.MustPassword,
			PasswordFreeze:      funds.PasswordFreeze,
			WithdrawMinAmount:   funds.WithdrawMinAmount,
			WithdrawMaxAmount:   funds.WithdrawMaxAmount,
			WithdrawFee:         funds.WithdrawFee,
			ProductFee:          funds.ProductFee,
			ProductQuickAmount:  funds.ProductQuickAmount,
		}
	}
	kf := model.SetKf{
		Status: model.StatusOk,
	}
	kfs := kf.List(true)
	if len(kfs) > 0 {
		kfSlice := make([]response.Kf, 0)
		for _, v := range kfs {

			if v.Id == 2 {
				i := response.Kf{
					Id:        v.Id,
					Name:      v.Name,
					StartTime: v.StartTime,
					EndTime:   v.EndTime,
					Link:      v.Link,
				}
				kfSlice = append(kfSlice, i)
			}

		}
		res.Kf = kfSlice
	}
	lang := model.SetLang{
		Status: model.StatusOk,
	}
	langs := lang.List(true)
	if len(langs) > 0 {
		langSlice := make([]response.Lang, 0)
		for _, v := range langs {
			i := response.Lang{
				Id:        v.Id,
				Name:      v.Name,
				Code:      v.Code,
				Icon:      v.Icon,
				IsDefault: v.IsDefault,
			}
			langSlice = append(langSlice, i)
		}
		res.Lang = langSlice
	}
	equity := model.Equity{}
	if equity.Get(true) {
		res.IsOpen = true
	}
	return res
}
