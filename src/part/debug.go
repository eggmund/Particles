package part

import (
	"constants"
	"github.com/gen2brain/raylib-go/raylib"
)

func (self *Particle) debugVel() {
	rl.DrawLineEx(rl.NewVector2(float32(self.Body.Position.X), float32(self.Body.Position.Y)), rl.NewVector2(float32(self.Body.Position.X+(self.Body.Velocity.X/(constants.SCALE*constants.SCALE*2))), float32(self.Body.Position.Y+(self.Body.Velocity.Y/(constants.SCALE*constants.SCALE*2)))), 2, rl.Green)
}

func (self *Particle) debugForce() {
	rl.DrawLineEx(rl.NewVector2(float32(self.Body.Position.X), float32(self.Body.Position.Y)), rl.NewVector2(float32(self.Body.Position.X+(self.Body.Force.X/(constants.SCALE*constants.SCALE*10000000000000000000000))), float32(self.Body.Position.Y+(self.Body.Force.Y/(constants.SCALE*constants.SCALE*10000000000000000000000)))), 2, rl.Red)
}
