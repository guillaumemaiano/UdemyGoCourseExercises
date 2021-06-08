package main

import "log"

func checkLinks( urls []string) {
	for _, url := range urls {
		checkLink(url)
	}
}

func checkLink(url string) {
	log.Printf("Checking %s\n", url)
}
