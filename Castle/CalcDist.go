package Castle

import (
	"Adventure-go/Utils/Queue"
)

func (castle *Castle) CalcDist(Pos_S int, Dist *[27]int) {
	for i := 0; i < 27; i++ {
		Dist[i] = -1
	}
	Q := Queue.Queue{Pos_S}
	Dist[Pos_S] = 0
	var tmp interface{}
	for P, P_Nxt := 0, 0; !Q.Empty(); {
		tmp, _ = Q.Front()
		P = tmp.(int)
		Q.Pop()
		for i := 0; i < 6; i++ {
			P_Nxt = castle.CalcMove(P, i)
			if P_Nxt != -1 && Dist[P_Nxt] == -1 {
				Q.Push(P_Nxt)
				Dist[P_Nxt] = Dist[P] + 1
			}
		}
	}
}
