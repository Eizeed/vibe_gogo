package services

import "github.com/Eizeed/vibe_gogo/models"

type PlaylistService struct {};

var playlistModel = models.PlaylistModel {};

func (s *PlaylistService) GetByUserUuid() (string, error) {
    return playlistModel.GetByUserUuid();
}

func (s *PlaylistService) GetByUuid() (string, error) {
    return playlistModel.GetByUuid();
}

func (s *PlaylistService) Create() (string, error) {
    return playlistModel.Create();
}

func (s *PlaylistService) AddTrack() (string, error) {
    return playlistModel.AddTrack();
}

func (s *PlaylistService) DeleteTrack() (string, error) {
    return playlistModel.DeleteTrack();
}

func (s *PlaylistService) Update() (string, error) {
    return playlistModel.Update()
}

func (s *PlaylistService) ChangeVisibility() (string, error) {
    return playlistModel.ChangeVisibility()
}

func (s *PlaylistService) Delete() (string, error) {
    return playlistModel.Delete();
}
