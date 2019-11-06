package routers

import (
	"../controllers"
	"../gegosc"
	"net/http"
)

func init() {

	http.HandleFunc("/home", gegosc.Handler(controllers.Home))

	http.HandleFunc("/generate", gegosc.Handler(controllers.Generate))

	//box := rice.MustFindBox("../assets")
	//cssFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	//http.Handle("/static/", cssFileServer)
	//
	//
	//project := http.StripPrefix("/load/", http.FileServer(rice.MustFindBox("../").HTTPBox()))
	//http.Handle("/load/", project)
	//
	//fmt.Println(tools.IOReadDir("./"))
}
