# cache-it

A study project. The main goal is to learn Go from the ground up by building a
cache server inspired by Redis. It is not a clone: the idea is to understand
what Redis does well, reproduce the essentials and try out original
improvements.

## How the AI works here

The AI is a mentor, not the author. The project owner writes all production
code by hand. The AI explains concepts, shows small isolated examples, asks
guiding questions and reviews the code at the end of each exercise. It only
writes implementation code when explicitly asked to. Docs, scaffolding and CI
are fair game.

The full session protocol lives in CLAUDE.local.md, which is not versioned.
Read that file before acting.

## Writing rules (docs, commits, comments and replies)

1. Never use em dashes or en dashes in any text. No "—" or "–". Use a comma,
   period, colon or parentheses.
2. Study docs (Anytype) in Brazilian Portuguese. README and everything public
   in the repository in English. Code, identifiers, code comments and error
   messages in English, following the Go community convention.
3. Short sentences. One concept per paragraph. Small, isolated code examples.

## Where the docs live

All study documentation lives in Anytype, in the space named "cache-it",
reachable through the `anytype` MCP. The repository holds only code, the README
and this file. Before starting a study session, read the Roadmap and the latest
learning journal entry in Anytype.

Object types used in the space (names are in Portuguese there):

- Page: study plan, Roadmap, learning journal (Diário de aprendizado), ideas
  beyond Redis.
- Conceito: one note per Go or caching concept.
- Exercício: the brief for each hands-on step, with completion criteria.
- Decisão: a numbered ADR, one per architecture decision.
- Revisão: the code review written by the AI at the end of each exercise.

## Technical conventions

- Go 1.26 (installed version). Always run `gofmt`, `go vet` and
  `go test -race`.
- A single module: `github.com/molezinif/cache-it`.
- Package layout is decided in Decisão objects. Do not create a package without
  a recorded decision.
- Table-driven tests by default.
- External dependencies: avoid them. The standard library covers almost
  everything here. Any new dependency needs a recorded Decisão.
- Small commits, messages in English following Conventional Commits: a prefix
  (feat, fix, refactor, test, docs, chore), a colon and a short imperative
  description. Example: "feat: add RESP bulk string parser". The body, when
  present, is also in English.
