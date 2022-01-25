# Adventure-go
Adventure is a CLI game, which is a project of the OOP course of ZJU. This is the go version of [ZJU-OOP-Adventure](https://github.com/CNHowardHu/ZJU-OOP-Adventure).

## Compile & Run

```shell
$ git clone https://github.com/CNHowardHu/Adventure-go.git
$ cd Adventure-go && go build -trimpath -ldflags "-w -s"
$ ./Adventure-go
```

## My Design

Adventure is a CLI game. The player has to explore in the castle with many levels and a lot of rooms. The task of the player is to find a  room where the princess is prisoned and take her leave the castle. There are many types of rooms, and each type of room has different types of  exits. Note that there's a monster in one of the rooms, which the exact  location is not able to be aware. But once the player meets a monster,  the game is over.

There are 6 kinds of rooms in the castle:

|       Type        | Quantity |
| :---------------: | :------: |
|       Lobby       |    1     |
| Princess's Prison |    1     |
|  Monster's Lair   |    1     |
|    Normal Room    |    18    |
|    Radar Room     |    4     |
|   Trigger Room    |    2     |

When the game starts, the player is in the lobby of the castle. Then  the program shows information about the lobby：name of the room, how many exits are there, and names of all exits (e.g.: "east", "south", "up"),  like:

```
Welcome to the lobby. There are 3 exits: east, west and up.
Enter your command:
```

The player then can input "go" followed by the name of one exit to enter the room connected to that door, like:

```
go east
```

The player goes into the room to the east. The program shows the  information about that room, like what happened in the lobby just now.  And the player may input command to choose another room.

Once the player enters a room with a monster, the program shows a  message and game over. Once the player enters the room of princess, the  program shows a message about the princess, and the princess is going to leave with the player. The player then has to find their way out the  castle. The only way to leave the castle is via the lobby.

Special rooms designed by me:

- **Radar Room** can tell you the minimum number of steps to Lobby, Princess's Prison or Monster's Lair.

- **Trigger Room** can randomly reset Exits of all rooms.

The map of this game is a cube with 3 rooms on each side, and the total number of rooms is 27. There are 12-5=7 exits in each floors and 2 exits between each two floors.

The room with monster or princess is randomly set. It's guaranteed that there exist at least one path to find Princess while avoiding Monster at any time, even if the Trigger Room has been activated.
