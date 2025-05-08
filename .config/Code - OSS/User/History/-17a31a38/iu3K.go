package node

type Dialogue struct {
	Textbox     `yaml:",inline"`
	RevealSpeed float64 `yaml:"RevealSpeed"`
}

func (d *Dialogue) Update(confirmations map[string]bool) {

}

func (d *Dialogue) Draw() {

}

func (d *Dialogue) Reset() {

}
