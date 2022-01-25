// Adventure-go project main.go
package main

import (
	"Adventure-go/Game"
	"Adventure-go/Utils/IO"
)

func main() {
	var Game_Inst Game.Game
	Game_Inst.Initialize()
	Game_Inst.Interact()
	IO.PressToExit()
}
