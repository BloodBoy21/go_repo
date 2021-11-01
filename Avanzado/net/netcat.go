package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var(
	port = flag.Int("p", 3090, "port to listen on")
	host = flag.String("h", "localhost", "hostname to listen on") 
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()
	CopyContents(conn, os.Stdin)
	conn.Close()
	<-done

}
func CopyContents(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}