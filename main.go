package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

var serverId string

func giveTimestamp(w http.ResponseWriter, r *http.Request) {
	time := time.Now()
	fmt.Fprintf(w, "From Pod: %s @ %s", serverId, time)
}

func main() {
	uuid := uuid.NewV4()

	serverId = uuid.String()
	fmt.Println("Starting Server on 8080 ....")

	http.HandleFunc("/", giveTimestamp)
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		panic(err)
	}
}
