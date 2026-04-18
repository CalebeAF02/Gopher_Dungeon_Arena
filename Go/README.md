# Guia de Aprendizado Go

Esta pasta `Go/` foi criada para ser um guia prático de aprendizado da linguagem Go, do mais simples ao mais avançado.

## Objetivo

- Mostrar exemplos reais de código em português do Brasil.
- Explicar conceitos com comentários detalhados.
- Organizar o aprendizado em blocos temáticos.
- Servir como referência rápida enquanto você estuda.

## Estrutura

### 01_Variaveis/
Contém exemplos de tipos básicos e diferentes formas de declarar e instanciar variáveis:
- `string.go` — manipulação de strings, concatenação, interpolação e literais multilinha.
- `integer.go` — todos os tipos inteiros nativos (`int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `byte`, `rune`, `uintptr`).
- `float.go` — `float32` e `float64`, precisão e comparação.
- `boolean.go` — valores booleanos e operadores lógicos.
- `instanciacao.go` — declaração com `var`, inferência com `:=`, `new()`, valores zero e visibilidade de pacote.
- `constantes.go` — declaração de constantes e uso de `iota`.
- `operadores.go` — operadores aritméticos, de atribuição e relacionais.
- `conversoes.go` — conversões seguras entre string, int e bool.
- `complex.go` — tipos complexos (`complex64`, `complex128`).

### 02_Estruturas_Controle/
Exemplos de controle de fluxo:
- `if_else.go` — condicionais simples e compostas.
- `for.go` — loop clássico incrementando com `i := 0; i < n; i++`.
- `while.go` — uso de `for` para comportamento de `while`.
- `range.go` — iteração sobre slices e maps.
- `switch_case.go` — switch por valores e switch de tipos.

### 03_Colecoes/
Estruturas de coleção:
- `arrays.go` — arrays de tamanho fixo.
- `slices.go` — slices dinâmicos, `append`, `copy`.
- `maps.go` — dicionários com chave/valor.
- `make.go` — pré-alocação de memória com `make()`.

### 04_Funcoes/
Conceitos de funções:
- `basico.go` — parâmetros e retorno de valores.
- `multiplos_retornos.go` — funções com retorno de valor e erro.
- `variadicas.go` — funções que aceitam N argumentos.
- `defer.go` — execução adiada para fechamento de recursos.
- `init.go` — função especial executada antes do `main`.

### 05_Estruturas_Dados/
Tipos e abstrações:
- `structs.go` — definição de structs.
- `metodos.go` — métodos associados a structs.
- `interfaces.go` — contratos e polimorfismo.
- `ponteiros.go` — uso de ponteiros com `*` e `&`.

### 06_Tratamento_Erros/
Erros em Go:
- `error_type.go` — tipo `error` nativo.
- `panic_recover.go` — `panic` e `recover`.
- `custom_errors.go` — criação de tipos de erro personalizados.

### 07_Concorrencia/
Concorrência e sincronização:
- `goroutines.go` — execução paralela com `go`.
- `channels.go` — comunicação entre goroutines.
- `select.go` — controle de múltiplos canais.
- `waitgroups.go` — sincronização manual com `sync.WaitGroup`.
- `mutex.go` — exclusão mútua com `sync.Mutex`.

### 08_Modulos_e_Pacotes/
Módulos e visibilidade:
- `go_mod.go` — explicação de `go.mod` e gerência de dependências.
- `visibilidade.go` — exportação com letra maiúscula e símbolos privados.

## Como usar

1. Abra a pasta `Go/` no VS Code.
2. Leia os arquivos na sequência, começando por `01_Variaveis`.
3. Use os comentários como explicação passo a passo.
4. Se quiser testar, copie o código para um arquivo `main.go` ou execute `go test ./Go/...` a partir da raiz do projeto.

## Como evoluir este guia

- Posso adicionar exemplos de `main` para cada capítulo.
- Posso incluir exercícios práticos e respostas.
- Posso criar um fluxograma de estudo com prioridades.

> A pasta já está organizada para ser um guia didático. Se quiser, posso continuar aprimorando com explicações ainda mais completas e exemplos de uso reais.