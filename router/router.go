package router

//"awesomeProject/controller"
//"log"
//"net/http"

/**
 * GetRouter
 * root router
 */
//ginを使うのか
//main.goから呼んで使っている
func GetRouter() {
	//func GetRouter() *gin.Engine {
	//	r := gin.Default()

	//static
	//r.Static("/js", "./public/js")
	//r.Static("/css", "./public/css")
	//r.Static("/images", "./public/images")
	//
	//r.LoadHTMLGlob("views/*")

	//r := mux.NewRouter()
	//
	//r.HandleFunc("/", controller.Index).Methods("GET")
	//r.HandleFunc("/login", controller.Login).Methods("GET")
	//r.HandleFunc("/oauth", controller.OAuth).Methods("GET")
	//r.HandleFunc("/api/tracks", controller.GetPlayList).Methods("GET")

	//r.GET("/", controller.Index)
	//r.GET("/login", controller.Login)
	//r.GET("/oauth", controller.OAuth)
	//r.GET("/api/tracks", controller.GetPlayList)

	//追加した
	//if err := http.ListenAndServe(":8080", r); err != nil {
	//	log.Fatal(err)
	//}
	//
	//return r
}
