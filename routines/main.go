package main

import "log"

func main() {
	urls := []string{
	"www.maiano.fr",
	"minicampaigns.dragonshard.org",
	"www.flyingfortressgames.com",
	"devlog.flyingfortressgames.com",
	}

	log.Println("Intend to check the urls in the array ")
	for _, value := range urls {
		  log.Printf("- %s\n", value)
	}

}
