# cache-it

Projeto de estudo. O objetivo principal é APRENDER Go, do início, construindo um
servidor de cache inspirado no Redis. Não é um clone: a ideia é entender o que o
Redis faz bem, reproduzir o essencial e experimentar melhorias próprias.

## Regras de escrita (valem para docs, commits, comentários e respostas)

1. Nunca usar travessão nem meia-risca em nenhum texto. Nada de "—" ou "–".
   Use vírgula, ponto, dois pontos ou parênteses.
2. Docs em português do Brasil. Código, identificadores, comentários de código e
   mensagens de erro em inglês, seguindo a convenção da comunidade Go.
3. Frases curtas. Um conceito por parágrafo. Exemplos de código pequenos e
   isolados.

## Onde ficam os docs

Toda a documentação de estudo fica no Anytype, no espaço chamado "cache-it",
acessível pelo MCP `anytype`. O repositório guarda só código, README e este
arquivo. Antes de começar uma sessão de estudo, ler no Anytype o Roadmap e a
última entrada do Diário de aprendizado.

Tipos de objeto usados no espaço:

- Page: Plano de estudos, Roadmap, Diário de aprendizado, Ideias além do Redis.
- Conceito: uma nota por conceito de Go ou de cache.
- Exercício: enunciado de cada etapa prática, com critérios de conclusão.
- Decisão: ADR numerado, um por decisão de arquitetura.
- Revisão: revisão de código feita pela IA ao final de cada exercício.

## Instruções locais

O modo de trabalho com a IA (mentoria, protocolo de sessão) fica em
CLAUDE.local.md, que não é versionado. Ler esse arquivo antes de agir.

## Convenções técnicas

- Go 1.26 (versão instalada). Usar sempre `gofmt`, `go vet` e `go test -race`.
- Um módulo só: `github.com/molezinif/cache-it`.
- Layout de pacotes decidido em objetos Decisão. Não criar pacote sem decisão
  registrada.
- Testes em tabela (table-driven) como padrão.
- Dependências externas: evitar. A biblioteca padrão resolve quase tudo aqui.
  Qualquer dependência nova precisa de uma Decisão registrada.
- Commits pequenos, mensagem curta em português no imperativo:
  "adiciona parser RESP para bulk string".
