package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var port = flag.Int("port", 8000, "port number")

func main() {
	tz, ok := os.LookupEnv("TZ")
	if !ok {
		tz = "Local"
	}
	location, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server is listening on port %d...\n", *port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConnection(conn, location)
	}
}

func handleConnection(c net.Conn, loc *time.Location) {
	fmt.Printf("Handle connection... Remote addr: %v\n", c.RemoteAddr())
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			fmt.Printf("Client disconnected. %v\n", c.RemoteAddr())
			return
		}
		time.Sleep(1 * time.Second)
	}
}
