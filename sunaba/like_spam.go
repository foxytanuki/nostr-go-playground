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
	bech, _ := nip19.EncodePublicKey(pub)
	fmt.Println(sk)
	fmt.Println(pub)
	fmt.Println(bech)

	// ["EVENT","f2a85e41-9323-4837-a65d-f7278fe0b640",{"content":"+","created_at":1676231763,"id":"6e3ba6c69656030ce5ddfd65d9473f8602c88341ec440e74cdf6354d44649163","kind":7,"pubkey":"c6dc2b963a3125b06dc4007fa21075405f53bbcafd3d1ae98d77ba2e434f6947","sig":"f59143f9706b22f659cdf5f50b81781a1c691b468d845c8fae4b86a1b9e9b6eb68421ac49c4381551d4ff18cc1be22cc128249b11d40fe2506c8fdcdcca24e0a","tags":[["e","60b91d62ec4529faa9e44260681e9e882e5a400b9766c7f957c521ffdde78583"],["p","c6dc2b963a3125b06dc4007fa21075405f53bbcafd3d1ae98d77ba2e434f6947"]]}]
	e := nostr.Tag{"e", "60b91d62ec4529faa9e44260681e9e882e5a400b9766c7f957c521ffdde78583"}
	p := nostr.Tag{"p", "c6dc2b963a3125b06dc4007fa21075405f53bbcafd3d1ae98d77ba2e434f6947"}
	tags := nostr.Tags{
		e,
		p,
	}

	ev := nostr.Event{
		PubKey:    pub,
		CreatedAt: time.Now(),
		Kind:      7,
		Tags:      tags,
		Content:   "+",
	}

	// calling Sign sets the event ID field and the event Sig field
	ev.Sign(sk)

	// publish the event to multiple relays
	for _, url := range []string{"wss://relay.damus.io", "wss://nos.lol", "wss://relay.snort.social"} {
		relay, e := nostr.RelayConnect(context.Background(), url)
		if e != nil {
			fmt.Println(e)
			continue
		}
		fmt.Println("published to ", url, relay.Publish(context.Background(), ev))
	}
}
