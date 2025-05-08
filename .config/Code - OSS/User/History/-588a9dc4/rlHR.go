package overlay

type Overlay interface {
	StartEnter()
	StartExit()
	Update()
	Draw()
}
