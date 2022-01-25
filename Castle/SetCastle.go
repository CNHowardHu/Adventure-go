package Castle

func (castle *Castle) SetRoomType() {

	for i := 0; i < 27; i++ {
		castle.RoomType[i] = 3
	}
	castle.Pos_Lobby = castle.NextRandRoom(9)
	castle.RoomType[castle.Pos_Lobby] = 0
	castle.Pos_Princess = castle.NextRandRoom(27)
	castle.RoomType[castle.Pos_Princess] = 1
	castle.Pos_Monster = castle.NextRandRoom(27)
	castle.RoomType[castle.Pos_Monster] = 2
	for i := 0; i < 4; i++ {
		castle.RoomType[castle.NextRandRoom(27)] = 4
	}
	for i := 0; i < 2; i++ {
		castle.RoomType[castle.NextRandRoom(27)] = 5
	}
}

func (castle *Castle) SetExit(Start_Point int) {
	for true {
		for i := 0; i < 36; i++ {
			castle.Exit[i] = true
		}
		for i := 36; i < 54; i++ {
			castle.Exit[i] = false
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 5; j++ {
				castle.Exit[castle.NextRandExit(i*12, 12, true)] = false
			}
		}
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				castle.Exit[castle.NextRandExit(36+i*9, 9, false)] = true
			}
		}
		if castle.CheckExit(Start_Point) {
			break
		}
	}
}
