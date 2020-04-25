package model

//type Response struct {
//	//CityName string   `json:"city_name"`
//	Playlist PlayList `json:"playlist"`
//	//Weather  string   `json:"weather"`
//}

type Response struct {
	Artists Artists `json:"artists"`
}

//type Response struct {
//	Tracks Tracks `json:"tracks"`
//}

type PlayList struct {
	Playlists struct {
		Href  string `json:"href"`
		Items []struct {
			Collaborative bool `json:"collaborative"`
			ExternalUrls  struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name  string `json:"name"`
			Owner struct {
				DisplayName  interface{} `json:"display_name"`
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"owner"`
			PrimaryColor interface{} `json:"primary_color"`
			Public       interface{} `json:"public"`
			SnapshotID   string      `json:"snapshot_id"`
			Tracks       struct {
				Href  string `json:"href"`
				Total int    `json:"total"`
			} `json:"tracks"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"items"`
		Limit    int         `json:"limit"`
		Next     string      `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"playlists"`
}

type Artists struct {
	Artists struct {
		Href  string `json:"href"`
		Items []Artist

		//Items []struct {
		//	ExternalUrls struct {
		//		Spotify string `json:"spotify"`
		//	} `json:"external_urls"`
		//	Followers struct {
		//		Href  interface{} `json:"href"`
		//		Total int         `json:"total"`
		//	} `json:"followers"`
		//	Genres []string `json:"genres"`
		//	Href   string   `json:"href"`
		//	ID     string   `json:"id"`
		//	Images []struct {
		//		Height int    `json:"height"`
		//		URL    string `json:"url"`
		//		Width  int    `json:"width"`
		//	} `json:"images"`
		//	Name       string `json:"name"`
		//	Popularity int    `json:"popularity"`
		//	Type       string `json:"type"`
		//	URI        string `json:"uri"`
		//} `json:"items"`
		Limit    int         `json:"limit"`
		Next     string      `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"artists"`
}

type Artist struct {
	External_urls struct {
		Spotify string
	}
	//"external_urls" : {
	//	"spotify" : "https://open.spotify.com/artist/1O8CSXsPwEqxcoBE360PPO"
	//},
	Followers struct {
		Href  string
		Total int
	}
	//"followers" : {
	//	"href" : null,
	//	"total" : 374781
	//},
	Genres []string
	//"genres" : [ "finnish metal", "finnish power metal", "melodic metal", "metal", "neo classical metal", "power metal", "progressive metal", "speed metal" ],
	Href string
	//"href" : "https://api.spotify.com/v1/artists/1O8CSXsPwEqxcoBE360PPO",
	Id string
	//"id" : "1O8CSXsPwEqxcoBE360PPO",
	Images []struct {
		Height int
		Url    string
		Width  int
	}
	//"images" : [ {
	//	"height" : 640,
	//	"url" : "https://i.scdn.co/image/8285d6793a22dab9982fafc1a7cb43a41ff6fae3",
	//	"width" : 640
	//}, {
	//	"height" : 320,
	//	"url" : "https://i.scdn.co/image/ff46b7ae4f095874bddc7a3738c235d5a49c7596",
	//	"width" : 320
	//}, {
	//	"height" : 160,
	//	"url" : "https://i.scdn.co/image/aa243f0e1686260e63df633fc6762f34de02adee",
	//	"width" : 160
	//} ],
	Name string
	//"name" : "Stratovarius",
	Popularity int
	//"popularity" : 57,
	Type string
	//"type" : "artist",
	Uri string
	//"uri" : "spotify:artist:1O8CSXsPwEqxcoBE360PPO"
}

type Tracks struct {
	Tracks struct {
		Href  string `json:"href"`
		Items []struct {
			Album struct {
				AlbumType string `json:"album_type"`
				Artists   []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"artists"`
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href   string `json:"href"`
				ID     string `json:"id"`
				Images []struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name                 string `json:"name"`
				ReleaseDate          string `json:"release_date"`
				ReleaseDatePrecision string `json:"release_date_precision"`
				TotalTracks          int    `json:"total_tracks"`
				Type                 string `json:"type"`
				URI                  string `json:"uri"`
			} `json:"album"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			DiscNumber  int  `json:"disc_number"`
			DurationMs  int  `json:"duration_ms"`
			Explicit    bool `json:"explicit"`
			ExternalIds struct {
				Isrc string `json:"isrc"`
			} `json:"external_ids"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href        string `json:"href"`
			ID          string `json:"id"`
			IsLocal     bool   `json:"is_local"`
			IsPlayable  bool   `json:"is_playable"`
			Name        string `json:"name"`
			Popularity  int    `json:"popularity"`
			PreviewURL  string `json:"preview_url"`
			TrackNumber int    `json:"track_number"`
			Type        string `json:"type"`
			URI         string `json:"uri"`
		} `json:"items"`
		Limit    int         `json:"limit"`
		Next     string      `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"tracks"`
}
