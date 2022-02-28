package player

type Args struct {
	Top    int
	Bottom int
	Left   int
	Right  int
}

type Player interface {
	Play(file string, args Args)
	Stop()
}

type PlayCmd struct {
	File string
	Args Args
}

var StopCmd = PlayCmd{}

type PlayCmdLogger struct {
	Log []PlayCmd
}

func (p *PlayCmdLogger) Play(file string, args Args) {
	p.Log = append(p.Log, PlayCmd{file, args})
}

func (p *PlayCmdLogger) Stop() {
	p.Log = append(p.Log, StopCmd)
}
