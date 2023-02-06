package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/elliotforbes/go-websockets-tutorial/controller"
)

var Alarm_controllers controller.Controller

func main() {
	// c := gocron.NewScheduler(time.Local)
	// c.Every(1).Minute().Do(
	// 	func() { AddConfigJson() },
	// )
	// c.StartAsync()

	AddConfigJson()
	fmt.Println("Go websocket")
	Alarm_controllers.SetupRoutes()
	log.Fatal(http.ListenAndServe(":9999", nil))
	// Alarm_controllers.AlgorithmeHashage()

}

func AddConfigJson() {

	var dataConfig []struct {
		Nombre string `json:"Nombre"`
	} = []struct {
		Nombre string `json:"Nombre"`
	}{
		{Nombre: "1"},
	}

	jsonInfo, _ := json.Marshal(dataConfig)

	file, err := os.Create("data/config.json")
	if err != nil {
		fmt.Println(err)
		log.Print("file create error", err)
		return
	} else {
		l, err := file.Write(jsonInfo)
		if err != nil {
			log.Print("file write error", err)
			file.Close()
			return
		}
		fmt.Println(l, "bytes written successfully")
		err = file.Close()
		if err != nil {
			log.Print("file close error", err)
			return
		}
	}
}
