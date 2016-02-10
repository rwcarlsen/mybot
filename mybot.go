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
		"I heard Pat got engaged to a state!",
	}

	for {
		// read each incoming message
		m, err := getMessage(ws)
		if err != nil {
			log.Print(err)
			continue
		}

		switch m.Type {
		case "message":
			if strings.Contains(m.Text, self.Name) && m.User != usermap["r"] {
				m.Text = "Careful what you say - I am watching..."
				postMessage(ws, m)
			}

			if strings.Contains(m.Text, self.Name) && m.User == usermap["r"] {
				i := rand.Intn(len(robtext))
				m.Text = robtext[i]
				postMessage(ws, m)
			}

			switch m.User {
			case usermap["shriwise"]:
				i := rand.Intn(len(pattext))
				m.Text = pattext[i]
				postMessage(ws, m)
			case usermap["r"]:
				m.Text = "Hi Robert, "
			}
		}
	}
}
