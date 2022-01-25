package Castle

import (
	"Adventure-go/Utils/Queue"
)

func (castle *Castle) CheckExit(Start_Point int) bool {
	var Vis [27]bool
	for i := 0; i < 27; i++ {
		Vis[i] = false
	}
	Q := Queue.Queue{Start_Point}
	Vis[Start_Point] = true
	var tmp interface{}
	for P, P_Nxt := 0, 0; !Q.Empty(); {
		tmp, _ = Q.Front()
		P = tmp.(int)
		Q.Pop()
		for i := 0; i < 6; i++ {
			P_Nxt = castle.CalcMove(P, i)
			if P_Nxt != -1 && castle.RoomType[P_Nxt] != 2 && !Vis[P_Nxt] {
				Q.Push(P_Nxt)
				Vis[P_Nxt] = true
			}
		}
	}
	return Vis[castle.Pos_Princess] && Vis[castle.Pos_Lobby]
}
