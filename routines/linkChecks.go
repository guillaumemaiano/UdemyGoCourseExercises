package main

import "log"
import "net/http"

func checkLinks( urls []string) {
	c := make(chan string)
	for _, url := range urls {
		go checkLink(url, c)
		// this is a blocking statement and it makes the previous line kinda pointless
		log.Println(<- c)
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
