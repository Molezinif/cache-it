# cache-it

Servidor de cache em Go, inspirado no Redis, escrito como projeto de estudo.

O Redis é um dos softwares mais bem desenhados que existem. A meta aqui é
entender por que, reconstruindo o núcleo dele em Go, e depois experimentar
ideias que o Redis não tem ou que ficariam mais naturais em Go.

## Estado

Fase 0: planejamento. Nenhuma linha de Go ainda.

## Como rodar

Ainda não roda. Quando rodar, vai ser:

```
go run ./cmd/cache-it
redis-cli -p 6380 PING
```

## Docs

Toda a documentação de estudo fica no Anytype, no espaço "cache-it".
