package overlay

import (
	"image/color"
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
	opText.ColorScale.ScaleWithColor(color.RGBA{textColor[0], textColor[1], textColor[2], uint8(t * 255)})
	opText.GeoM.Translate(rendering.GameWidth/2, rendering.GameHeight/2)

	text.Draw(rendering.RenderLayers.UI,
		tc.text,
		&text.GoTextFace{
			Source: tc.font,
			Size:   32,
		}, opText)
}

func (tc *TitleCard) ChangeText(text string) {
	tc.text = text
}

func NewTitleCard(text string) OverlayContent {
	return &TitleCard{
		text:  text,
		font:  fonts.GetFont("JSE_AmigaAMOS"),
		image: ebiten.NewImage(rendering.GameWidth, rendering.GameHeight),
	}
}
