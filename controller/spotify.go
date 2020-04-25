package controller

import (
	"awesomeProject/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

//追加した
var config oauth2.Config

//別サイトから取ってきた
// cookieの設定を行う
func setCookies(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "hoge",
		Value: "bar",
	}
	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Cookieの設定ができたよ")
}

//gin使うか
//func Login(c *gin.Context) {
//	c.Redirect(http.StatusTemporaryRedirect, service.GetRedirectURL())
//}

func OAuth(w http.ResponseWriter, r *http.Request) {
	//func OAuth(c *gin.Context) {
	//code := c.Query("code")
	code := r.FormValue("code")
	fmt.Println(code)

	config = oauth2.Config{
		//取れてるはず
		ClientID: os.Getenv("client_id"),
		//取れてるはず
		ClientSecret: os.Getenv("client_secret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		},

		//RedirectURL: "https://musi-app.now.sh/oauth",
		RedirectURL: "http://localhost:8080/oauth",
		//今回はリダイレクトしない
		//RedirectURL: "urn:ietf:wg:oauth:2.0:oob",
		//Scopes:      []string{"playlist-modify", "user-read-private", "user-library-read"},
		Scopes: []string{},
	}

	//追加した
	url := config.AuthCodeURL("test")
	fmt.Println(url)

	//token, err := service.GetToken(code)
	//serviceでやりたいが一旦はここで
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		//log.Fatal(err)
		log.Fatal("token取得失敗")
	}

	client := config.Client(oauth2.NoContext, token) //httpクライアントを取得
	log.Println(client)

	//if err != nil {
	//	//c.AbortWithError(http.StatusInternalServerError, err)
	//	//return
	//	log.Fatal("token取得失敗")
	//}

	//別サイト参考に作った
	cookie := &http.Cookie{
		Name:  "spotify-token",
		Value: token.AccessToken,
	}

	log.Println(cookie)
	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Cookieの設定ができたよ")

	//要修正か
	//c.SetCookie("spotify-token", token.AccessToken, 1000*60*60*24*7, "/", "https://musi-app.now.sh", false, false)
	//c.Redirect(http.StatusTemporaryRedirect, "/")
}

//func GetPlayList(c *gin.Context) {
func GetPlayList(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("spotify-token")
	log.Println(cookie)
	if err != nil {
		log.Println("cookieが取得できない")
		log.Fatal("Cookie:", err)
	}
	//cookie, _ := c.Cookie("spotify-token")
	if cookie.Value == "" {
		//c.AbortWithStatus(http.StatusUnauthorized)
		//return
		log.Println("cookieが空白")
	}

	//ここら辺は不要
	//	lat, err := strconv.ParseFloat(c.Query("latitude"), 64)
	//	if err != nil {
	//		c.AbortWithStatus(http.StatusBadRequest)
	//		return
	//	}
	//
	//	lng, err := strconv.ParseFloat(c.Query("longitude"), 64)
	//	if err != nil {
	//		c.AbortWithStatus(http.StatusBadRequest)
	//		return
	//	}

	//不要
	//location := model.GeoLocation{
	//	Longitude: lng,
	//	Latitude:  lat,
	//}

	//playlist, err := service.GetTracks(cookie, location)
	//playlist, err := service.GetTracks(cookie)
	playlist, err := service.GetTracks(cookie.Value)
	log.Println(playlist)
	if err != nil {
		//c.AbortWithStatus(http.StatusInternalServerError)
		//return
		log.Fatal("プレイリスト取得失敗")
	}

	//c.JSON(http.StatusOK, playlist)
	v, err := json.Marshal(playlist)
	if err != nil {
		println(string(v))
	}
	w.Write(v)
}
