package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QuestionHandler(ctx *gin.Context) {
	questionID := ctx.Query("qid")
	rPath := fmt.Sprintf("problem-%s.html", questionID)
	ctx.HTML(http.StatusOK, rPath, nil)
}
