package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	_ "time"

	log "github.com/Sirupsen/logrus"
	_ "github.com/hpcloud/tail"
	. "github.com/julienschmidt/sse"
)

const listen = ":7001"

func main() {

	streamer := New()

	go func() {

		for {
			reader, _ := os.OpenFile("/tmp/events", os.O_RDONLY, 0600)
			scanner := bufio.NewScanner(reader)
			for scanner.Scan() {
				fmt.Println(scanner.Text()) // Println will add back the final '\n'
				streamer.SendString("", "", scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				log.Fatal("reading standard input:", err)
			}
		}
	}()

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Client Subscribed to events!")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		streamer.ServeHTTP(w, nil)
	})

	// Start the server and listen forever on port 8001
	log.Fatal(http.ListenAndServe(listen, nil))
	log.Warn(fmt.Sprintf("SSE Server running on port %s", listen))
}
