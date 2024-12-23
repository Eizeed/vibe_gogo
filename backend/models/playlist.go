package models

import (
	"errors"
	"fmt"

	"github.com/Eizeed/vibe_gogo/db"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Playlist struct {  
    UUID            uuid.UUID       `json:"uuid" gorm:"primaryKey"`
    Title           string          `json:"title"`
    UserUUID        uuid.UUID       `json:"user_uuid"`
    IsPrivate       bool            `json:"is_private"`
    Tracklist       pq.Int64Array   `json:"tracklist" gorm:"type:bigint[]"`
}

type PlaylistModel struct {};

func (m *PlaylistModel) GetByUserUuid(uuid uuid.UUID) ([]Playlist, error) {
    db := db.GetDB();

    var playlists []Playlist;
    if err := db.Where("user_uuid = ?", uuid).Find(&playlists).Error; err != nil {
        return []Playlist{}, err
    }
    return playlists, nil
}

func (m *PlaylistModel) GetByUuid(uuid uuid.UUID) (Playlist, error) {
    db := db.GetDB();
    
    var playlist Playlist;

    if err := db.Where("uuid = ?", uuid).First(&playlist).Error; err != nil {
        return Playlist{}, err;
    }
    return playlist, nil
}

func (m *PlaylistModel) Create(userUuid uuid.UUID, uuid uuid.UUID, title string) (Playlist, error) {
    db := db.GetDB();

    newPlaylist := Playlist {
        UUID: uuid,
        Title: title,
        UserUUID: userUuid,
        IsPrivate: true,
        Tracklist: []int64{},
    };

    if err := db.Create(&newPlaylist).Error; err != nil{
        return Playlist{}, err
    }
    
    var playlist Playlist;
    if err := db.Where("uuid", newPlaylist.UUID).First(&playlist).Error; err != nil {
        return Playlist{}, err
    }

    return playlist, nil;
}

func (m *PlaylistModel) AddTrack(userUuid uuid.UUID, playlistUuid uuid.UUID, trackId int64) (Playlist, error) {
    db := db.GetDB();

    var playlist Playlist;

    if err := db.Where("uuid = ? AND user_uuid = ?", playlistUuid, userUuid).First(&playlist).Error; err != nil {
        return Playlist{}, err;
    }

    for _, v := range playlist.Tracklist {
        if v == trackId {
            return Playlist{}, errors.New("Track is already in playlist");
        }
    }

    
    playlist.Tracklist = append(playlist.Tracklist, trackId);
    fmt.Println(playlist.Tracklist)

    if err := db.Where("uuid = ?", playlist.UUID).Save(&playlist).Error; err != nil {
        return Playlist{}, err;
    }

    return playlist, nil
}

func (m *PlaylistModel) DeleteTrack(userUuid uuid.UUID, playlistUuid uuid.UUID, trackId int64) (Playlist, error) {
    db := db.GetDB();

    var playlist Playlist;

    if err := db.Where("uuid = ? AND user_uuid = ?", playlistUuid, userUuid).First(&playlist).Error; err != nil {
        return Playlist{}, err;
    }

    var newTracklist = []int64{};
    for _, t := range playlist.Tracklist {
        if t != trackId {
            newTracklist = append(newTracklist, t);
        }
    }

    if len(newTracklist) == len(playlist.Tracklist) {
        return Playlist{}, errors.New("Track is not found in this playlist")
    }

    playlist.Tracklist = newTracklist;

    if err := db.Save(&playlist).Error; err != nil {
        return Playlist{}, err;
    }

    return playlist, nil
}

func (m *PlaylistModel) Update() (string, error) {
    return "Update playlist", nil
}

func (m *PlaylistModel) ChangeVisibility(userUuid uuid.UUID, playlistUuid uuid.UUID) (Playlist, error) {
    db := db.GetDB();

    var playlist Playlist;
    
    if err := db.Where("uuid = ? AND user_uuid = ?", playlistUuid, userUuid).First(&playlist).Error; err != nil {
        return Playlist{}, err;
    }

    playlist.IsPrivate = !playlist.IsPrivate;

    if err := db.Save(&playlist).Error; err != nil {
        return Playlist{}, err;
    }

    return playlist, nil;
}

func (m *PlaylistModel) Delete() (string, error) {
    return "Playlist deleted", nil
}

















