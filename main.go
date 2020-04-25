package main

import (
	"awesomeProject/config"
	"awesomeProject/controller"

	//"time"

	//"awesomeProject/mylib"
	//"awesomeProject/router"
	"awesomeProject/utils"
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"

	//"github.com/joho/godotenv"
	"log"
	"net/http"
	//"os"
	//"github.com/konojunya/musi/router"
	//"awesomeProject/router"
)

type User struct {
	//主キーはidか
	gorm.Model
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
	//Nicknameはなぜかエラー出た
	//Nicknames []string `json:"nicknames"`
	//!作成されていない。外部キーを主キーのCredit CardのIDではなく、UserNameにしたい
	//CreditCard CreditCard `gorm:"foreignkey:UserName"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

type Song struct {
	//主キーはidか
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

//Belongs To
//type Profile struct {
//	gorm.Model
//	//外部キー設定されていないか
//	UserID int
//	//作成されていない
//	User User
//	Name string `json:"name"`
//	Age  int    `json:"age,omitempty"`
//	//Nicknames []string `json:"nicknames"`
//}

//Has One(UserがCreditCardを1つ持っている例)
//type CreditCard struct {
//	//主キーはid
//	gorm.Model
//	Number string
//	//符号なし整数 外部キーできていないっぽい
//	UserID   uint
//	UserName string
//}

type Event struct {
	gorm.Model
	Name string `json:name`
}

func gormConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/laravel6?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("接続成功")
	return db
}

// レスポンスにエラーを突っ込んで、返却するメソッド
func errorInResponse(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
	return
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	user := User{}
	error := Error{}

	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println("signup中")
	fmt.Println(email)
	fmt.Println(password)

	//使えるようにしたい
	//fmt.Println(r.Body)

	//一旦断念するが後でしたいJson.NewEncoder, json.NewDecoder で、エンコード(構造体から文字列)、デコード(文字列から構造体)の処理を行なっている。
	//json.NewDecoder(r.Body).Decode(&user)

	if email == "" {
		//if user.Email == "" {
		error.Message = "Emailは必須です。"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	if password == "" {
		//if user.Password == "" {
		error.Message = "パスワードは必須です。"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	fmt.Println(user)

	// dump も出せる
	fmt.Println("---------------------")
	spew.Dump(user)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	//hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("パスワード: ", password)
	//fmt.Println("パスワード: ", user.Password)
	fmt.Println("ハッシュ化されたパスワード", hash)

	user.Email = email
	user.Password = string(hash)
	//ほんとはuserに入れて処理していきたい
	password = string(hash)
	//fmt.Println("コンバート後のパスワード: ", user.Password)
	fmt.Println("コンバート後のパスワード: ", password)

	db := gormConnect()
	defer db.Close()
	db.Create(&User{Email: email, Password: password})
	//TOD　エラーハンドリング
	//err = db.Create(&User{Email: email, Password: password})
	//if err != nil {
	//	error.Message = "サーバーエラー"
	//	errorInResponse(w, http.StatusInternalServerError, error)
	//	return
	//}
	// DB に登録できたらパスワードをからにしておく
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	// 使えなかった JSON 形式で結果を返却
	//responseByJSON(w, user)

	v, err := json.Marshal(user)
	if err != nil {
		println(string(v))
	}
	w.Write(v)
}

func createToken(user User) (string, error) {
	//func createToken(w http.ResponseWriter, r *http.Request) (string, error) {
	var err error

	secret := "secret"

	//うまく取得できないので、自分で追記
	//email := r.FormValue("email")

	// Token を作成
	// jwt -> JSON Web Token - JSON をセキュアにやり取りするための仕様
	// jwtの構造 -> {Base64 encoded Header}.{Base64 encoded Payload}.{Signature}
	// HS254 -> 証明生成用(https://ja.wikipedia.org/wiki/JSON_Web_Token)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		//"email": email,
		"iss": "__init__", // JWT の発行者が入る(文字列(__init__)は任意)
	})

	//fmt.Println("user.Email:", email)

	//Dumpを吐く
	spew.Dump(token)

	tokenString, err := token.SignedString([]byte(secret))

	fmt.Println("-----------------------------")
	fmt.Println("tokenString:", tokenString)

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//userもpasswordも2つずつ用意するのはスマートではないか？
	//w.Write([]byte("successfully called login"))
	user := User{}
	//スマートでないので直したい
	user2 := User{}
	error := Error{}
	jwt := JWT{}
	//json.NewDecoder(r.Body).Decode(&user)

	//自分で追記した
	email := r.FormValue("email")
	password := r.FormValue("password")
	//user.Email = email
	//user.Password = password

	//if user.Email == "" {
	if email == "" {
		error.Message = "Email は必須です。"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	//if user.Password == "" {
	if password == "" {
		error.Message = "パスワードは、必須です。"
		errorInResponse(w, http.StatusBadRequest, error)
	}

	user.Email = email
	user.Password = password

	//password = user.Password

	//golang5が返ってくる
	fmt.Println("request_password:", password)

	db := gormConnect()
	defer db.Close()

	//これでうまくいくのはよくわからない
	row := db.Where("email = ?", user.Email).Find(&user2)

	//あくまでエラーハンドリング 要修正か
	_, err := json.Marshal(row)
	if err != nil {
		error.Message = "該当するアカウントが見つかりません。"
		errorInResponse(w, http.StatusUnauthorized, error)
		return
	}

	//signup4@example.comのユーザー情報
	fmt.Println("requestUser", user)
	fmt.Println("databaseUser", user2)
	password2 := user2.Password
	//データベース上のパスワード
	fmt.Println("database_password", password2)

	//fmt.Println(row)
	//修正する hash化？
	//err := row.Scan(&user.id, &user.Email, &user.Password)
	//err := row.Scan(&user.Email, &user.Password)
	//err := row.Scan(&user.Password)

	fmt.Println("databasePasswordBeforeHash", password2)

	// (リクエストパラメータの)パスワードのハッシュ化 ハッシュ化の必要あるのか
	//hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//if err != nil {
	//	return
	//}
	//hasedPassword := string(hash)
	//fmt.Println("hasedPassword:", hasedPassword)

	//if err != nil {
	//修正したい
	//if err == sql.ErrNoRows { // https://golang.org/pkg/database/sql/#pkg-variables
	//	error.Message = "ユーザが存在しません。"
	//	errorInResponse(w, http.StatusBadRequest, error)
	//} else {
	//	log.Fatal(err)
	//}
	//log.Fatal(err)
	//}
	//hasedPassword := user.Password
	//fmt.Println("hasedPassword:", hasedPassword)
	//password2はすでにハッシュ化されているのか
	err = bcrypt.CompareHashAndPassword([]byte(password2), []byte(password))

	if err != nil {
		//リクエストパスワードをなんとかする必要あるか
		//if password != password2 {
		error.Message = "無効なパスワードです。"
		errorInResponse(w, http.StatusUnauthorized, error)
		return
	}

	token, err := createToken(user)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	jwt.Token = token

	//responseByJSON(w, jwt)

	v2, err := json.Marshal(jwt)
	if err != nil {
		println(string(v2))
	}
	w.Write(v2)
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	//event := Event{
	//	Name: "example",
	//}
	//クエリパラメータの取得
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println("logging!")

	db := gormConnect()
	defer db.Close()

	eventEx := Event{}
	test1 := db.Where("id = ?", id).Find(&eventEx)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//Encodeメソッドでは、引数の構造体をマッピングします。
	//json.NewEncoder(w).Encode(event)
	//あくまでエラーハンドリングか
	v1, err := json.Marshal(test1)
	if err != nil {
		println(string(v1))
	}
	//v1, _ := json.Marshal(v)
	//fmt.Println(v.Type())
	//println(string(v1))
	//json.NewEncoder(w).Encode(v1)
	//w.Write(v1)
	//変数eventExをJSONにしたものをレスポンスに使う
	v, err := json.Marshal(eventEx)
	if err != nil {
		println(string(v))
	}
	w.Write(v)
	//w.WriteJson(v1)
}

func AllEventsHandler(w http.ResponseWriter, r *http.Request) {
	db := gormConnect()
	defer db.Close()
	allEvents := []Event{}

	db.Find(&allEvents)
	v2, _ := json.Marshal(allEvents)
	w.Write(v2)
}

//func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
//	//フォームの値を取得
//	name := r.FormValue("Name")
//	log.Println(name)
//
//	db := gormConnect()
//	defer db.Close()
//	//eventEx := Event{}
//
//	db.Create(&Event{Name: name})
//	//db.Model(&eventEx).Where("id = ?", id).Update("name", name)
//}

func CreateSongHandler(w http.ResponseWriter, r *http.Request) {
	//フォームの値を取得
	name := r.FormValue("name")
	email := r.FormValue("email")

	db := gormConnect()
	defer db.Close()
	//eventEx := Event{}
	db.Create(&Song{Name: name, Email: email})
}

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	//クエリパラメータの取得
	vars := mux.Vars(r)
	id := vars["id"]

	//フォームの値を取得
	name := r.FormValue("Name")
	log.Println(name)

	//UPDATE成功
	db := gormConnect()
	defer db.Close()
	eventEx := Event{}

	db.Model(&eventEx).Where("id = ?", id).Update("name", name)
}

//func LoggingSettings(logFile string){
//	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	multiLogFile := io.Multi￿Writer(os.Stdout, logfile)
//	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
//	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
//	log.SetOutput(multiLogFile)
//}

func main() {
	db := gormConnect()
	defer db.Close()
	db.Create(&User{Name: "test", Email: "test@example.com", Age: 20, Password: "golang"})

	//まだ試してない
	//r := router.GetRouter()
	//r.Run(":3000")

	//router.goは一旦コメントアウト
	r := mux.NewRouter()

	//r.HandleFunc("/", controller.Index).Methods("GET")
	//r.HandleFunc("/login", controller.Login).Methods("GET")
	r.HandleFunc("/oauth", controller.OAuth).Methods("GET")
	r.HandleFunc("/api/tracks", controller.GetPlayList).Methods("GET")

	r.HandleFunc("/song", CreateSongHandler).Methods("POST")
	r.HandleFunc("/signup", SignUpHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")

	//追加した
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

	//gorilla/muxを使ってみる
	//rでいいのでは
	//router := mux.NewRouter()
	//r := mux.NewRouter()
	//s := r.PathPrefix("/api").Subrouter()
	//r.PathPrefix("/api")
	//Subrouterメソッドでドメイン名設定を1回だけして、それを使い回せます。
	//s := r.Host("www.example.com").Subrouter()
	//httpかhttpsかを指定できる
	//r.Schemes("https")
	//r.Host("http://localhost:8080")
	//https://akirachiku.com/post/2017-04-08-go-net-http-api-server-4/には以下のようにあったが、.Handler以降がよくわからない
	//r.Methods("GET").Path("/hello/{name}").Handler(chain.Then(AppHandler{h: app.GreetingWithName}))

	//一旦コメントアウト
	//http.HandleFunc("/event", controller.Index)
	//http.HandleFunc("/login", controller.Login)
	//http.HandleFunc("/oauth", controller.OAuth)
	//http.HandleFunc("/api/tracks", controller.GetPlayList)

	//if err := http.ListenAndServe(":8080", r); err != nil {
	////if err := http.ListenAndServe(":8080", router); err != nil {
	//	log.Fatal(err)
	//}

	//こういう書き方もできるよう
	//r.HandleFunc("/products", ProductsHandler).
	//Host("www.example.com").
	//Methods("GET").
	//Schemes("http")

	//godotenv使用成功
	//err := godotenv.Load()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(os.Getenv("client_id"))
	//fmt.Println(os.Getenv("client_secret"))

	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	//設定ファイルの読み込みテスト
	//fmt.Printf("%T %v\n", config.Config.Port, config.Config.Port)
	//fmt.Printf("%T %v\n", config.Config.DbName, config.Config.DbName)
	//fmt.Println(config.Config.SQLDriver)
	//fmt.Println(config.Config.ApiSecret)
	//fmt.Println(config.Config.ApiKey)

	//math.goのテスト実行のために 成功
	//s := []int{1, 2, 3, 4, 5}
	//fmt.Println(mylib.Average(s))

	//log.Println("logging!")
	//log.Printf("%T %v", "test", "test")
	//
	//log.Fatalf("%T %v", "test", "test")
	////ここでコード終了してしまう
	//
	//log.Fatalln("error!!")
	////ここでコード終了してしまう
	//
	//log.Println("ok!!")

	//db := gormConnect()
	//defer db.Close()

	//テーブル作成
	//if !db.HasTable(&User{}) {
	//	db.CreateTable(&User{})
	//}
	//db.HasTable(&User{})
	//db.CreateTable(&User{})
	//db.CreateTable(&Profile{})
	//db.CreateTable(&CreditCard{})

	//テーブル削除
	//テーブルが存在する場合だけ削除
	//if db.HasTable(&User{}) {
	//	db.DropTable(&User{})
	//}
	//db.DropTable(&User{})
	//db.DropTable(&Profile{})
	//db.DropTable(&CreditCard{})

	//SELECT文(単一)
	//構造体のインスタンス化
	//eventEx := Event{}
	//IDの指定
	//eventEx.Id = 1
	//nameの指定
	//eventEx.Name = "test"
	// 指定したIDを元にレコードを１つ引っ張ってくる
	//test1 := db.First(&eventEx)
	//test1 := db.Where("id = ?", 5).Find(&eventEx)

	//v, _ := json.Marshal(test1)
	//fmt.Println(string(v))
	//http.HandleFunc("/event", EventHandler)
	//http.HandleFunc("/events", AllEventsHandler)
	//http.HandleFunc("/update", UpdateEventHandler)
	//r := mux.NewRouter()
	//r.HandleFunc("/event/{id}", EventHandler).Methods("GET")
	//r.HandleFunc("/event", CreateEventHandler).Methods("POST")
	//r.HandleFunc("/event/{id}", UpdateEventHandler).Methods("PUT")

	//sql-migrateでマイグレーション後に成功
	r.HandleFunc("/song", CreateSongHandler).Methods("POST")

	r.HandleFunc("/signup", SignUpHandler).Methods("POST")
	//r.HandleFunc("/login", LoginHandler).Methods("POST")
	//①どちらでもできてそう
	//srv := &http.Server{
	//	Handler:      r,
	//	Addr:         "127.0.0.1:8080",
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}
	//log.Fatal(srv.ListenAndServe())

	//②どちらでもできてそう
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
	//http.ListenAndServe(":8080", nil)
}
