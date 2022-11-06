package main

import (
	"flag"

	"github.com/030/multipart/internal/app/multipart/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	url := flag.String("url", "", "To what URL the multipart should be posted")
	user := flag.String("user", "", "What username should be used to authenticate to the URL")
	pass := flag.String("pass", "", "The password that should be used to authenticate to the URL")
	elements := flag.String("F", "", "The elements that should be constructed as a multipart")
	debug := flag.Bool("d", false, "Whether debug logging should be enabled")

	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}

	u := utils.Upload{URL: *url, Username: *user, Password: *pass}
	if err := u.MultipartUpload(*elements); err != nil {
		log.Fatal(err)
	}
}
