package cmd

import (
	"flag"
	"log"
	"time"

	"github.com/shuhard1/L2/develop/dev10/internal/client"
	"github.com/shuhard1/L2/develop/dev10/internal/server"
)

var (
	timeoutFlag string
)

func init() {
	flag.StringVar(&timeoutFlag, "timeout", "10s", "timeout durations")
	flag.Parse()
}

func RunNetwork() {

	host := flag.Arg(0)
	port := flag.Arg(1)

	networkType := flag.Arg(2)
	// запускаем сервер
	if networkType == "server" {
		server.Start(server.Config{Host: host, Port: port})
	} else {
		timeout, err := time.ParseDuration(timeoutFlag)
		if err != nil {
			log.Fatalf("incorrect timeout")
		}
		// задаём конфиг клиента
		var conf = client.Config{Host: host, Port: port, Timeout: timeout}
		client.Start(conf)
	}
}
