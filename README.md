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

## Useful Links

- https://github.com/aljazceru/awesome-nostr
- https://github.com/fiatjaf/noscl
