package datasource

import (
	"net/http"
	"strconv"
)

type Datasource struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Alias        string `json:"alias"`
	Introduction string `json:"introduction"`
	Sql          string `json:"sql"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
	Cate         int    `json:"cate"`
}

func NewDatasource() *Datasource {
	return &Datasource{}
}

func NewDatasourceSet() *DatasourceSet {
	return &DatasourceSet{
		Items: []*Datasource{},
	}
}

type DatasourceSet struct {
	Total int           `json:"total"`
	Items []*Datasource `json:"items"`
}

func (s *DatasourceSet) Add(item *Datasource) {
	s.Items = append(s.Items, item)
}

type QueryDatasourceRequest struct {
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	Keywords   string `json:"kws"`
	Cate       int8   `json:"cate"`
}

func (req *QueryDatasourceRequest) GetPageSize() uint {
	return uint(req.PageSize)
}

func (req *QueryDatasourceRequest) OffSet() int64 {
	return int64((req.PageNumber - 1) * req.PageSize)
}

func NewQueryDatasourceHTTP(r *http.Request) *QueryDatasourceRequest {
	req := NewQueryDatasourceRequest()
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

	cate := qs.Get("cate")
	if cate != "" {
		if n, err := strconv.Atoi(cate); err == nil {
			req.Cate = int8(n)
		}
	}

	req.Keywords = qs.Get("kws")
	return req
}

func NewQueryDatasourceRequest() *QueryDatasourceRequest {
	return &QueryDatasourceRequest{
		PageSize:   10,
		PageNumber: 1,
	}
}
