package Game

import (
	"fmt"
)

func (game *Game) PrintRoom() {
	fmt.Println("Welcome to the " + RoomName[game.RoomType[game.Pos_Player]] + ".")
	switch game.RoomType[game.Pos_Player] {
	case 0:
		if game.With_Princess {
			fmt.Println("Congratulations, you have successfully taken the princess away from the castle!")
			game.GameOver = true
			return
		}
	case 1:
		if !game.With_Princess {
			fmt.Println("You have found princess, and she is going to leave with you. Please take her to the Lobby.")
			game.With_Princess = true
		}
	case 2:
		fmt.Println("Unfortunately, you have been killed by the monster. Game Over!")
		game.GameOver = true
		return
	case 4:
		fmt.Println("Welcome to Radar Room! You will get some data about distance:")
		var dist [27]int
		game.CalcDist(game.Pos_Player, &dist)
		fmt.Println("0. There are at least", dist[game.Pos_Lobby], "step(s) for you to reach the Lobby.")
		fmt.Println("1. There are at least", dist[game.Pos_Princess], "step(s) for you to reach the Princess's Prison.")
		fmt.Println("2. There are at least", dist[game.Pos_Monster], "step(s) for you to reach the Monster's Lair.")
	case 5:
		fmt.Println("Welcome to Trigger Room! Exits of all rooms have been reset randomly, so you need to look for a new path!")
		game.SetExit(game.Pos_Player)
	}
	game.PrintExit()
}
