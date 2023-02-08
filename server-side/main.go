package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	createSocket()
}
func createSocket() {
    sockAddr:= flag.String("sockAddr", "/tmp/echo.sock", "daemon control listen multiaddr")
    flag.Parse()
	if err := os.RemoveAll(*sockAddr); 
	err != nil {
        log.Fatal(err)
    }

	socket, err := net.Listen("unix", *sockAddr)
    if err != nil {
		log.Fatal("listen error:", err)
    }
	defer socket.Close()

	log.Println("Socket created, socket path:",*sockAddr)
    log.Println("Waiting message...")
    for {
        conn, err := socket.Accept()
        if err != nil {
            log.Fatal("accept error:", err)
        }
        go doEchoSocket(conn)
    }
}
func doEchoSocket(conn net.Conn) {
    defer conn.Close()
    for {
		//read message
            buf := make([]byte, 512)
            count, err := conn.Read(buf)
            if err != nil {
                if err != io.EOF {
                    log.Println("Error reading message", "error", err)
                }
                return
            }
            data := buf[0:count]
            log.Println("Incoming message:", string(data))

		//write message
		message:=[]byte("Hello!")
        _, err = conn.Write(message)
        if err != nil {
            log.Fatal("Error write message: ", err)
			return
			// break
        }
    }
}