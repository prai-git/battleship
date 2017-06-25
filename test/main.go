package main

import "github.com/prai-git/battleship"

func main() {

	// MINUTE TIME INTERVAL
	const TIME_INTERVAL int64 = 1

	battleArea := battleship.BattleArea{}

	// Area Width

	message := "1 <= Width of Battle area (M’) <= 9  \n"
	battleArea.Width = battleship.GetNumberFromInputAs0to9(message)

	// Area Height
	message = "A <= Height of Battle area (N’) <= Z \n"
	battleArea.Height = battleship.GetNumberFromInputAsAtoZ(message)

	// Number of battleShips
	message = "1 <= Number of battleShips <= M’ * N’ \n"
	totalShips := battleship.GetNumberFromInputAs0to9(message)

	//  player 1
	player1 := battleship.Player{Name: "Player-1"}
	player1.Ships = make([]battleship.Ship, totalShips)
	player1.AddBattleBoard(battleArea)

	// player 2
	player2 := battleship.Player{Name: "Player-2"}
	player2.Ships = make([]battleship.Ship, totalShips)
	player2.AddBattleBoard(battleArea)

	for i := 0; i < totalShips; i++ {

		// Type of ship
		message = "Type of ship = {‘P’, ‘Q’} \n"
		sType := battleship.GetShipTypeFromInput(message)

		// ship Area
		shipArea := battleship.ShipArea{}

		message = "1 <= Width of battleship <= M’  \n"
		shipArea.Width = battleship.GetNumberFromInputAs0to9(message)

		message = "1 <= Height of battleship <= N’ \n"
		shipArea.Height = battleship.GetNumberFromInputAs0to9(message)

		// player 1 CellPosition
		cellPosition1 := battleship.CellPosition{}

		message = "A <= Y coordinate of ship <= N’ for Player 1  \n"
		cellPosition1.Y = battleship.GetNumberFromInputAsAtoZ(message)

		message = "1 <= X coordinate of ship <= M’ for Player 1 \n"
		cellPosition1.X = battleship.GetNumberFromInputAs0to9(message)

		// player 2 ship location
		cellPosition2 := battleship.CellPosition{}
		message = "A <= Y coordinate of ship <= N’ for Player 2  \n"
		cellPosition2.Y = battleship.GetNumberFromInputAsAtoZ(message)

		message = "1 <= X coordinate of ship  <= M’ for Player 2 \n"
		cellPosition2.X = battleship.GetNumberFromInputAs0to9(message)

		player1.AddShip(shipArea, cellPosition1, sType, i)
		player2.AddShip(shipArea, cellPosition2, sType, i)

	}

	// Now both players are ready lets play the game.
	battleship.Play(&player1, &player2, TIME_INTERVAL)
}
