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
func GetTracks(token string) (*model.Response, error) {
	// weather
	//w, err := owm.NewCurrent("F", "EN", openweatherAPIKey)
	//if err != nil {
	//	return nil, err
	//}
	//w.CurrentByCoordinates(&owm.Coordinates{
	//	Longitude: location.Longitude,
	//	Latitude:  location.Latitude,
	//})

	values := url.Values{}
	log.Println("values", values)
	//values.Add("q", w.Weather[0].Main)
	values.Add("type", "playlist")
	log.Println(values)
	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)
	req.URL.RawQuery = values.Encode()
	req.Header.Set("Authorization", "Bearer "+token)
	log.Println(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var playlist model.PlayList
	json.Unmarshal(b, &playlist)

	response := &model.Response{
		//CityName: w.Name,
		Playlist: playlist,
		//Weather:  w.Weather[0].Main,
	}

	return response, nil
}
