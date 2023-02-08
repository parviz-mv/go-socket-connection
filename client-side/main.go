package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		_, err:= conn.Write([]byte("hi"))
		if err != nil {
			log.Fatal("Write error:", err)
			break
		}
		time.Sleep(3 * 1e9)
		go doEchoSocket(conn)
	}

}
func doEchoSocket(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf[:])
		if err != nil {
			if err != io.EOF {
				log.Println("Error reading message", "error", err)
			}
			return
		}
		data := buf[0:n]
		log.Println("Incoming messsage:", string(data))		
	}
}
