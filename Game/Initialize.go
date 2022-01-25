package Game

func (game *Game) Initialize() {
	game.SetRoomType()
	game.SetExit(game.Pos_Lobby)
	game.With_Princess = false
	game.GameOver = false
	game.Pos_Player = game.Pos_Lobby
}
