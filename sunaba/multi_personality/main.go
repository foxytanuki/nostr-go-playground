package main

import (
	"context"
	"fmt"
	"log"

	nostrevent "github.com/foxytanuki/go-nostr-event"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func main() {
	sk, _ := newAccount()
	relays := []string{
		"wss://relay.nostr.wirednet.jp",
		"wss://relay.damus.io",
	}
	m := nostrevent.MetadataContent{
		Name:        "multi_personality",
		DisplayName: "多重人格くん",
	}
	mev := generateMetadataEvent(sk, m)

	for _, r := range relays {
		c := fmt.Sprintf("note from %s", r)
		ev := generateTextEvent(sk, c)
		publishEvent(mev, r)
		publishEvent(ev, r)
	}
}

func newAccount() (string, string) {
	sk := nostr.GeneratePrivateKey()
	pub, err := nostr.GetPublicKey(sk)
	if err != nil {
		log.Fatalln("failed to get public key")
	}
	bech, err := nip19.EncodePublicKey(pub)
	if err != nil {
		log.Fatalln("failed to encode public key to bech32")
	}
	fmt.Printf("sk: %s\npub: %s\nbech: %s\n", sk, pub, bech)
	return sk, pub
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

func publishEvent(ev nostr.Event, relayUrl string) {
	relay, err := nostr.RelayConnect(context.Background(), relayUrl)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("published to ", relayUrl, relay.Publish(context.Background(), ev))
}
