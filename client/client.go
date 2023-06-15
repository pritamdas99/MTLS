package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	cert, err := ioutil.ReadFile("../cert/ca.crt")
	if err != nil {
		log.Fatalf("could not open certificate file: %v\n", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert)
	certificate, err := tls.LoadX509KeyPair("../cert/client.crt", "../cert/client.key")
	client := http.Client{
		Timeout: time.Minute * 1,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{certificate},
			},
		},
	}

	res, err := client.Get("https://pritam:9091")
	if err != nil {
		log.Fatalf("error making get request: %v\n", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading response: %v\n", err)
	}
	fmt.Println(string(body))
}
