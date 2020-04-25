package controller

import (
	"awesomeProject/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

//追加した
var config oauth2.Config

//別サイトから取ってきた
// cookieの設定を行う
//func setCookies(w http.ResponseWriter, r *http.Request) {
//	cookie := &http.Cookie{
//		Name:  "hoge",
//		Value: "bar",
//	}
//	http.SetCookie(w, cookie)
//
//	fmt.Fprintf(w, "Cookieの設定ができたよ")
//}

//gin使うか
func Login(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, service.GetRedirectURL())
}

//コンテキスト使うのか
//func OAuth(w http.ResponseWriter, r *http.Request) {
//func OAuth(ctx context.Context) {

//code := ctx.Value("code")
func OAuth(c *gin.Context) {

	code := c.Query("code")
	//code := r.FormValue("code")
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
	//url := config.AuthCodeURL("test")
	//fmt.Println(url)

	//token, err := service.GetToken(code)
	//serviceでやりたいが一旦はここで
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		//log.Fatal(err)
		log.Fatal("controller//token取得失敗")
	}

	//追記した
	//client := config.Client(oauth2.NoContext, token) //httpクライアントを取得
	//log.Println(client)

	//if err != nil {
	//	//c.AbortWithError(http.StatusInternalServerError, err)
	//	//return
	//	log.Fatal("token取得失敗")
	//}

	////別サイト参考に作った
	//cookie := &http.Cookie{
	//	Name:  "spotify-token",
	//	Value: token.AccessToken,
	//}

	//log.Println(cookie)
	//http.SetCookie(w, cookie)
	//
	//fmt.Fprintf(w, "Cookieの設定ができたよ")

	//要修正か
	//c.SetCookie("spotify-token", token.AccessToken, 1000*60*60*24*7, "/", "https://musi-app.now.sh", false, false)
	c.SetCookie("spotify-token", token.AccessToken, 1000*60*60*24*7, "/", "https://localhost:8080", false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

//ginで
func GetPlayList(c *gin.Context) {
	//func GetPlayList(w http.ResponseWriter, r *http.Request) {

	//一旦はここでやる
	//c.SetCookie("spotify-token", token.AccessToken, 1000*60*60*24*7, "/", "https://localhost:8080", false, false)
	c.SetCookie("spotify-token", "BQAQ8IXL2MkfLMrNWK0Bs-b_jRdVIVfFMZvTSk_eCifwXedDodKrxch5hCRAZQXDyWdCa1HQv7F9RZJc2mw", 1000*60*60*24*7, "/", "http://localhost:8080", false, false)

	//gin使わない場合
	//cookie, err := r.Cookie("spotify-token")
	//log.Println(cookie)
	//if err != nil {
	//	log.Println("cookieが取得できない")
	//	log.Fatal("Cookie:", err)
	//}

	//gin使う場合
	cookie, _ := c.Cookie("spotify-token")

	if cookie == "" {
		//gin使わない場合
		//if cookie.Value == "" {
		//c.AbortWithStatus(http.StatusUnauthorized)
		//return
		log.Println("cookieが空白")
	}
	log.Println("cookie:", cookie)

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

	//gin使う場合
	//playlist, err := service.GetTracks(cookie)

	//gin使わない場合
	//playlist, err := service.GetTracks(cookie.Value)
	//log.Println("playlist:", playlist)

	//if err != nil {
	//	//c.AbortWithStatus(http.StatusInternalServerError)
	//	//return
	//	log.Fatal("プレイリスト取得失敗")
	//}

	//gin使う場合
	//c.JSON(http.StatusOK, playlist)

	//gin使わない場合
	//v, err := json.Marshal(playlist)
	//if err != nil {
	//	println(string(v))
	//}
	//w.Write(v)
}

func GetArtist(c *gin.Context) {

	//一旦はここでやる
	//c.SetCookie("spotify-token", token.AccessToken, 1000*60*60*24*7, "/", "https://localhost:8080", false, false)
	//c.SetCookie("spotify-token", "BQAQ8IXL2MkfLMrNWK0Bs-b_jRdVIVfFMZvTSk_eCifwXedDodKrxch5hCRAZQXDyWdCa1HQv7F9RZJc2mw", 1000*60*60*24*7, "/", "http://localhost:8080", false, false)
	c.SetCookie("spotify-token", "BQDCQ75jmVWWFQoJWv2lC6iNFZAN1KCvT_3eA7IxUJ0CGf9I8dVyLp5ysUrX29lqJQ9PiwYIL2pkOarfISA", 1000*60*60*24*7, "/", "http://localhost:8080", false, false)

	//gin使う場合
	cookie, _ := c.Cookie("spotify-token")

	if cookie == "" {
		//gin使わない場合
		//if cookie.Value == "" {

		c.AbortWithStatus(http.StatusUnauthorized)
		return
		//log.Println("cookieが空白")
	}
	log.Println("cookie:", cookie)

	//playlist, err := service.GetTracks(cookie, location)

	//gin使う場合
	//固定したIdでアーティスト取得
	//artist, err := service.GetMusicArtistId(cookie)

	//アーティスト検索
	artists, err := service.SearchMusicArtists(cookie)

	//gin使わない場合
	//playlist, err := service.GetTracks(cookie.Value)
	log.Println("artists:", artists)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
		log.Fatal("アーティスト取得失敗")
	}

	//gin使う場合
	c.JSON(http.StatusOK, artists)

	//gin使わない場合
	//v, err := json.Marshal(playlist)
	//if err != nil {
	//	println(string(v))
	//}
	//w.Write(v)
}
