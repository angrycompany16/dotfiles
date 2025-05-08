package node

type Dialogue struct {
	Textbox         `yaml:",inline"`
	RevealTime      float64 `yaml:"RevealSpeed"`
	t               float64
	revealIndicator int
	revealedText    string
}

// This is cursed but we will receive the input inside the UI node
func (d *Dialogue) Update(confirmations map[string]bool) {
	d.UpdateChildren(confirmations)
	d.t += 1 / 60
	// if confirm key pressed
	// finish or print more text

	// else timer
}

func (d *Dialogue) Draw(offsetX, offsetY float64, parentWidth, parentHeight float64) {
	w, h := inheritSize(d.Width, d.Height, parentWidth, parentHeight)
	d.DrawChildren(offsetX+d.PosX, offsetY+d.PosY, w, h)
}

func (d *Dialogue) Reset() {
	d.ResetChildren()
}
