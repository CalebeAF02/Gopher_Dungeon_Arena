package goguide

// Player é um exemplo de struct que agrupa campos relacionados.
type Player struct {
	Name   string
	Health int
	Level  int
}

// NewPlayer cria um novo Player com valores iniciais padrão.
func NewPlayer(name string) Player {
	return Player{Name: name, Health: 100, Level: 1}
}

// LevelUp aumenta o nível do jogador.
func (p *Player) LevelUp() {
	p.Level++
}
