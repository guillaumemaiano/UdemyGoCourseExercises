package main

import "log"
import "net/http"

func checkLinks( urls []string) {
	c := make(chan string)
	for _, url := range urls {
		go checkLink(url, c)
	}
	reportFails := 0
	for i:= 0; i < len(urls); i++ {
		if ("success" != <-c) {
			reportFails += 1
		}
	}
	if reportFails > 0 {
		log.Printf("Encountered %d failures.\n", reportFails)
	}
}

func checkLink(url string, c chan string) {
	log.Printf("Checking %s\n", url)
	httpsUrl := "https://" + url
	_, err := http.Get(httpsUrl)
	if err != nil {
		log.Println(httpsUrl, "could not be reached.")
		c <- "failed"
	} else {
		log.Println(httpsUrl, "is online.")
		c <- "success"
	}
}
