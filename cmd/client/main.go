package main

import (
	"log"
	"net/http"
	"time"

	"github.com/deadlyengineer/socket-golang"
)

func main() {
	ws := websockit.NewWebsocket()
	clientOpts := []websockit.WebsocketClientOption{
		websockit.ClientWithHandshakeTimeout(5 * time.Second),
		websockit.ClientWithProxy(http.ProxyFromEnvironment),
	}
	client, err := ws.ClientSocket("ws://127.0.0.1:8080", nil, clientOpts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.CloseNice()

	pingOpts := []websockit.ClientPingOption{
		websockit.PingWithPongTimeout(10 * time.Second),
		websockit.PingWithPongLog(true),
	}
	client.EnableServerPings(pingOpts...)

	readChan := make(chan []byte)
	go func() {
		if err := client.ServerPingLoop(); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		if err := client.ReadLoop(readChan); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for range ticker.C {
			log.Println("write: 'FooBar'")
			_ = client.WriteTextMessage([]byte("FooBar"))
		}
	}()

	for msg := range readChan {
		log.Printf("read: %s\n", string(msg))
	}
}
