package Castle

import (
	"Adventure-go/Utils/Random"
)

func (castle *Castle) NextRandRoom(Limit int) int {
	Res := Random.RandInt(0, Limit)
	for castle.RoomType[Res] != 3 {
		Res = Random.RandInt(0, Limit)
	}
	return Res
}

func (castle *Castle) NextRandExit(L, Len int, op bool) int {
	Res := Random.RandInt(L, Len)
	for castle.Exit[Res] != op {
		Res = Random.RandInt(L, Len)
	}
	return Res
}
