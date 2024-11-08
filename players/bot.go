package players

type Bot struct {
	GetNamer
}

func (bot Bot) PlayPiece() string {
	return "dop"
}
