package Game

import (
	"Adventure-go/Castle"
)

var RoomName = [6]string{
	"Lobby",
	"Princess's Prison",
	"Monster's Lair",
	"Normal Room",
	"Radar Room",
	"Trigger Room"}
var Direction = [6]string{
	"east",
	"south",
	"west",
	"north",
	"up",
	"down"}

type Game struct {
	With_Princess, GameOver bool
	Pos_Player              int
	Castle.Castle
}
