package sunaba

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

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
