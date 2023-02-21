package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	nostrevent "github.com/foxytanuki/go-nostr-event"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func main() {
	fmt.Println("welcome to sunaba")
	sk, _ := newAccount()

	like(sk, "c90c8571c4e4b85f8c5367305cd455e4de488c72dca9ba763109b24f96af47f4", "c6dc2b963a3125b06dc4007fa21075405f53bbcafd3d1ae98d77ba2e434f6947")
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

func like(sk string, note string, event string) {
	// ["EVENT","f2a85e41-9323-4837-a65d-f7278fe0b640",{"content":"+","created_at":1676231763,"id":"6e3ba6c69656030ce5ddfd65d9473f8602c88341ec440e74cdf6354d44649163","kind":7,"pubkey":"c6dc2b963a3125b06dc4007fa21075405f53bbcafd3d1ae98d77ba2e434f6947","sig":"f59143f9706b22f659cdf5f50b81781a1c691b468d845c8fae4b86a1b9e9b6eb68421ac49c4381551d4ff18cc1be22cc128249b11d40fe2506c8fdcdcca24e0a","tags":[["e","60b91d62ec4529faa9e44260681e9e882e5a400b9766c7f957c521ffdde78583"],["p","c6dc2b963a3125b06dc4007fa21075405f53bbcafd3d1ae98d77ba2e434f6947"]]}]
	// The last e tag MUST be the id of the note that is being reacted to.
	// The last p tag MUST be the pubkey of the event being reacted to.
	e := nostr.Tag{"e", note}
	p := nostr.Tag{"p", event}
	tags := nostr.Tags{
		e,
		p,
	}

	ev := nostrevent.NewEvent(nostr.KindReaction, tags, "+")

	// calling Sign sets the event ID field and the event Sig field
	ev.SignPk(sk)

	b, _ := json.Marshal(ev)
	fmt.Println(string(b))

	// publish the event to multiple relays
	for _, url := range []string{"wss://relay.damus.io", "wss://nos.lol", "wss://relay.snort.social"} {
		relay, e := nostr.RelayConnect(context.Background(), url)
		if e != nil {
			fmt.Println(e)
			continue
		}
		fmt.Println("published to ", url, relay.Publish(context.Background(), ev.Event))
	}
}

func nip19Gen() {
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

func variableProfile() {
	sk := nostr.GeneratePrivateKey()
	pub, _ := nostr.GetPublicKey(sk)
	bech, _ := nip19.EncodePublicKey(pub)
	fmt.Println(sk)
	fmt.Println(pub)
	fmt.Println(bech)

	src := "https://api.nostr.watch/v1/online"
	resp, err := http.Get(src)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	s := string(b)
	sp := s[2 : len(s)-2]
	sa := strings.Split(sp, "\",\"")
	// fmt.Println(sa)
	// fmt.Println(len(sa))

	for _, url := range sa {
		fmt.Println(url)
		if url == "wss://nproxy.zerologin.co" {
			continue
		}
		content := fmt.Sprintf("{\"banner\":\"\",\"website\":\"\",\"picture\":\"\",\"lud16\":\"\",\"display_name\":\"\",\"about\":\"Bye\",\"name\":\"%s\",\"nip05\":\"\",\"nip05valid\":false}", url)
		fmt.Println(content)
		ev := nostr.Event{
			PubKey:    pub,
			CreatedAt: time.Now(),
			Kind:      0,
			Tags:      nil,
			Content:   content,
		}
		// calling Sign sets the event ID field and the event Sig field
		ev.Sign(sk)

		relay, e := nostr.RelayConnect(context.Background(), url)
		if e != nil {
			fmt.Println(e)
			continue
		}
		fmt.Println("published to ", url, relay.Publish(context.Background(), ev))
	}
}
