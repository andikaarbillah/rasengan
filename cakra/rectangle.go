package cakra

type Jurus struct {
	Ninjutsu int
	Jutsu    int
}

func (jurus *Jurus) Kyubi() int {
	return jurus.Ninjutsu - jurus.Jutsu
}
