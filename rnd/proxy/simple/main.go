package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const (
	proxyAddr    = ":4000"
	upstreamAddr = ":4001"
)

func main() {
	fmt.Println("main. proxy.")

	ln, err := net.Listen("tcp", proxyAddr)
	if err != nil {
		log.Fatalf("listen fail: %v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept fail: %v", err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	log.Printf("proxy conn accepted: local - %v, remote - %v", conn.LocalAddr(), conn.RemoteAddr())
	defer func() {
		log.Printf("proxy conn disconnected: local - %v, remote - %v", conn.LocalAddr(), conn.RemoteAddr())
	}()

	upConn, err := net.Dial("tcp", upstreamAddr)
	if err != nil {
		log.Printf("upstream dial fail: %v", err)
	}

	go func() {
		_, err := io.Copy(upConn, conn)
		if err != nil {
			log.Printf("proxy conn to upConn copy fail: %v", err)
		}
	}()
	if _, err := io.Copy(conn, upConn); err != nil {
		log.Printf("proxy upConn to conn copy fail: %v", err)
	}
}
