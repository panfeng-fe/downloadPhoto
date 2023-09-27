package tools

type Glimit struct {
	max int
	c   chan struct{}
}

func (g *Glimit) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}

func Workers(max int) *Glimit {
	return &Glimit{
		max: max,
		c:   make(chan struct{}, max),
	}
}
