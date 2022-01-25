package Game

import (
	"fmt"
)

func (game *Game) PrintExit() {
	cnt := 0
	var ExitDir [6]int
	for i := 0; i < 6; i++ {
		if game.CalcMove(game.Pos_Player, i) != -1 {
			ExitDir[cnt] = i
			cnt++
		}
	}
	if cnt == 1 {
		fmt.Println("There is", cnt, "exit: "+Direction[ExitDir[0]]+".")
	} else {
		fmt.Print("There is ", cnt, " exit: "+Direction[ExitDir[0]])
		for i := 1; i <= cnt-2; i++ {
			fmt.Print(", " + Direction[ExitDir[i]])
		}
		fmt.Println(" and " + Direction[ExitDir[cnt-1]] + ".")
	}
}
