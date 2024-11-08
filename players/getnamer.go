package players
type GetNamer struct {
	Name string
}

func (g GetNamer) GetName() string {
	return g.Name
}


