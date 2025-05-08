package node

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Dialogue struct {
	Textbox         `yaml:",inline"`
	Name            string   `yaml:"Name"`
	RevealTime      float64  `yaml:"RevealSpeed"`
	Lines           []string `yaml:"Lines"`
	activeLine      string
	t               float64
	revealIndicator int
}

// This is cursed but we will receive the input inside the UI node
func (d *Dialogue) Update(confirmations map[string]bool) {
	if d.activeLine == "" {
		d.activeLine = d.Lines[0]
	}
	d.UpdateChildren(confirmations)
	d.t += 1.0 / 60.0
	if d.t > d.RevealTime {
		if d.revealIndicator == len(d.activeLine) {
			return
		}

		d.Text = strings.Join([]string{
			d.Text, string(d.activeLine[d.revealIndicator]),
		}, "")
		d.t = 0
		d.revealIndicator += 1
	}

	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		if d.revealIndicator == len(d.fullText) {
			confirmations[d.Name] = true
		} else {
			d.revealIndicator = len(d.fullText)
		}
	}
}

func (d *Dialogue) Draw(offsetX, offsetY float64, parentWidth, parentHeight float64) {
	w, h := inheritSize(d.Width, d.Height, parentWidth, parentHeight)
	d.DrawChildren(offsetX+d.PosX, offsetY+d.PosY, w, h)
	d.Textbox.Draw(offsetX, offsetY, parentWidth, parentHeight)
}

func (d *Dialogue) Reset() {
	d.ResetChildren()
}
