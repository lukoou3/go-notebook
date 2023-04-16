package shell

import (
	"net/http"
	"strconv"
)

// 模型的定义
type ShellCode struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Desc       string `json:"desc"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type ShellCodeSet struct {
	Total int          `json:"total"`
	Items []*ShellCode `json:"items"`
}

type QueryShellCodeRequest struct {
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	Keywords   string `json:"kws"`
}

func NewQueryShellCodeHTTP(r *http.Request) *QueryShellCodeRequest {
	req := NewQueryShellCodeRequest()
	// query string
	qs := r.URL.Query()
	pss := qs.Get("page_size")
	if pss != "" {
		req.PageSize, _ = strconv.Atoi(pss)
	}

	pns := qs.Get("page_number")
	if pns != "" {
		req.PageNumber, _ = strconv.Atoi(pns)
	}

	req.Keywords = qs.Get("kws")
	return req
}

func NewQueryShellCodeRequest() *QueryShellCodeRequest {
	return &QueryShellCodeRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}
