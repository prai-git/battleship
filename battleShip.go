package battleship

import (
	"fmt"
	"strings"
)

type BattleArea Area
type ShipArea Area
type Area struct {
	Width  int
	Height int
}

type CellPosition struct {
	X int
	Y int
}

type ShipType int
type CellStatus int
type ShipStatus int
type MissileStatus string

const (
	Q ShipType = 2
	P ShipType = 1
)

const (
	Q_SHIP     CellStatus = 2
	P_SHIP     CellStatus = 1
	SHINK_SHIP CellStatus = 0
	NO_SHIP    CellStatus = -1
)

const (
	HIT   MissileStatus = "hit"
	MISS  MissileStatus = "miss"
	SHINK MissileStatus = "shink"
)

const (
	SHIP_FLOAT     ShipStatus = 2
	SHIP_HALFFLOAT ShipStatus = 1
	SHIP_SHUNK     ShipStatus = 0
)

type Cell struct {
	position CellPosition
	status   CellStatus
}

type Ship struct {
	cells    []CellPosition
	status   ShipStatus
	shipType ShipType
}

type BattleShipBoard struct {
	cells [][]Cell
}

func Main(filePath string) (err error) {

	// MINUTE TIME INTERVAL
	const TIME_INTERVAL int64 = 1

	battleArea := BattleArea{}

	// Area Width

	lines, err := ReadLines(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// message := "1 <= Width of Battle area (M’) <= 9  \n"
	var line0 string = lines[0]
	battleAreas := strings.Split(line0, " ")
	battleArea.Width, _ = GetNumberFrom0to9(Trim(battleAreas[0]))

	// Area Height
	// message = "A <= Height of Battle area (N’) <= Z \n"
	battleArea.Height, _ = GetNumberFromAtoZ(Trim(battleAreas[1]))

	// Number of battleShips
	// message = "1 <= Number of battleShips <= M’ * N’ \n"
	line1 := lines[1]
	totalShips, _ := GetNumberFrom0to9(Trim(line1))

	//  player 1
	player1 := Player{Name: "Player-1"}
	player1.Ships = make([]Ship, totalShips)
	player1.AddBattleBoard(battleArea)

	// player 2
	player2 := Player{Name: "Player-2"}
	player2.Ships = make([]Ship, totalShips)
	player2.AddBattleBoard(battleArea)

	shipLine := 2
	for i := 0; i < totalShips; i++ {

		shipLine := lines[shipLine+i]
		shipInput := strings.Split(string(shipLine), " ")
		// Type of ship
		// message = "Type of ship = {‘P’, ‘Q’} \n"
		sType, _ := GetShipType(Trim(shipInput[0]))

		// ship Area
		shipArea := ShipArea{}

		// message = "1 <= Width of battleship <= M’  \n"
		shipArea.Width, _ = GetNumberFrom0to9(Trim(shipInput[1]))

		// message = "1 <= Height of battleship <= N’ \n"
		shipArea.Height, _ = GetNumberFrom0to9(Trim(shipInput[2]))

		// player 1 CellPosition
		cellPosition1 := CellPosition{}

		// message = "A <= Y coordinate of ship <= N’ for Player 1  \n"
		cellPosition1.Y, _ = GetNumberFromAtoZ(string([]rune(Trim(shipInput[3]))[0]))

		// message = "1 <= X coordinate of ship <= M’ for Player 1 \n"
		cellPosition1.X, _ = GetNumberFrom0to9(string([]rune(Trim(shipInput[3]))[1]))

		// player 2 ship location
		cellPosition2 := CellPosition{}
		//message = "A <= Y coordinate of ship <= N’ for Player 2  \n"
		cellPosition2.Y, _ = GetNumberFromAtoZ(string([]rune(Trim(shipInput[4]))[0]))

		//message = "1 <= X coordinate of ship  <= M’ for Player 2 \n"
		cellPosition2.X, _ = GetNumberFrom0to9(string([]rune(Trim(shipInput[4]))[1]))

		player1.AddShip(shipArea, cellPosition1, sType, i)
		player2.AddShip(shipArea, cellPosition2, sType, i)

	}

	line4 := lines[4]
	line5 := lines[5]

	positionsInput1 := strings.Split(string(line4), " ")
	positionsInput2 := strings.Split(string(line5), " ")

	input1 := make([]CellPosition, len(positionsInput1))
	input2 := make([]CellPosition, len(positionsInput2))

	for i := 0; i < len(positionsInput1); i++ {
		cellPosition := CellPosition{}
		cellPosition.Y, _ = GetNumberFromAtoZ(string([]rune(Trim(positionsInput1[i]))[0]))
		cellPosition.X, _ = GetNumberFrom0to9(string([]rune(Trim(positionsInput1[i]))[1]))
		input1[i] = cellPosition
	}

	for j := 0; j < len(positionsInput2); j++ {
		cellPosition := CellPosition{}
		cellPosition.Y, _ = GetNumberFromAtoZ(string([]rune(Trim(positionsInput2[j]))[0]))
		cellPosition.X, _ = GetNumberFrom0to9(string([]rune(Trim(positionsInput2[j]))[1]))
		input2[j] = cellPosition
	}

	// Now both players are ready lets play the game.
	err = Play(&player1, &player2, input1, input2, TIME_INTERVAL)
	return

}
