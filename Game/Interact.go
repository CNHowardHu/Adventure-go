package Game

import (
	"errors"
	"fmt"
)

func (game *Game) Interact() {
	game.PrintRoom()
	for !game.GameOver {
		fmt.Println("Enter your command:")
		if Dir, err := InputMove(); err != nil {
			fmt.Println("Your Command is wrong!")
		} else {
			if NextPos := game.CalcMove(game.Pos_Player, Dir); NextPos == -1 {
				fmt.Println("Your command is wrong!")
			} else {
				game.Pos_Player = NextPos
				game.PrintRoom()
			}
		}
	}
}

func InputMove() (int, error) {
	var op, dir string
	fmt.Scanln(&op, &dir)
	if op == "go" {
		for i := 0; i < 6; i++ {
			if dir == Direction[i] {
				return i, nil
			}
		}
	}
	return -1, errors.New("Your command is wrong!")
}
