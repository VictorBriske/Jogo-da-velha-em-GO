# Briske's TOE

**Briske's TOE** é um jogo da velha (tic-tac-toe) desenvolvido em Go utilizando a biblioteca gráfica [Ebiten](https://ebiten.org/). O jogo possui três níveis de dificuldade e uma introdução animada com uma bola e o título estilizado.

## 🎮 Funcionalidades

- Três níveis de dificuldade (Fácil, Médio, Difícil)
- Animação de introdução com texto revelado
- Interface gráfica com mouse
- Jogador contra IA com lógica ajustável
- Reinício rápido com botão ou tecla `R`

## 🛠️ Tecnologias

- [Go (Golang)](https://golang.org/)
- [Ebiten (2D Game Library)](https://ebiten.org/)

## ▶️ Como executar

1. **Clone o repositório:**

```bash
git clone https://github.com/VictorBriske/Jogo-da-velha-em-GO
cd tictactoe
```

2. **Execute o projeto:**

```bash
go run main.go
```

> Certifique-se de ter o Go instalado: https://golang.org/dl/

## 📂 Estrutura

```
.
├── game/
│   ├── ai.go         # Lógica da IA (fácil, médio, difícil)
│   ├── game.go       # Controle principal do jogo (estado, eventos)
│   └── ui.go         # Desenho do tabuleiro e peças
├── main.go           # Ponto de entrada
├── go.mod
└── go.sum
```

## 🧠 Sobre a IA

- **Fácil:** movimentos aleatórios
- **Médio:** tenta ganhar ou bloquear
- **Difícil:** usa algoritmo Minimax

## 📸 Screenshot

![image](https://github.com/user-attachments/assets/b8125ff3-f9c2-47f0-bff4-62759afa5847)


## 📄 Licença

Este projeto é livre para fins educacionais ou pessoais. Modificações são bem-vindas.
