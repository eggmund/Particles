package constants

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

const (
	// App
	SCREEN_W float32 = 1000
	SCREEN_H float32 = 800

	// Universal
	ChrMult = 8990000000 // 1/4piE0 =~ 8.99 x 10^9
)

var (
	sclFactor float64 = 3
	SCALE     float64 = sclFactor / (float64(0.8751) * math.Pow10(-15)) // Radius of proton on screen / actuall radius

	// Proton
	ProtonR    float64  = sclFactor
	ProtonChr  float64  = (1.6 * math.Pow10(-16)) * SCALE
	ProtonM    float64  = (1.67 * math.Pow10(-27)) * SCALE
	ProtonDiam float64  = 2 * ProtonR
	ProtonClr  rl.Color = rl.Maroon

	AntiProtonClr rl.Color = rl.SkyBlue
)
