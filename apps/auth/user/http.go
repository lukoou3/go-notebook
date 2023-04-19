package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
	"go-notebook/apps"
)

var handler = &Handler{}

type Handler struct {
}

func (h *Handler) query(c *gin.Context) {
	param := make(map[string]interface{})
	c.BindJSON(&param)
	fmt.Println(param)
	data := map[string]string{"token": "token"}
	/*respBytes, _ := json.Marshal(data)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(respBytes)*/
	response.Success(c.Writer, data)
}

func (h *Handler) Config() {
}

// 完成了 Http Handler的注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/auth/user/login", h.query)
}

func (h *Handler) Name() string {
	return registerName()
}

// 完成Http Handler注册
func init() {
	apps.RegistryGin(handler)
}
