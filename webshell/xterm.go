package webshell

import (
	"fmt"
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type windowSize struct {
	Rows uint16 `json:"rows"`
	Cols uint16 `json:"cols"`
	X    uint16
	Y    uint16
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	h, err := models.MachineFind(4)
	if err != nil {
		log.Fatal("错误的SSH服务器ID ", err)
	}
	if err := flx.RunSshXtermJs(h, true, w, r); err != nil {
		fmt.Println(err)
	}

}

func RunXterm() {

	r := mux.NewRouter()
	r.HandleFunc("/term", handleWebsocket)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../assets")))
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.WithError(err).Fatal("Something went wrong with the webserver")
	}
}
