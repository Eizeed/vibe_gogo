package handlers

import (
	"net/http"

	"github.com/Eizeed/vibe_gogo/models"
	"github.com/Eizeed/vibe_gogo/services"
	"github.com/gin-gonic/gin"
)

type DeezerHandler struct {};

var deezerService services.DeezerService;

func (h *DeezerHandler) Search(c *gin.Context) {
    var params models.SearchParams;
    err := c.ShouldBindBodyWithJSON(&params);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, err);
        return
    }

    res, err := deezerService.Search(params);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, err);
        return
    }

    c.IndentedJSON(http.StatusOK, res)
}

