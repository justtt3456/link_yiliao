package service

import (
	"finance/app/admin/swag/response"
	"finance/common"
)

func FormatPage(page common.Page) response.Page {
	return response.Page{
		Page:     page.Page,
		PageSize: page.PageSize,
		Record:   page.Record,
		Total:    page.Total,
	}
}
