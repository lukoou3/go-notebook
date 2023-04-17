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

func NewShellCode() *ShellCode {
	return &ShellCode{}
}

func NewShellCodeSet() *ShellCodeSet {
	return &ShellCodeSet{
		Items: []*ShellCode{},
	}
}

type ShellCodeSet struct {
	Total int          `json:"total"`
	Items []*ShellCode `json:"items"`
}

func (s *ShellCodeSet) Add(item *ShellCode) {
	s.Items = append(s.Items, item)
}

type QueryShellCodeRequest struct {
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	Keywords   string `json:"kws"`
}

func (req *QueryShellCodeRequest) GetPageSize() uint {
	return uint(req.PageSize)
}

func (req *QueryShellCodeRequest) OffSet() int64 {
	return int64((req.PageNumber - 1) * req.PageSize)
}

func NewQueryShellCodeHTTP(r *http.Request) *QueryShellCodeRequest {
	req := NewQueryShellCodeRequest()
	// query string
	qs := r.URL.Query()
	pss := qs.Get("page_size")
	if pss != "" {
		if n, err := strconv.Atoi(pss); err == nil {
			req.PageSize = n
		}
	}

	pns := qs.Get("page_number")
	if pns != "" {
		if n, err := strconv.Atoi(pns); err == nil {
			req.PageNumber = n
		}
	}

	req.Keywords = qs.Get("kws")
	return req
}

func NewQueryShellCodeRequest() *QueryShellCodeRequest {
	return &QueryShellCodeRequest{
		PageSize:   10,
		PageNumber: 1,
	}
}
