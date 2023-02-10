package main

import (
	"fmt"

	"github.com/nbd-wtf/go-nostr/nip19"
)

func main() {
	npub := "npub1422a7ws4yul24p0pf7cacn7cghqkutdnm35z075vy68ggqpqjcyswn8ekc"

	if prefix, v, err := nip19.Decode(npub); err == nil {
		pub := v.(string)
		fmt.Println(prefix)
		fmt.Println(pub)
		if s, err := nip19.EncodePublicKey(pub); err == nil {
			fmt.Println(s)
		}
	}
}
