package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

const (
	RELAY = "ws://localhost:8008"
)

func main() {
	relay, err := nostr.RelayConnect(context.Background(), RELAY)
	if err != nil {
		panic(err)
	}

	npub := "npub1lx8clymua5e5sfje6kax7shmc3ffwxj8945qu4uwzj5cmnuv0p2swtanp7"
	// nsec := "nsec1v375c9x9vzsa9gdta6c36dq2nhf6pra98etndu0yet78svy62a0qceu75r"

	var filters nostr.Filters
	if _, v, err := nip19.Decode(npub); err == nil {
		pub := v.(string)
		filters = []nostr.Filter{{
			Kinds:   []int{1},
			Authors: []string{pub},
			Limit:   1,
		}}
	} else {
		panic(err)
	}

	// ctx, cancel := context.WithCancel(context.Background())
	ctx := context.Background()
	log.Default().Println("start subscribing")
	sub := relay.Subscribe(ctx, filters)

	go func() {
		<-sub.EndOfStoredEvents
		// handle end of stored events (EOSE, see NIP-15)
	}()

	for ev := range sub.Events {
		// handle returned event.
		// channel will stay open until the ctx is cancelled (in this case, by calling cancel())

		fmt.Println(ev.ID)
	}
}
