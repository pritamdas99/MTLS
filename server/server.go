package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {
		log.Println("new request")
		fmt.Fprintf(writter, "hello world\n")
	})
	caCertFile, err := ioutil.ReadFile("../cert/ca.crt")
	if err != nil {
		log.Fatalf("error reading CA certificate: %v", err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCertFile)

	server := http.Server{
		Addr:    ":9091",
		Handler: handler,
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  certPool,
			MinVersion: tls.VersionTLS12,
		},
	}

	if err := server.ListenAndServeTLS("../cert/server.crt", "../cert/server.key"); err != nil {
		log.Fatalf("error listening to port: %v\n", err)
	}
}
