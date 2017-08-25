package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/barcatfigaro/namtsuabot/namtsuabot"
)

func main() {

	file, err := ioutil.ReadFile("token")
	if err != nil {
		log.Fatalf("missing token file; %v\n", err)
	}
	bot := namtsuabot.New(strings.TrimSpace(string(file)))
	if bot != nil {
		bot.Connect(true)
	}
}
