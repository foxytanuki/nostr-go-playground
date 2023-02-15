# nostr-go-playground

### Generate randam privkey
`openssl rand -hex 32`


### Nostrill
`./nostrill/nostrill`

https://github.com/jb55/nostril


```
nostril --envelope --sec <key> --content "this is a message" | websocat ws://localhost:8008
```

### noscl
`noscl publish <content>` 


### relayer
```
docker compose up

docker exec -it basic-postgres-1 psql -U nostr nostr
nostr-# \dt
nostr=# SELECT * from event;
```

### nostreq
https://github.com/blakejakopovic/nostreq

````
nostreq --authors c8238017bbd3e488c7814c8f49201e8d21a2f5d560627964e31799563fec80c1 | nostcat wss://relay.damus.io
```

## Useful Links

- https://www.nostr.net/
  - https://github.com/aljazceru/awesome-nostr
- https://github.com/fiatjaf/noscl
- https://github.com/nbd-wtf/go-nostr
- https://github.com/blakejakopovic/nostreq
- https://github.com/hoytech/strfry
- https://api.nostr.watch/
- https://nproxy-test.zerologin.co/
- https://nostr-proxy.inosta.cc/
