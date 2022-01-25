package Castle

func (castle *Castle) CalcMove(P, Dir int) int {
	switch Dir {
	case 0:
		if P%3 == 2 {
			return -1
		}
		ExitID := 12*(P/9) + P%9 + 2*(P%9/3)
		if P%9 != 7 {
			ExitID += P % 3
		}
		if !castle.Exit[ExitID] {
			return -1
		}
		return P + 1
	case 1:
		if P%9 >= 6 {
			return -1
		}
		ExitID := 12*(P/9) + P%9 + 2*(P%9/3) + P%3
		if P%3 != 2 {
			ExitID += 1
		}
		if !castle.Exit[ExitID] {
			return -1
		}
		return P + 3
	case 2:
		if P%3 == 0 {
			return -1
		}
		Q := P - 1
		ExitID := 12*(Q/9) + Q%9 + 2*(Q%9/3)
		if Q%9 != 7 {
			ExitID += Q % 3
		}
		if !castle.Exit[ExitID] {
			return -1
		}
		return Q
	case 3:
		if P%9 < 3 {
			return -1
		}
		Q := P - 3
		ExitID := 12*(Q/9) + Q%9 + 2*(Q%9/3) + Q%3
		if Q%3 != 2 {
			ExitID += 1
		}
		if !castle.Exit[ExitID] {
			return -1
		}
		return Q
	case 4:
		if P >= 18 || !castle.Exit[36+P] {
			return -1
		} else {
			return P + 9
		}
	case 5:
		if P < 9 || !castle.Exit[36+P-9] {
			return -1
		} else {
			return P - 9
		}
	default:
		return -1
	}
}
