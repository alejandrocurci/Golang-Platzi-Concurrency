package main

// netcat.go -> Client -> it requires the Server to be running first
// build and run this file in different terminals so each running client can talk to everyone else

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("p", 3090, "port")
	host = flag.String("h", "localhost", "host")
)

func main() {
	flag.Parse()

	// Listen on TCP port
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	// Start a goroutine to receive data from the server
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	// Copy what we got to the console line
	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done
}

// CopyContent copies content from src to dst
func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
