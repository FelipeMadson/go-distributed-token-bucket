# Limitador de Taxa Token Bucket Concorrente em Go

[![Go CI](https://github.com/FelipeMadson/go-distributed-token-bucket/actions/workflows/ci.yml/badge.svg)](https://github.com/FelipeMadson/go-distributed-token-bucket/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/badge/release-v1.0.0-green.svg)](https://github.com/FelipeMadson/go-distributed-token-bucket/releases)

Motor de controle de taxa (`Rate Limiting`) thread-safe e de alta concorrência em Go. Implementa o algoritmo Token Bucket com suporte a particionamento por chaves e zero corridas de dados.

## Recursos de Engenharia

- **Isolamento de Concorrência**: Sincronização via `sync.RWMutex` eliminando contenção desnecessária em leituras.
- **Matemática Contínua de Reabastecimento**: Baseada no delta de nanossegundos, dispensando timers periódicos pesados.
- **Auditoria de Concorrência**: 100% dos testes aprovados com detector de raça (`go test -race`).

## Execução dos Testes

```bash
go test -v -race ./...
```

## Autor

Desenvolvido por **[Felipe Madison](https://github.com/FelipeMadson)**.
