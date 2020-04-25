package service

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var config oauth2.Config

func init() {
	err := godotenv.Load()
	if err != nil {
		//panic(err)
		log.Println("godotenv失敗")
		//log.Fatal(err)
	}

	//fmt.Println(os.Getenv("client_id"))
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
		//RedirectURL: "http://localhost:8080/oauth",
		//今回はリダイレクトしない
		RedirectURL: "urn:ietf:wg:oauth:2.0:oob",
		//Scopes:      []string{"playlist-modify", "user-read-private", "user-library-read"},
		Scopes: []string{},
	}
	fmt.Println(config)
	//fmt.Println(os.Getenv("client_secret"))
}

//以下よくわからない
func GetRedirectURL() string {
	// TODO: CSRF対策
	return config.AuthCodeURL("state")
}

//ControllerのアクションOAuthで呼び出される
func GetToken(code string) (*oauth2.Token, error) {
	//config.Exchangeとは
	//return config.Exchange(oauth2.NoContext, code)

	//他のサイト参考に書いた
	log.Println("code:", code)

	//アクセストークン取得
	tok, err := config.Exchange(oauth2.NoContext, code)

	if err != nil {
		log.Println("config.Exchange失敗")
		log.Fatal(err)
	}

	//log.Println(tok)
	//return tok
	return tok, nil
}
