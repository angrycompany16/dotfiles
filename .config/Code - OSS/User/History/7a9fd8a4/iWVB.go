package overlay

import (
	"mask_of_the_tomb/internal/game/UI/fonts"
	"mask_of_the_tomb/internal/game/core/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var textColor = []uint8{255, 0, 0}

type TitleCard struct {
	text  string
	font  *text.GoTextFaceSource
	image *ebiten.Image
}

func (tc *TitleCard) Draw(t float64) {
	opText := &text.DrawOptions{}
	opText.LayoutOptions.LineSpacing = 40
	opText.LayoutOptions.PrimaryAlign = text.AlignCenter
	opText.LayoutOptions.SecondaryAlign = text.AlignCenter
	opText.ColorScale = ebiten.ColorScale{}
	opText.ColorScale.ScaleWithColor(textColor)
	opText.GeoM.Translate(rendering.GameWidth/2, rendering.GameHeight/2)

	text.Draw(rendering.RenderLayers.UI,
		tc.text,
		&text.GoTextFace{
			Source: tc.font,
			Size:   32,
		}, opText)
}

func NewTitleCard(text string) OverlayContent {
	return &TitleCard{
		text:  text,
		font:  fonts.GetFont("JSE_AmigaAMOS"),
		image: ebiten.NewImage(rendering.GameWidth, rendering.GameHeight),
	}
}
