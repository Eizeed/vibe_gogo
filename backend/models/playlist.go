package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Playlist struct {  
    UUID            uuid.UUID       `json:"uuid"`
    Title           string          `json:"title"`
    UserUUID        uuid.UUID       `json:"user_uuid"`
    IsPrivate       string          `json:"is_private"`
    Tracklist       pq.Int64Array   `json:"tracklist"`
}

type PlaylistModel struct {};

func (m *PlaylistModel) GetByUserUuid() (string, error) {
    return "playlists from user", nil
}

func (m *PlaylistModel) GetByUuid() (string, error) {
    return "playlist by uuid", nil
}

func (m *PlaylistModel) Create() (string, error) {
    return "Create playlist", nil
}

func (m *PlaylistModel) AddTrack() (string, error) {
    return "Track added", nil
}

func (m *PlaylistModel) DeleteTrack() (string, error) {
    return "Track deleted", nil
}

func (m *PlaylistModel) Update() (string, error) {
    return "Update playlist", nil
}

func (m *PlaylistModel) ChangeVisibility() (string, error) {
    return "Playlist visibility changed", nil
}

func (m *PlaylistModel) Delete() (string, error) {
    return "Playlist deleted", nil
}

















