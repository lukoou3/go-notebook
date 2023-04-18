package datasource

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
	"go-notebook/apps"
)

var handler = &Handler{}

type Handler struct {
	svc DatasourceService
}

func (h *Handler) query(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := NewQueryDatasourceHTTP(c.Request)

	// 进行接口调用, 返回 肯定有成功或者失败
	set, err := h.svc.Query(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}

func (h *Handler) Config() {
	// 从IOC里面获取HostService的实例对象
	h.svc = apps.GetImpl(h.Name()).(DatasourceService)
}

// 完成了 Http Handler的注册
func (h *Handler) Registry(r gin.IRouter) {
	r.GET("/sql/datasource", h.query)
}

func (h *Handler) Name() string {
	return registerName()
}

// 完成Http Handler注册
func init() {
	apps.RegistryGin(handler)
}
