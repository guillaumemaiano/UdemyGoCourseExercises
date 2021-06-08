package main

import "log"
import "fmt"

func main() {
	urls := []string{
	"www.maiano.fr",
	"guillaume.maiano.fr",
	"minicampaigns.dragonshard.org",
	"www.flyingfortressgames.com",
	"devlog.flyingfortressgames.com",
	}

	fmt.Println("Intend to check the urls in the array ")
	for _, value := range urls {
		  log.Printf("- %s\n", value)
	}
	fmt.Println("### Checking urls now. ###")
	checkLinks(urls)
}
