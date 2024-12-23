package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Eizeed/vibe_gogo/models"
)

type DeezerService struct {};

func (s *DeezerService) Search(searchParams models.SearchParams) ([]models.SearchTrack, error) {
    searchStr := "https://api.deezer.com/search?q=";

    var data []models.SearchTrack;

  //  hasArtist := false;
    if searchParams.Artist != "" {
    //    hasArtist = true;
        escaped := url.QueryEscape(searchParams.Artist);
        searchStr = fmt.Sprintf(`%sartist:"%s"`, searchStr, escaped);
    }
    if searchParams.Track != "" {
        escaped := url.QueryEscape(searchParams.Track);
        searchStr = fmt.Sprintf(`%strack:"%s"`, searchStr, escaped);
    }
    if searchParams.Album != "" {
        escaped := url.QueryEscape(searchParams.Album);
        searchStr = fmt.Sprintf(`%salbum:"%s"`, searchStr, escaped);
    }
    if searchParams.Strict {
        searchStr = fmt.Sprintf(`%s&strict=on`, searchStr);
    }
    
    client := &http.Client {};
    for {
        req, err := http.NewRequest("GET", searchStr, nil);
        if err != nil {
            return []models.SearchTrack{}, err;
        }

        res, err := client.Do(req);
        if err != nil {
            return []models.SearchTrack{}, err;
        }
        defer res.Body.Close();

        body, err := io.ReadAll(res.Body);
        if err != nil {
            return []models.SearchTrack{}, err;
        }

        var sd models.SearchData;

        err = json.Unmarshal(body, &sd);
        if err != nil {
            return []models.SearchTrack{}, err;
        }

        data = append(data, sd.Data...)

        fmt.Println(searchStr)

        if sd.Next != "" {
            if sd.Next == searchStr {
                break;
            } else {
                searchStr = sd.Next
            }
        } else {
            break;
        }
    }


    return data, nil
}

func (s *DeezerService) FetchTracks(trackIds []int64) ([]models.Track, error) {
    searchStr := "https://api.deezer.com/track/";
    client := &http.Client {};

    var data []models.Track;
    for _, v := range trackIds {
        searchStr = fmt.Sprintf("%s%d", searchStr, v);
        req, err := http.NewRequest("GET", searchStr, nil);
        if err != nil {
            return []models.Track{}, err;
        }

        res, err := client.Do(req);
        if err != nil {
            return []models.Track{}, err;
        }
        defer res.Body.Close();

        body, err := io.ReadAll(res.Body);
        if err != nil {
            return []models.Track{}, err;
        }

        var track models.Track;
        err = json.Unmarshal(body, &track);
        if err != nil {
            return []models.Track{}, err;
        }

        data = append(data, track)
    }

    return data, nil;
}


















