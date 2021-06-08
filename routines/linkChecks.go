package main

import "log"
import "net/http"

func checkLinks( urls []string) {
	for _, url := range urls {
		checkLink(url)
	}
}

func checkLink(url string) {
	log.Printf("Checking %s\n", url)
	httpsUrl := "https://" + url
	_, err := http.Get(httpsUrl)
	if err != nil {
		log.Println(httpsUrl, "could not be reached.")
	} else {
		log.Println(httpsUrl, "is online.")
	}
}
