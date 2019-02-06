package part

import (
  "github.com/gen2brain/raylib-go/raylib"
)

func (self *Particle) debugVel() {
	rl.DrawLineEx(rl.NewVector2(float32(self.Body.Position.X), float32(self.Body.Position.Y)), rl.NewVector2(float32(self.Body.Position.X+self.Body.Velocity.X), float32(self.Body.Position.Y+self.Body.Velocity.Y)), 2, rl.Green)
}

func (self *Particle) debugForce() {
	rl.DrawLineEx(rl.NewVector2(float32(self.Body.Position.X), float32(self.Body.Position.Y)), rl.NewVector2(float32(self.Body.Position.X+self.Body.Force.X), float32(self.Body.Position.Y+self.Body.Force.Y)), 2, rl.Red)
}
