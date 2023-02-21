package sunaba

import (
	"encoding/json"
	"fmt"
	"log"

	nostrevent "github.com/foxytanuki/go-nostr-event"
	"github.com/nbd-wtf/go-nostr"
)

func useNostrevent() {
	sk := nostr.GeneratePrivateKey()
	cev := nostrevent.NewNote("hi")
	if err := cev.SignPk(sk); err != nil {
		log.Fatal(err)
	}

	cev.SignPk(sk)

	// print the event
	b, err := json.MarshalIndent(cev, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
