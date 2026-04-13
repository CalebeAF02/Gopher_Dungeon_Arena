# Plano de Projeto: Gopher Dungeon PvPvE (Versão Final)

Este documento estabelece o roteiro completo para o desenvolvimento de um jogo de arena concorrente em **Go**, focado em demonstrar o domínio do modelo de concorrência CSP e a eficiência da linguagem em sistemas de tempo real.

## 1. Visão Geral do Jogo
- **Conceito:** Uma dungeon onde dois times (3 vs 3) competem para chegar ao fim do mapa.
- **Inimigos:** Bots (Environment) que atacam qualquer jogador próximo.
- **Objetivo Final:** Pelo menos um membro do time chegar vivo à "Extração".
- **Visual:** Geometria simples (Quadrados e Círculos) para foco na lógica de código.

---

## 2. Metas e Passos de Desenvolvimento

### Meta 1: Fundação e Game Loop (A Infraestrutura)
- [Meta 1: Fundação e Game Loop](metas/meta-1-fundacao-game-loop.md)
* **1.1 Inicialização:** Configurar `go mod init` e instalar `ebiten/v2`.
* **1.2 Engine Base:** Implementar a interface `ebiten.Game` com `Update`, `Draw` e `Layout`.
* **1.3 Sistema de Coordenadas:** Criar a matriz do mapa e os limites da janela.

### Meta 2: Entidades e Tipagem (Arquitetura Go)
- [Meta 2: Entidades e Tipagem](metas/meta-2-entidades-tipagem.md)
* **2.1 Interface Entity:** Definir métodos comuns: `Position()`, `Health()`, `Team()`.
* **2.2 Composição de Structs:** Criar structs para `Player`, `Bot` e `Projectile`.
* **2.3 Sistema de Cores:** Atribuição visual (Time A: Verde, Time B: Azul, Bots: Vermelho).

### Meta 3: Concorrência Massiva (O Diferencial de Go)
- [Meta 3: Concorrência Massiva](metas/meta-3-concorrencia-massiva.md)
* **3.1 Goroutines de Bot:** Disparar cada bot como uma goroutine independente (`go runBot(b)`).
* **3.2 Canais de Comunicação:** Criar o `ActionChannel` para receber pedidos de movimento/ataque.
* **3.3 Tickers Assíncronos:** Usar `time.Ticker` para que cada entidade tenha a sua própria "cadência" de decisão.



### Meta 4: Mecânicas de Combate e Projéteis
- [Meta 4: Mecânicas de Combate e Projéteis](metas/meta-4-mecanicas-combate-projeteis.md)
* **4.1 Projéteis em Goroutines:** Cada disparo é uma goroutine que percorre uma trajetória até colidir ou expirar.
* **4.2 Detecção de Colisão (AABB):** Algoritmo simples para verificar interseção entre quadrados.
* **4.3 Feedback de Dano:** Alteração momentânea de cor (flash branco) ao ser atingido.

### Meta 5: Sincronização e Estado Seguro
- [Meta 5: Sincronização e Estado Seguro](metas/meta-5-sincronizacao-estado-seguro.md)
* **5.1 Mutexes de Saúde:** Proteger a variável `HP` com `sync.Mutex` para evitar condições de corrida (Race Conditions).
* **5.2 Árbitro Global:** Uma goroutine central que valida as mensagens do canal e atualiza o estado oficial do mundo.
* **5.3 Logs de Sistema:** Output no terminal mostrando o nascimento e morte de goroutines.

### Meta 6: Dinâmica da Dungeon (Gameplay)
- [Meta 6: Dinâmica da Dungeon](metas/meta-6-dinamica-dungeon.md)
* **6.1 Spawn Progressivo:** Aumentar a quantidade de bots ativos conforme o tempo passa.
* **6.2 Zona de Extração:** Área no final do mapa que valida a condição de vitória do time.
* **6.3 Obstáculos Estáticos:** Paredes simples que bloqueiam movimento e tiros.

### Meta 7: Interface e Input de Utilizador
- [Meta 7: Interface e Input de Utilizador](metas/meta-7-interface-input-usuario.md)
* **7.1 Multi-Input:** Mapeamento de teclas para múltiplos jogadores locais (ou simulados).
* **7.2 HUD (Heads-Up Display):** Barras de vida desenhadas sobre os jogadores.
* **7.3 Ecrãs de Estado:** Mensagens de "Game Over", "Vitória" e "Membros Restantes".

### Meta 8: Finalização Académica (A Defesa)
- [Meta 8: Finalização Académica](metas/meta-8-finalizacao-academica.md)
* **8.1 Gestão com Context:** Usar `context.Context` para cancelar todas as goroutines de forma limpa.
* **8.2 Gestão de Erros:** Tratamento explícito de erros no estilo idiomático do Go.
* **8.3 Relatório Comparativo:** Análise técnica sobre as vantagens do Go face ao paradigma Funcional (Haskell) e Lógico (Prolog).

---

## 3. Justificação Técnica (Para a UnB)

O projeto utiliza as seguintes especialidades da linguagem Golang:

1.  **Modelo CSP:** O jogo não usa memória partilhada para decidir posições; ele comunica intenções através de canais.
2.  **Lightweight Threading:** Demonstra a capacidade de correr centenas de IAs de bots simultaneamente sem sobrecarga de sistema.
3.  **Composição sobre Herança:** Mostra como o Go resolve problemas complexos de jogos sem a necessidade de hierarquias de classes rígidas.
4.  **Segurança e Velocidade:** Combina a performance de uma linguagem compilada com a segurança de um Garbage Collector moderno.

---

## 4. Requisitos de Sistema
- Go 1.20+
- Ebitengine v2.x
- Sistema Operativo: Windows, Linux ou macOS (com suporte a drivers gráficos básicos).