package handlers

import (
	"net/http"
	"strconv"

	"github.com/Eizeed/vibe_gogo/forms"
	"github.com/Eizeed/vibe_gogo/models"
	"github.com/Eizeed/vibe_gogo/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PlaylistHandler struct {};

var playlistService = services.PlaylistService {};

func (h *PlaylistHandler) GetByUserUuid(c *gin.Context) {
    var tokenUuid uuid.UUID;
    cookie, err := c.Cookie("jwt_token");
    if err == nil {
        var jwt models.JWT;
        claims, err := jwt.DecodeToken(cookie);
        if err == nil {
            tokenUuid = claims.UserUUID;
        }
    }

    ownerUuid, err := uuid.Parse(c.Params.ByName("uuid"));
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user uuid"})
        return
    }
    

    res, err := playlistService.GetByUserUuid(tokenUuid, ownerUuid);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": "Got user's playlists", "playlists": res})
}

func (h *PlaylistHandler) GetByUuid(c *gin.Context) {
    var tokenUuid uuid.UUID;
    cookie, err := c.Cookie("jwt_token");
    if err == nil {
        var jwt models.JWT;
        claims, err := jwt.DecodeToken(cookie);
        if err == nil {
            tokenUuid = claims.UserUUID;
        }
    }

    plalistUuid, err := uuid.Parse(c.Params.ByName("uuid"));
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist uuid"})
        return
    }
    res, err := playlistService.GetByUuid(tokenUuid, plalistUuid);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    tracks, err := deezerService.FetchTracks(res.Tracklist);
    if err != nil { 
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


    c.IndentedJSON(http.StatusOK, gin.H{"message": "Got playlist by uuid", "playlist": res, "tracklist": tracks})
}

func (h *PlaylistHandler) Create(c *gin.Context) {
    token, exists := c.Get("token");
    if !exists {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    claims, ok := token.(models.Claims);
    if !ok {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unable to get claims from token"})
        return
    }

    var f forms.PlaylistCreateForm;

    err := c.ShouldBindBodyWithJSON(&f);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    res, err := playlistService.Create(claims.UserUUID, f.Title);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func (h *PlaylistHandler) AddTrack(c *gin.Context) {
    token, exists := c.Get("token");
    if !exists {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    claims, ok := token.(models.Claims);
    if !ok {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unable to get claims from token"})
        return
    }

    var body map[string]interface{};
    err := c.ShouldBindBodyWithJSON(&body);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to parse body"})
        return
    }
    trackStr, exists := body["track_id"].(string);
    if !exists {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "track_id field is required"})
        return
    }
    trackId, err := strconv.ParseInt(trackStr, 0, 64);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "unable to parse track_id"})
        return
    }
    playlistUuid, err := uuid.Parse(c.Params.ByName("uuid"));
    if err != nil { 
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid playlist uuid"})
        return
    }
    
    res, err := playlistService.AddTrack(claims.UserUUID, playlistUuid, trackId);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": "Added track to playlist", "playlist": res})
}

func (h *PlaylistHandler) DeleteTrack(c *gin.Context) {
    token, exists := c.Get("token");
    if !exists {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    claims, ok := token.(models.Claims);
    if !ok {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unable to get claims from token"})
        return
    }


    var body map[string]interface{};
    err := c.ShouldBindBodyWithJSON(&body);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to parse body"})
        return
    }

    trackStr, exists := body["track_id"].(string);
    if !exists {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "track_id field is required"})
        return
    }
    trackId, err := strconv.ParseInt(trackStr, 0, 64);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "unable to parse track_id"})
        return
    }
    playlistUuid, err := uuid.Parse(c.Params.ByName("uuid"));
    if err != nil { 
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid playlist uuid"})
        return
    }

    res, err := playlistService.DeleteTrack(claims.UserUUID, playlistUuid, trackId);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted track from playlist", "playlist": res})
}

func (h *PlaylistHandler) Update(c *gin.Context) {
    res, err := playlistService.Update();
    if err != nil {
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func (h *PlaylistHandler) ChangeVisibility(c *gin.Context) {
    token, exists := c.Get("token");
    if !exists {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    claims, ok := token.(models.Claims);
    if !ok {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unable to get claims from token"})
        return
    }

    uuid, err := uuid.Parse(c.Params.ByName("uuid"));
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid plalist uuid"})
        return
    }
    
    res, err := playlistService.ChangeVisibility(claims.UserUUID, uuid);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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













