package services

import (
	"errors"

	"github.com/Eizeed/vibe_gogo/models"
	"github.com/google/uuid"
)

type PlaylistService struct {};

var playlistModel = models.PlaylistModel {};

func (s *PlaylistService) GetByUserUuid(uuid uuid.UUID, ownerUuid uuid.UUID) ([]models.Playlist, error) {
    res, err := playlistModel.GetByUserUuid(ownerUuid);
    if err != nil {
        return []models.Playlist{}, err;
    }

    if uuid == ownerUuid {
        return res, nil;
    }

    var filteredPl = []models.Playlist{};
    for _, pl := range res {
        if pl.IsPrivate {
            if pl.UserUUID == uuid {
                filteredPl = append(filteredPl, pl);
            }
        } else {
            filteredPl = append(filteredPl, pl)
        }
    }

    return filteredPl, nil;
}

func (s *PlaylistService) GetByUuid(userUuid uuid.UUID, playlistUuid uuid.UUID) (models.Playlist, error) {
    res, err := playlistModel.GetByUuid(playlistUuid);
    if err != nil {
        return models.Playlist{}, err;
    }

    if res.IsPrivate {
        if res.UserUUID == userUuid {
            return res, nil;
        } else {
            return models.Playlist{}, errors.New("Playlist not found");
        }
    } else {
        return res, nil
    }
}

func (s *PlaylistService) Create(userUuid uuid.UUID, title string) (models.Playlist, error) {
    uuid, err := uuid.NewV7();
    if err != nil {    
        return models.Playlist{}, err;
    }

    return playlistModel.Create(userUuid, uuid, title);
}

func (s *PlaylistService) AddTrack(userUuid uuid.UUID, playlistUuid uuid.UUID, trackId int64) (models.Playlist, error) {
    res, err := playlistModel.AddTrack(userUuid, playlistUuid, trackId);
    if err != nil {
        return models.Playlist{}, err;
    }

    return res, nil;
}

func (s *PlaylistService) DeleteTrack(userUuid uuid.UUID, playlistUuid uuid.UUID, trackId int64) (models.Playlist, error) {
    res, err := playlistModel.DeleteTrack(userUuid, playlistUuid, trackId);
    if err != nil {
        return models.Playlist{}, err;
    }

    return res, nil;
}

func (s *PlaylistService) Update() (string, error) {
    return playlistModel.Update()
}

func (s *PlaylistService) ChangeVisibility(userUuid uuid.UUID, playlistUuid uuid.UUID) (models.Playlist, error) {
    return playlistModel.ChangeVisibility(userUuid, playlistUuid)
}

func (s *PlaylistService) Delete() (string, error) {
    return playlistModel.Delete();
}















