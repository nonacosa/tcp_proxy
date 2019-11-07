package main

import (
	"../gegosc"
	_ "../routers"
	"../tools"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//todo custom
const (
	devEs     = "dev.es"
	devConsul = "dev.consul"
	devKong   = "dev.kong"
	devKonga  = "dev.konga"
	devMinio  = "dev.minio"

	prodEs     = "prod.es"
	prodConsul = "prod.consul"
	prodKong   = "prod.kong"
	prodKonga  = "prod.konga"
	prodMinio  = "prod.minio"

	globElk = "glob.elk"

	springboardMachine = "http://192.168.18.153:"
	//springboardMachine = "http://localhost:"
)

//todo custom

// remote address and local port dictionary
var remoteAddress = map[string][]string{
	devEs:     {"9200", "9200", "172.24.167.144", "/_plugin/head/"},
	devConsul: {"8500", "8500", "172.24.167.144", ""},
	devKong:   {"8000", "8000", "172.24.167.144", ""},
	devKonga:  {"1337", "1337", "172.24.167.144", ""},
	devMinio:  {"9000", "9000", "172.24.167.144", ""},

	prodEs:     {"9200", "19200", "172.24.167.143", "/_plugin/head/"},
	//prodEs:     {"60603", "19200", "39.104.96.233", "/_plugin/head/"},
	prodConsul: {"8500", "18500", "172.24.167.146", ""},
	prodKong:   {"8000", "18000", "172.24.167.143", ""},
	prodKonga:  {"1337", "11337", "172.24.167.143", ""},
	prodMinio:  {"9000", "19000", "172.24.167.143", ""},

	globElk: {"5601", "5601", "172.24.167.150", ""},
}

func main() {

	for _,v :=range remoteAddress{
		fromPort, err := strconv.Atoi(v[1])
		if err != nil {
			log.Println("address error or cast number error")
		}
		toPort, err := strconv.Atoi(v[0])
		if err != nil {
			log.Println("address error or cast number error")
		}
		ip := v[2]
		go gegosc.ProxyStart(fromPort, toPort, ip)

	}

	log.Println("英立讯阿里云VPN端口转发程序已启动 ...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		url := strings.TrimLeftFunc(strings.TrimSpace(r.RequestURI), tools.FirstIsSlash)

		remoteInfoInfo := remoteAddress[url]
		fromPort, err := strconv.Atoi(remoteInfoInfo[1])
		if err != nil {
			log.Println("address error or cast number error")
		}
		//toPort, err := strconv.Atoi(remoteInfoInfo[0])
		//if err != nil {
		//	log.Println("address error or cast number error")
		//}
		//ip := remoteInfoInfo[2]
		path := remoteInfoInfo[3]

		//go gegosc.ProxyStart(fromPort, toPort, ip)
		http.Redirect(w, r, springboardMachine+strconv.Itoa(fromPort)+path, http.StatusMovedPermanently)

		println()
		w.Write([]byte("无匹配的转发，请核对地址。"))
		sayBye(w, r)
	})

	http.HandleFunc("/bye", sayBye)
	log.Println("Starting v1 server  httpServer power by golang !")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,this is v1 httpServer power by @nonacosa "))
}
