package service

import (
	"awesomeProject/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	//owm "github.com/briandowns/openweathermap"
	//"github.com/konojunya/musi/model"
)

//この記述自体は、位置情報→天気→Spoitfyのよう
//https://note.com/konojunya/n/ncf4b3c7c4d51を参考に

//func GetTracks(token string, location model.GeoLocation) (*model.Response, error) {

//func GetTracks(token string) (*model.Response, error) {
//	// weather
//	//w, err := owm.NewCurrent("F", "EN", openweatherAPIKey)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//w.CurrentByCoordinates(&owm.Coordinates{
//	//	Longitude: location.Longitude,
//	//	Latitude:  location.Latitude,
//	//})
//
//	values := url.Values{}
//	//log.Println("values", values)
//	//values.Add("q", w.Weather[0].Main)
//	values.Add("type", "playlist")
//	log.Println(values)
//
//	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)
//	req.URL.RawQuery = values.Encode()
//	req.Header.Set("Authorization", "Bearer "+token)
//
//	client := &http.Client{}
//
//	log.Println("req:", req)
//
//	log.Println("client:", client)
//
//	resp, err := client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	log.Println("resp:", resp)
//
//	defer resp.Body.Close()
//
//	b, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println("err:", err)
//		return nil, err
//	}
//	log.Println("b:", b)
//
//	var playlist model.PlayList
//
//	json.Unmarshal(b, &playlist)
//
//	response := &model.Response{
//		//CityName: w.Name,
//		Playlist: playlist,
//		//Weather:  w.Weather[0].Main,
//	}
//
//	log.Println("response:", response)
//
//	return response, nil
//}

//func GetMusicArtistId(token string) (*model.Response, error) {
//
//	values := url.Values{}
//	//log.Println("values", values)
//	//values.Add("q", w.Weather[0].Main)
//
//	//とりあえず手動で
//	//アーティストのuuid?
//	//values.Add("q", "1O8CSXsPwEqxcoBE360PPO")
//
//	//values.Add("type", "artist")
//	log.Println("values:", values)
//
//	//req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)
//	//req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/artists", nil)
//	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/artists/1O8CSXsPwEqxcoBE360PPO", nil)
//	req.URL.RawQuery = values.Encode()
//	req.Header.Set("Authorization", "Bearer "+token)
//
//	client := &http.Client{}
//	log.Println("req:", req)
//	log.Println("client:", client)
//
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println("client.Do失敗")
//		return nil, err
//	}
//	log.Println("resp.Body:", resp.Body)
//
//	defer resp.Body.Close()
//
//	b, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println("err:", err)
//		return nil, err
//	}
//	log.Println("b:", b)
//
//	var artist model.Artist
//
//	json.Unmarshal(b, &artist)
//
//	//log.Println("b:", b)
//
//	response := &model.Response{
//		//CityName: w.Name,
//		Artist: artist,
//		//Weather:  w.Weather[0].Main,
//	}
//
//	log.Println("response:", response)
//
//	return response, nil
//}

func SearchMusicArtists(token string) (*model.Response, error) {

	values := url.Values{}

	values.Add("type", "artist")
	values.Add("q", "Radiohead")

	log.Println("values:", values)

	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)
	req.URL.RawQuery = values.Encode()
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	log.Println("req:", req)
	//log.Println("client:", client)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do失敗")
		return nil, err
	}
	log.Println("resp:", resp)

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	//log.Println("b:", b)

	var artists model.Artists
	log.Print("model.Artists:", artists)

	json.Unmarshal(b, &artists)

	response := &model.Response{
		//CityName: w.Name,
		Artists: artists,
		//Weather:  w.Weather[0].Main,
	}

	log.Println("response:", response)

	return response, nil
}
