package sunaba

import (
	"fmt"
	"log"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

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
	fmt.Println(sk, pub, bech)
	return sk, pub
}
