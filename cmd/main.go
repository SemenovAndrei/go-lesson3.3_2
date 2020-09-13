package main

import (
	"context"
	"github.com/i-hit/go-lesson3.3_2.git/pkg/qr"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const line = "www.netology.ru"
const filename = "netology.png"

func main() {

	urlQR, ok := os.LookupEnv("URL")
	if !ok {
		log.Println("URL err")
		os.Exit(1)
	}
	versionQR, ok := os.LookupEnv("VERSION")
	if !ok {
		log.Println("version err")
		os.Exit(1)
	}
	timeoutQR, ok := os.LookupEnv("TIMEOUT")
	if !ok {
		log.Println("timeout err")
		os.Exit(1)
	}
	timeout, err := strconv.Atoi(timeoutQR)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	con, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	svc := qr.NewService(
		urlQR,
		versionQR,
		con,
		&http.Client{},
	)
	data, err := svc.Encode(line)

	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}