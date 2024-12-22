package handlers

import (
	"net/http"

	"github.com/Eizeed/vibe_gogo/services"
	"github.com/gin-gonic/gin"
)

type PlaylistHandler struct {};

var playlistService = services.PlaylistService {};

func (h *PlaylistHandler) GetByUserUuid(c *gin.Context) {
    res, err := playlistService.GetByUserUuid();
    if err != nil {
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func (h *PlaylistHandler) GetByUuid(c *gin.Context) {
    res, err := playlistService.GetByUuid();
    if err != nil {
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func (h *PlaylistHandler) Create(c *gin.Context) {
    res, err := playlistService.Create();
    if err != nil {
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func (h *PlaylistHandler) AddTrack(c *gin.Context) {
    res, err := playlistService.AddTrack();
    if err != nil {
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func (h *PlaylistHandler) DeleteTrack(c *gin.Context) {
    res, err := playlistService.DeleteTrack();
    if err != nil {
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func (h *PlaylistHandler) Update(c *gin.Context) {
    res, err := playlistService.Update();
    if err != nil {
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func (h *PlaylistHandler) ChangeVisibility(c *gin.Context) {
    res, err := playlistService.ChangeVisibility();
    if err != nil {
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func (h *PlaylistHandler) Delete(c *gin.Context) {
    res, err := playlistService.Delete();
    if err != nil {
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}













