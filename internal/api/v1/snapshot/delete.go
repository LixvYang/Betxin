package snapshot

import (
	"github.com/lixvyang/betxin/internal/model"

	v1 "github.com/lixvyang/betxin/internal/api/v1"

	"github.com/lixvyang/betxin/internal/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func DeleteSnapshot(c *gin.Context) {
	traceId := c.Param("traceId")
	code := model.DeleteMixinNetworkSnapshot(traceId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
