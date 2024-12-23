package models

import ()

type SearchData struct {
    Data        []SearchTrack   `json:"data"`
    Total       uint            `json:"total"`
    Next        string          `json:"next"`
    Prev        string          `json:"prev"`
}

type SearchParams struct {
    Artist      string      `json:"artist"`
    Track       string      `json:"track"`
    Album       string      `json:"album"`
    Strict      bool        `json:"strict"`
}

type SearchTrack struct {
    Id                          int64       `json:"id"`
    Readable                    bool        `json:"readable"`
    Title                       string      `json:"title"`
    TitleShort                  string      `json:"title_short"`
    TitleVersion                string      `json:"title_version"`
    Link                        string      `json:"link"`
    Duration                    uint16      `json:"duration"`
    Rank                        uint32      `json:"rank"`
    ExplicitLyrics              bool        `json:"explicit_lyrics"`
    ExplicitContentLyrics       uint8       `json:"explicit_content_lyrics"`
    ExplicitContentCover        uint8       `json:"explicit_content_cover"`
    Preview                     string      `json:"preview"`
    Artist                      Artist      `json:"artist"`
    Album                       Album       `json:"album"`
    Type                        string      `json:"type"`
}

type Artist struct {
    Id                          int64       `json:"id"`
    Name                        string      `json:"name"`
    Link                        string      `json:"link"`
    Picture                     string      `json:"picture"`
    PictureSmall                string      `json:"picture_small"`
    PictureMedium               string      `json:"picture_medium"`
    PictureBig                  string      `json:"picture_big"`
    PictureXl                   string      `json:"picture_xl"`
    Type                        string      `json:"type"`
}

type Album struct {
    Id                          int64       `json:"id"`
    Title                       string      `json:"title"`
    Cover                       string      `json:"cover"`
    CoverSmall                  string      `json:"cover_small"`
    CoverMedium                 string      `json:"cover_medium"`
    CoverBig                    string      `json:"cover_big"`
    CoverXl                     string      `json:"cover_xl"`
    Type                        string      `json:"type"`
}

// It's the same as SearchRes
// But got ReleaseDate field in it
// Use this one when searching tracks by id
type Track struct {
    Id                          int64       `json:"id"`
    Readable                    bool        `json:"readable"`
    Title                       string      `json:"title"`
    TitleShort                  string      `json:"title_short"`
    TitleVersion                string      `json:"title_version"`
    Link                        string      `json:"link"`
    Duration                    uint16      `json:"duration"`
    Rank                        uint32      `json:"rank"`
    ReleaseDate                 string      `json:"release_date"`
    ExplicitLyrics              bool        `json:"explicit_lyrics"`
    ExplicitContentLyrics       uint8       `json:"explicit_content_lyrics"`
    ExplicitContentCover        uint8       `json:"explicit_content_cover"`
    Preview                     string      `json:"preview"`
    Artist                      Artist      `json:"artist"`
    Album                       Album       `json:"album"`
    Type                        string      `json:"type"`
}



















