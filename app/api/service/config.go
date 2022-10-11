package service

import (
	"finance/app/api/swag/response"
	"finance/model"
)

type Config struct {
}

func (this Config) Get() response.Config {
	res := response.Config{}
	base := model.SetBase{}
	if base.Get() {
		res.Base = response.Base{
			AppName:      base.AppName,
			AppLogo:      base.AppLogo,
			VerifiedSend: float64(base.VerifiedSend) / model.UNITY,
			RegisterSend: float64(base.RegisterSend) / model.UNITY,
			OneSend:      float64(base.OneSend) / model.UNITY,
			TwoSend:      float64(base.TwoSend) / model.UNITY,
			ThreeSend:    float64(base.ThreeSend) / model.UNITY,
			SendDesc:     base.SendDesc,
			RegisterDesc: base.RegisterDesc,
			TeamDesc:     base.TeamDesc,
		}
	}
	funds := model.SetFunds{}
	if funds.Get() {
		res.Funds = response.Funds{
			RechargeStartTime:   funds.RechargeStartTime,
			RechargeEndTime:     funds.RechargeEndTime,
			RechargeMinAmount:   float64(funds.RechargeMinAmount) / 100,
			RechargeMaxAmount:   float64(funds.RechargeMaxAmount) / 100,
			RechargeQuickAmount: funds.RechargeQuickAmount,
			WithdrawStartTime:   funds.WithdrawStartTime,
			WithdrawEndTime:     funds.WithdrawEndTime,
			MustPassword:        funds.MustPassword,
			PasswordFreeze:      funds.PasswordFreeze,
			WithdrawMinAmount:   float64(funds.WithdrawMinAmount) / 100,
			WithdrawMaxAmount:   float64(funds.WithdrawMaxAmount) / 100,
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

			if v.ID == 2 {
				i := response.Kf{
					ID:        v.ID,
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
				ID:        v.ID,
				Name:      v.Name,
				Code:      v.Code,
				Icon:      v.Icon,
				IsDefault: v.IsDefault,
			}
			langSlice = append(langSlice, i)
		}
		res.Lang = langSlice
	}
	return res
}
