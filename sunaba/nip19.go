package main

import (
	"fmt"

	"github.com/nbd-wtf/go-nostr"
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
	fmt.Println("---")

	priv := "647d4c14c560a1d2a1abeeb11d340a9dd3a08fa53e5736f1e4cafc78309a575e"
	if bech32priv, err := nip19.EncodePrivateKey(priv); err == nil {
		fmt.Println(bech32priv)
		fmt.Println(priv)
		if pub, err := nostr.GetPublicKey(priv); err == nil {
			if bech32pub, err := nip19.EncodePublicKey(pub); err == nil {
				fmt.Println(bech32pub)
				fmt.Println(pub)
			}
		}
	}
	fmt.Println("---")
}
