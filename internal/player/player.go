package player

type Args struct {
	Top    int
	Bottom int
	Left   int
	Right  int
}

type Player interface {
	Play(file string, args Args)
}

type PlayCmd struct {
	File string
	Args Args
}

type PlayCmdLogger struct {
	Log []PlayCmd
}

func (p *PlayCmdLogger) Play(file string, args Args) {
	p.Log = append(p.Log, PlayCmd{file, args})
}
