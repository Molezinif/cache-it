# cache-it

Servidor de cache em Go, inspirado no Redis, escrito como projeto de estudo.

O Redis é um dos softwares mais bem desenhados que existem. A meta aqui é
entender por que, reconstruindo o núcleo dele em Go, e depois experimentar
ideias que o Redis não tem ou que ficariam mais naturais em Go.

## Como rodar

```
go run ./cmd/cache-it
```

O servidor escuta na porta 6380 e fala RESP, o protocolo do Redis. Qualquer
cliente Redis funciona:

```
redis-cli -p 6380 PING
```

Para ver os bytes crus da resposta:

```
printf 'PING\r\n' | nc localhost 6380
```

## Desenvolvimento

```
gofmt -l .
go vet ./...
go test -race ./...
```

Commits em inglês, no formato Conventional Commits. Convenções completas em
`CLAUDE.md`.

### Git hooks

As checagens acima rodam sozinhas via [lefthook](https://lefthook.dev).
Instale o binário e ative os hooks uma vez depois de clonar:

```
brew install lefthook
lefthook install
```

Sem Homebrew, `go install github.com/evilmartians/lefthook@latest` também
funciona.

- `pre-commit`: `gofmt -l` nos arquivos Go em stage e `go vet ./...`.
- `pre-push`: `go test -race ./...`.

Sem o binário instalado, os hooks só avisam e o commit passa. A configuração
fica em `lefthook.yml`.

## Layout

- `cmd/cache-it`: o executável. Só monta as peças e chama o servidor.
- `internal/`: toda a lógica, em pacotes por responsabilidade. Criados conforme
  o projeto avança.

## Docs

Plano de estudos, roadmap, decisões de arquitetura e diário ficam no Anytype,
fora deste repositório.
