package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	proxyAddr    = "localhost:4000"
	upstreamAddr = "localhost:4001"
)

func main() {
	fmt.Println("main. proxy.")

	// Flags
	certFile := flag.String("certfile", "./tmp/cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "./tmp/key.pem", "key PEM file")
	clientCertFile := flag.String("clientcert", "./tmp/clientcert.pem", "certificate PEM for client authentication")
	clientKeyFile := flag.String("clientkey", "./tmp/clientkey.pem", "key PEM for client")
	flag.Parse()

	// Load our client certificate and key.
	clientCert, err := tls.LoadX509KeyPair(*clientCertFile, *clientKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Trusted client certificate.
	clientCertBytes, err := os.ReadFile(*clientCertFile)
	if err != nil {
		log.Fatal(err)
	}
	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCertBytes)

	// Trusted server certificate.
	srvCertBytes, err := os.ReadFile(*certFile)
	if err != nil {
		log.Fatal(err)
	}
	srvCertPool := x509.NewCertPool()
	if ok := srvCertPool.AppendCertsFromPEM(srvCertBytes); !ok {
		log.Fatalf("unable to parse cert from %s", *certFile)
	}

	// server cert
	servCert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatalf("serv cert load: %v", err)
	}

	// server tls config
	tlsConf := &tls.Config{
		MinVersion:               tls.VersionTLS13,
		PreferServerCipherSuites: true,
		ClientCAs:                clientCertPool,
		ClientAuth:               tls.RequireAndVerifyClientCert,
		Certificates:             []tls.Certificate{servCert},
	}

	// client tls config
	clnTLSConf := &tls.Config{
		RootCAs:      srvCertPool,
		Certificates: []tls.Certificate{clientCert},
	}

	ln, err := tls.Listen("tcp", proxyAddr, tlsConf)
	if err != nil {
		log.Fatalf("listen fail: %v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept fail: %v", err)
		}

		go handleConn(conn, clnTLSConf)
	}
}

func handleConn(conn net.Conn, tlsConf *tls.Config) {
	log.Printf("proxy conn accepted: local - %v, remote - %v", conn.LocalAddr(), conn.RemoteAddr())
	defer func() {
		log.Printf("proxy conn disconnected: local - %v, remote - %v", conn.LocalAddr(), conn.RemoteAddr())
	}()

	upConn, err := tls.Dial("tcp", upstreamAddr, tlsConf)
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
