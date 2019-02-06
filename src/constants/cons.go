package constants

import (
  "github.com/gen2brain/raylib-go/raylib"
)

const (
	// App
  SCREEN_W float32 = 1000
  SCREEN_H float32 = 800

	// Universal
	ChrMult = 8990000000 // 1/4piE0 =~ 8.99 x 10^9

	// Proton
	ProtonR float32 = 5
	ProtonChr float32 = 0.00016
	ProtonM float32 = 0.01
)

var (
	// Proton
	ProtonDiam float32 = 2*ProtonR
	ProtonClr rl.Color = rl.Maroon

	AntiProtonClr rl.Color = rl.SkyBlue
)
