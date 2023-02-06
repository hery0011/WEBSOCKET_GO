package controller

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/elliotforbes/go-websockets-tutorial/helper"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

type Controller struct {
}

var Helper helper.Helper

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	log.Println("Client Successfull Connected...")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	VerificationFichier(ws)
	go Helper.Reader(ws)

}

func VerificationFichier(ws *websocket.Conn) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					data, err := ioutil.ReadFile(event.Name)
					if err != nil {
						log.Println("error reading file:", err)
						continue
					}

					hash := sha256.Sum256(data)
					fmt.Printf("SHA-256 hash of %s: %x\n", event.Name, hash)
					fmt.Println(event.Name)

					/*==============debut traitement fichier changer========================*/
					Helper.Traitement(ws, event.Name)
					/*==========================fin traitement===================================*/

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("data")
	if err != nil {
		log.Fatal(err)
	}
	<-done

}

func (*Controller) SetupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

// func AlgorithmeHashage(w http.ResponseWriter, r *http.Request, ws *websocket.Conn) {
// 	var lastHash string
// 	for {
// 		content, err := ioutil.ReadFile("data/config.json")
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		hash := sha256.Sum256(content)
// 		hashHex := hex.EncodeToString(hash[:])

// 		if lastHash != "" && lastHash != hashHex {

// 			prodDir := "data/config.json"

// 			content, err := ioutil.ReadFile(prodDir)

// 			if err != nil {
// 				log.Fatal("Error when opening file: ", err)
// 			}

// 			var payload []entity.Type
// 			err = json.Unmarshal(content, &payload)
// 			if err != nil {
// 				log.Fatal("Error during Unmarshal(): ", err)
// 			}

// 			//fmt.Println(payload[0].Nombre)
// 			if payload[0].Nombre == "1" {
// 				// sender(ws)
// 				log.Printf("============================================== File has changed, new hash: %s", hashHex)

// 				sender(ws)
// 				// reader(ws)

// 			} else {
// 				fmt.Println("pas de message envoyer")
// 			}
// 		} else {
// 			log.Printf("File has not changed, hash: %s", hashHex)
// 		}

// 		lastHash = hashHex

// 		time.Sleep(time.Second)
// 	}
// }
