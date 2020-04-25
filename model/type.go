package model

//type Response struct {
//	//CityName string   `json:"city_name"`
//	Playlist PlayList `json:"playlist"`
//	//Weather  string   `json:"weather"`
//}

type Response struct {
	//Artist Artist `json:"artist"`
	Artists Artists `json:"artists"`
}

//type GeoLocation struct {
//	Latitude  float64 `json:"latitude"`
//	Longitude float64 `json:"longitude"`
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
		Items []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Followers struct {
				Href  interface{} `json:"href"`
				Total int         `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Href   string   `json:"href"`
			ID     string   `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name       string `json:"name"`
			Popularity int    `json:"popularity"`
			Type       string `json:"type"`
			URI        string `json:"uri"`
		} `json:"items"`
		Limit    int         `json:"limit"`
		Next     string      `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"artists"`
}

//type Artists struct {
//	artists struct {
//		Href string
//		//Items    *[]Artist
//		Items    []Artist
//		Limit    int
//		Next     string
//		Offset   int
//		Previous string
//		Total    int
//		//"limit": 20,
//		//"next": null,
//		//"offset": 0,
//		//"previous": null,
//		//"total": 1
//	}
//
//	Href string
//	//Items    *[]Artist
//	Items    []Artist
//	Limit    int
//	Next     string
//	Offset   int
//	Previous string
//	Total    int
//	//"limit": 20,
//	//"next": null,
//	//"offset": 0,
//	//"previous": null,
//	//"total": 1
//}

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

//"artists": {
//	"href": "https://api.spotify.com/v1/search?query=tania+bowra&offset=0&limit=20&type=artist",
//	"items": [ {
//	"external_urls": {
//		"spotify": "https://open.spotify.com/artist/08td7MxkoHQkXnWAYD8d6Q"
//	},
//	"genres": [ ],
//	"href": "https://api.spotify.com/v1/artists/08td7MxkoHQkXnWAYD8d6Q",
//	"id": "08td7MxkoHQkXnWAYD8d6Q",
//	"images": [ {
//		"height": 640,
//		"url": "https://i.scdn.co/image/f2798ddab0c7b76dc2d270b65c4f67ddef7f6718",
//		"width": 640
//	}, {
//		"height": 300,
//		"url": "https://i.scdn.co/image/b414091165ea0f4172089c2fc67bb35aa37cfc55",
//		"width": 300
//	}, {
//		"height": 64,
//		"url": "https://i.scdn.co/image/8522fc78be4bf4e83fea8e67bb742e7d3dfe21b4",
//		"width": 64
//	} ],
//	"name": "Tania Bowra",
//	"popularity": 0,
//	"type": "artist",
//	"uri": "spotify:artist:08td7MxkoHQkXnWAYD8d6Q"
//} ],
//"limit": 20,
//"next": null,
//"offset": 0,
//"previous": null,
//"total": 1
//}
