//go:build ignore

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func main() {

	sk := nostr.GeneratePrivateKey()
	pub, _ := nostr.GetPublicKey(sk)
	fmt.Println(sk)
	fmt.Println(pub)
	bech, _ := nip19.EncodePublicKey(pub)
	fmt.Println(bech)

	ev0 := nostr.Event{
		PubKey:    pub,
		CreatedAt: time.Now(),
		Kind:      0,
		Tags:      nil,
		Content:   "{\"banner\":\"\",\"website\":\"\",\"picture\":\"\",\"lud16\":\"\",\"display_name\":\"\",\"about\":\"hahaha\",\"name\":\"decoy\",\"nip05\":\"\",\"nip05valid\":false}",
	}
	// calling Sign sets the event ID field and the event Sig field
	ev0.Sign(sk)

	// publish the event to two relays
	for _, url := range []string{"wss://relay.damus.io", "wss://nos.lol", "wss://relay.snort.social"} {
		relay, e := nostr.RelayConnect(context.Background(), url)
		if e != nil {
			fmt.Println(e)
			continue
		}
		fmt.Println("published to ", url, relay.Publish(context.Background(), ev0))
	}
}
