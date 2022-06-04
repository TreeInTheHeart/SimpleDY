package json_response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
	"net/http"
)

type BaseResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func Response(c *gin.Context, code int32, msg string, data interface{}) {
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}

	marshal, _ = sjson.SetBytes(marshal, "status_code", code)
	marshal, _ = sjson.SetBytes(marshal, "status_msg", msg)
	c.Data(http.StatusOK, "application/json", marshal)
	//c.JSON(http.StatusOK, marshal)

}

func OK(c *gin.Context, msg string, data interface{}) {
	Response(c, 0, msg, data)
}

func Error(c *gin.Context, code int32, msg string) {
	Response(c, code, msg, nil)
}
