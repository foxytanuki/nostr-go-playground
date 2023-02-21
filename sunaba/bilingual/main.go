package main

import (
	"context"
	"fmt"
	"log"
	"os"

	nostrevent "github.com/foxytanuki/go-nostr-event"
	"github.com/joho/godotenv"
	"github.com/nbd-wtf/go-nostr"
)

type Lang struct {
	Relays      []string
	Metadata    nostrevent.MetadataContent
	TextContent string
}

func main() {
	loadEnv()

	sk := os.Getenv("PRIVATE_KEY")
	fmt.Println(sk)

	ja := Lang{
		Relays: []string{
			"wss://relay.nostr.wirednet.jp",
		},
		Metadata: nostrevent.MetadataContent{
			Name:        "bilingual_bot",
			DisplayName: "バイリンガル君",
		},
		TextContent: "やっほ〜",
	}
	en := Lang{
		Relays: []string{
			"wss://relay.damus.io",
		},
		Metadata: nostrevent.MetadataContent{
			Name:        "bilingual_bot",
			DisplayName: "bilingual boy",
		},
		TextContent: "hiii",
	}

	for _, lang := range []Lang{ja, en} {
		ev := generateMetadataEvent(sk, lang.Metadata)
		publishEvent(ev, lang.Relays)
		ev = generateTextEvent(sk, lang.TextContent)
		publishEvent(ev, lang.Relays)
	}
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("failed to load env file")
	}
}

func generateMetadataEvent(sk string, c nostrevent.MetadataContent) nostr.Event {
	ev := nostrevent.NewMetadata(c)
	if err := ev.SignPk(sk); err != nil {
		log.Fatalln(err)
	}
	return ev.Event
}

func generateTextEvent(sk string, c string) nostr.Event {
	ev := nostrevent.NewNote(c)
	if err := ev.SignPk(sk); err != nil {
		log.Fatalln(err)
	}
	return ev.Event
}

func publishEvent(ev nostr.Event, relays []string) {
	for _, url := range relays {
		relay, e := nostr.RelayConnect(context.Background(), url)
		if e != nil {
			fmt.Println(e)
			continue
		}
		fmt.Println("published to ", url, relay.Publish(context.Background(), ev))
	}
}
