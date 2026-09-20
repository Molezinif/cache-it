# cache-it

A cache server in Go, inspired by Redis, written as a study project.

Redis is one of the best designed pieces of software out there. The goal here
is to understand why, by rebuilding its core in Go, and then to experiment with
ideas Redis does not have or that feel more natural in Go.

## Status

Early stage. Today the server accepts concurrent TCP connections, one goroutine
per client, and answers `PING` with `PONG` over RESP. The next step is a full
RESP parser.

## Roadmap

- [x] TCP server that answers `PING`.
- [ ] RESP parser and serializer, with table-driven tests.
- [ ] Strings: `SET`, `GET`, `DEL`, `EXISTS`, safe under concurrent access.
- [ ] Expiration: `EXPIRE`, `TTL`, passive and active expiry, graceful shutdown.
- [ ] Lists and hashes.
- [ ] Persistence: AOF with configurable fsync and a binary snapshot.
- [ ] Performance: pprof under `redis-benchmark`, comparing a single mutex,
      sharding and a data-owning goroutine.
- [ ] Sorted sets backed by a skip list.
- [ ] Pub/sub and transactions.
- [ ] Replication.
- [ ] Ideas beyond Redis.

No external dependencies: Go standard library only.

## Running

```
go run ./cmd/cache-it
```

The server listens on port 6380 and speaks RESP, the Redis protocol. Any Redis
client works:

```
redis-cli -p 6380 PING
```

To see the raw response bytes:

```
printf 'PING\r\n' | nc localhost 6380
```

## Development

```
gofmt -l .
go vet ./...
go test -race ./...
```

Commits follow Conventional Commits, in English. Full conventions live in
`CLAUDE.md`.

### Git hooks

The checks above run automatically through [lefthook](https://lefthook.dev).
Install the binary and enable the hooks once after cloning:

```
brew install lefthook
lefthook install
```

Without Homebrew, `go install github.com/evilmartians/lefthook@latest` works
too.

- `pre-commit`: `gofmt -l` on staged Go files and `go vet ./...`.
- `pre-push`: `go test -race ./...`.

If the binary is not installed, the hooks only warn and the commit goes
through. The configuration lives in `lefthook.yml`.

## Layout

- `cmd/cache-it`: the executable. It only wires the pieces together and starts
  the server.
- `internal/`: all the logic, in packages split by responsibility. Created as
  the project moves forward.

## Docs

The study plan, roadmap, architecture decisions and learning journal are kept
in Anytype, outside this repository.

## License

[MIT](LICENSE)
