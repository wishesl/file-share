package assets

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type MyDefaultTheme struct{}

var _ fyne.Theme = (*MyDefaultTheme)(nil)

//go:embed simkai.ttf
var arialNovaLight []byte

//go:embed logo.png
var logoData []byte

//go:embed share.png
var shareData []byte

//go:embed uit--twitter-alt.png
var twitterData []byte

var TwitterDataSR = &fyne.StaticResource{
	StaticName:    "uit--twitter-alt.png",
	StaticContent: twitterData,
}

//go:embed uit--multiply.png
var multiplyData []byte
var MultiplyDataSR = &fyne.StaticResource{
	StaticName:    "uit--multiply.png",
	StaticContent: multiplyData,
}

var LogoDataSR = &fyne.StaticResource{
	StaticName:    "logo.png",
	StaticContent: logoData,
}

var ShareDataSR = &fyne.StaticResource{
	StaticName:    "share.png",
	StaticContent: shareData,
}

var arialNovaLightSR = &fyne.StaticResource{
	StaticName:    "simkai.ttf",
	StaticContent: arialNovaLight,
}

func (*MyDefaultTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	return arialNovaLightSR
}

func (*MyDefaultTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameButton:
		return color.Transparent
	default:
		return theme.DefaultTheme().Color(n, v)
	}
}

func (*MyDefaultTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*MyDefaultTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
