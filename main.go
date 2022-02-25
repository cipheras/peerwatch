//go:generate goversioninfo -icon=peerwatch.ico -64
package main

import (
	"embed"
	"log"
	"net"
	"net/http"
	"peerwatch/src"
)

var (
	//go:embed www
	fs  embed.FS
	ln  net.Listener
	err error
)

func main() {
	ln, err = net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal("failed to start listener")
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(http.FS(fs)))
	src.Peerwatch(ln)
}
