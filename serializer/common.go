package serializer

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status_Code int
	Data        interface{}
	Msg         string
	Error       string
}

func (r *Response) CtxSetResponse(c *gin.Context) {
	c.JSON(r.Status_Code, gin.H{"code": r.Status_Code, "msg": r.Msg, "data": r.Data, "error:": r.Error})
}
