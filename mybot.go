package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
)

func init() {
	flag.Usage = func() {
		log.Fatal("Usage: mybot slack-bot-token\n")
	}
}

func main() {
	log.SetFlags(0)
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("Usage: mybot slack-bot-token\n")
	}

	token := flag.Arg(0)

	// start a websocket-based Real Time API session
	ws, self, users := slackConnect(token)
	fmt.Println("mybot listening...")
	_ = users

	usermap := map[string]string{}
	for _, user := range users {
		usermap[user.Name] = user.Id
	}

	pattext := []string{
		"a little less Pat",
		"too much Pat",
		"just right Pat",
		"more Pat",
	}

	robtext := []string{
		"All are behaving within acceptable paremeters.",
		"That biondo - he's a bit of trouble...",
		"Just another day in CNERG.",
		"I heard Pat got engaged to a southern state!",
	}

	for {
		// read each incoming message
		m, err := getMessage(ws)
		if err != nil {
			log.Print(err)
			continue
		}

		lowertext := strings.ToLower(m.Text)
		if m.Type != "message" {
			continue
		}

		switch {
		case m.User == usermap["shriwise"]:
			i := rand.Intn(len(pattext))
			m.Text = pattext[i]
			postMessage(ws, m)
		case m.User == usermap["r"]:
			if strings.Contains(lowertext, self.Name) && strings.Contains(m.Text, "?") {
				if strings.Contains(lowertext, "ready") {
					m.Text = "Locked and loaded."
					postMessage(ws, m)
				} else {
					i := rand.Intn(len(robtext))
					m.Text = robtext[i]
					postMessage(ws, m)
				}
			}
		default:
			if strings.Contains(lowertext, self.Name) {
				m.Text = "Careful what you say - I am watching..."
				postMessage(ws, m)
			}
		}
	}
}
