package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elliotforbes/go-websockets-tutorial/entity"
	"github.com/gorilla/websocket"
)

type Helper struct {
}

func Sender(conn *websocket.Conn) {

	message := []byte("message envoyer avec success")
	err := conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		return
	}

}

func (*Helper) Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func (*Helper) Traitement(ws *websocket.Conn, event string) {
	content, err := ioutil.ReadFile(event)
	if err != nil {
		log.Fatal(err)
	}
	if json.Valid(content) {
		var payload []entity.Type
		err = json.Unmarshal(content, &payload)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}

		if payload[0].Nombre == "1" {
			log.Printf("============================================== File has changed, new hash")
			fmt.Println("content", content)
			Sender(ws)

		} else {
			fmt.Println("pas de message envoyer")
		}
	}
}
