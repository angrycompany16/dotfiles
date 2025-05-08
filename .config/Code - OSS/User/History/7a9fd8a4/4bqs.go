package overlay

import (
	"image/color"
	"mask_of_the_tomb/internal/game/core/rendering"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var fontColor = color.RGBA{255, 0, 0, 255}

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
	opText.ColorScale.ScaleWithColor(t.Color.DarkColor)
	if t.ScreenAlign == screenCentered {
		opText.GeoM.Translate(parentWidth/2, parentHeight/2)
	}

	text.Draw(rendering.RenderLayers.UI,
		t.Text,
		&text.GoTextFace{
			Source: t.Font.GoTextFaceSource,
			Size:   t.FontSize,
		}, opText)

	opText.ColorScale = ebiten.ColorScale{}
	opText.ColorScale.ScaleWithColor(t.Color.BrightColor)

	opText.GeoM.Translate(-t.ShadowX, t.ShadowY)
	text.Draw(rendering.RenderLayers.UI, t.Text, &text.GoTextFace{
		Source: t.Font.GoTextFaceSource,
		Size:   t.FontSize,
	}, opText)
}
