package service

import (
	"china-russia/app/api/swag/response"
	"china-russia/common"
)

func FormatPage(page common.Page) response.Page {
	return response.Page{
		Page:     page.Page,
		PageSize: page.PageSize,
		Record:   page.Record,
		Total:    page.Total,
	}
}
