package player

type Args struct {
	Top    int
	Bottom int
	Left   int
	Right  int
}

type Player interface {
	play(file string, args Args)
}

type PlayCmd struct {
	file string
	args Args
}

type PlayCmdLogger struct {
	log []PlayCmd
}

func (p PlayCmdLogger) play(file string, args Args) {
	p.log = append(p.log, PlayCmd{file, args})
}

func Logger() *Player {
	var p Player
	p = PlayCmdLogger{}
	return &p
}
