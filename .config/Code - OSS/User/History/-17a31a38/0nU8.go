package node

type Dialogue struct {
	Textbox     `yaml:",inline"`
	RevealSpeed float64 `yaml:"RevealSpeed"`
}

func (d *Dialogue) Update(confirmations map[string]bool) {

}

func (d *Dialogue) Draw(offsetX, offsetY float64, parentWidth, parentHeight float64) {
	w, h := inheritSize(d.Width, d.Height, parentWidth, parentHeight)
	d.DrawChildren(offsetX+d.PosX, offsetY+d.PosY, w, h)
}

func (d *Dialogue) Reset() {

}
