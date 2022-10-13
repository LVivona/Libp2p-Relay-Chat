# Libp2p-Relay-Chat (Local Network)
## How to run

### 1. Start up Relay
```
go run ./Relay

console:
2022/05/29 15:09:08 ID: Qmba4p7o9SipbwTABNvV1nuNPeE4gXdewxXbJJjvy4qm4h
2022/05/29 15:09:08 Adder: /ip4/x.x.x.x/tcp/36941
2022/05/29 15:09:08 Adder: /ip4/x.x.x.x/tcp/36941
2022/05/29 15:09:08 Adder: /ip6/::1/tcp/43563
```

### 2. Start Node B
```
go run ./NodeB -Address Qmba4p7o9SipbwTABNvV1nuNPeE4gXdewxXbJJjvy4qm4h -id /ip4/x.x.x.x/tcp/36941

console:
2022/05/29 13:58:16 Adder: /ip4/x.x.x.x/tcp/34095
2022/05/29 13:58:16 Adder: /ip4/x.x.x.x/tcp/34095
2022/05/29 13:58:16 Adder: /ip6/::1/tcp/35661
```
### 3. Start Node A
```
go run ./NodeA -Address Qmba4p7o9SipbwTABNvV1nuNPeE4gXdewxXbJJjvy4qm4h -id /ip4/x.x.x.x/tcp/36941 -d /ip4/x.x.x.x/tcp/37549/p2p/QmUc4jfYCRWViUXQS5XC5XZhVswtsPJCBQao3LAMJPcBWX -p /chat

console:
2022/05/29 15:11:16 made it
> 

```

### 4. Start Chatting
enjoy!!!
