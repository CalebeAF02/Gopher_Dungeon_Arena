# Meta 8: Finalização Académica

Passos para concluir o projeto com boas práticas e documentação técnica.

## Passo 1 – Usar `context.Context`
1. Crie um contexto raiz no `main`:
   ```go
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   ```
2. Passe `ctx` para goroutines de bots e projéteis.
3. Cada goroutine deve sair quando `ctx.Done()` for acionado.

## Passo 2 – Cancelar goroutines de forma limpa
1. Quando o jogo terminar, chame `cancel()`.
2. As goroutines recebem o sinal e finalizam suas loops.
3. Isso evita goroutines “órfãs” após o fim do jogo.

## Passo 3 – Tratar erros idiomaticamente
1. Sempre capture erros em funções que possam falhar.
2. Use o padrão Go:
   ```go
   if err != nil {
       return err
   }
   ```
3. No `main()`, logue ou encerre conforme necessário.

## Passo 4 – Preparar o relatório comparativo
1. Explique por que Go é bom neste projeto:
   - CSP e canais para concorrência
   - goroutines leves
   - composição simples
2. Compare com:
   - Haskell: foco funcional e imutabilidade
   - Prolog: lógica declarativa e busca
3. Destaque vantagens práticas do Go para jogos em tempo real.

## Passo 5 – Documentar o projeto
1. Adicione comentários claros no código.
2. Crie um `README.md` ou inclua um relatório no próprio `SUMARIO.md`.
3. Liste requisitos, comandos de execução e decisões técnicas.

## Resultado esperado
- O projeto finaliza sem goroutines vazando.
- Os erros são tratados de forma limpa.
- Você tem material acadêmico para explicar as escolhas.

## Conclusão
Com essas etapas concluídas, seu projeto estará pronto para defesa e demonstração.
