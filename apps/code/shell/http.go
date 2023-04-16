package shell

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
	"go-notebook/apps"
	"go-notebook/apps/code"
)

// 面向接口, 真正Service的实现, 在服务实例化的时候传递进行
// 也就是(CLI)  Start时候
var handler = &Handler{}

// 通过写一个实例类, 把内部的接口通过HTTP协议暴露出去
// 所以需要依赖内部接口的实现
// 该实体类, 会实现Gin的Http Handler
type Handler struct {
	svc ShellCodeService
}

func (h *Handler) query(c *gin.Context) {
	println(99)
	// 从http请求的query string 中获取参数
	req := NewQueryShellCodeHTTP(c.Request)

	println(111)

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
	h.svc = apps.GetImpl(h.Name()).(ShellCodeService)
}

// 完成了 Http Handler的注册
func (h *Handler) Registry(r gin.IRouter) {
	println(222)
	r.GET("/code/shellCodes", h.query)
}

func (h *Handler) Name() string {
	return code.AppName + "-" + ModelName
}

// 完成Http Handler注册
func init() {
	apps.RegistryGin(handler)
}
