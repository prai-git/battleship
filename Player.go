package battleship

import (
	"errors"
	"fmt"
)

type Player struct {
	battleBoard BattleShipBoard
	Ships       []Ship
	Name        string
	chanInput   chan bool
	testData    []CellPosition
}

func (p Player) HitTarget(position CellPosition) (result MissileStatus) {
	if len(p.battleBoard.cells) < position.Y || len(p.battleBoard.cells[position.Y-1]) < position.X {
		return MISS
	}
	cell := &p.battleBoard.cells[position.Y-1][position.X-1]
	// fmt.Println(cell)
	if cell.status > SHINK_SHIP {
		cell.status--
		p.UpdateShip(&cell.position)
		return HIT
	}

	return MISS
}

/*
 UpdateShip -
*/
func (p *Player) UpdateShip(cellposition *CellPosition) {

	// fmt.Println(p)
	//fmt.Println(p.Ships)
	for index := range p.Ships {
		ship := &p.Ships[index]
		isShunk := true
		for i := range ship.cells {
			cPosition := &ship.cells[i]
			if p.battleBoard.cells[cPosition.Y][cPosition.X].status > 0 {
				isShunk = false
				break
			}
		}

		if isShunk == true {
			ship.status = SHIP_SHUNK
		} else {
			ship.status = SHIP_FLOAT
		}
	}
}

/**
 IsAllShunk -
**/
func (p *Player) IsAllShunk() (result bool) {

	for index := range p.Ships {
		ship := &p.Ships[index]

		if ship.status == SHIP_FLOAT {
			result = false
			return
		}
	}
	result = true
	return
}

func (p *Player) IsWinner(p2 *Player) (result bool) {
	if p2.IsAllShunk() {
		// fmt.Println(p.Name, "won the battle ")
		return true
	}
	result = false
	return
}

func (p *Player) AddShip(shipArea ShipArea,
	cellPosition CellPosition,
	sType ShipType, index int) {

	ship := Ship{cells: make([]CellPosition,
		shipArea.Width*shipArea.Height),
		shipType: sType,
		status:   SHIP_FLOAT}

	row := cellPosition.Y - 1
	column := cellPosition.X - 1
	for i := 0; i < shipArea.Height; i++ {
		for j := 0; j < shipArea.Width; j++ {
			ship.cells[i*shipArea.Width+j] = CellPosition{X: column, Y: row}
			//fmt.Println(row, ":", column)
			//fmt.Println(p.battleBoard)
			cell := &p.battleBoard.cells[row][column]
			if sType == Q {
				cell.status = Q_SHIP

			} else {
				cell.status = P_SHIP
			}

			column++

		}
		row++
	}
	p.Ships[index] = ship
}

func (p *Player) AddBattleBoard(battleArea BattleArea) {

	//fmt.Println(battleArea)

	p.battleBoard = BattleShipBoard{cells: make([][]Cell, battleArea.Height)}
	for row := 0; row < battleArea.Height; row++ {

		p.battleBoard.cells[row] = make([]Cell, battleArea.Width)
		for column := 0; column < battleArea.Height; column++ {
			cellPosition := CellPosition{X: column, Y: row}
			p.battleBoard.cells[row][column] = Cell{position: cellPosition, status: NO_SHIP}
		}
	}
}

func (p *Player) Play(missileLocationChan chan CellPosition, input []CellPosition) (err error) {
	index := 0

	for {
		select {
		case <-p.chanInput:
			var cellPosition CellPosition

			if input != nil && len(input) >= index {
				//fmt.Println("index : ", index)
				if index < len(input) {
					cellPosition = input[index]
				} else {
					fmt.Printf("%s does not have data to play \n", p.Name)
					err = errors.New("Insufficient input data")
					return
				}
				index++

			} else {

				cellPosition = CellPosition{}
				fmt.Print("\n Player ", p.Name, " is Playing \n")

				message := "A <=  Y coordinate of Target Ship <= N’ \n"
				//var inputY string
				cellPosition.Y, _ = GetYCordFromInputAsAtoZ(message)

				message = "1 <= X coordinate of Target ship <= M’  \n"
				//var inputX string
				cellPosition.X, _ = GetXCordFromInputAs0to9(message)
			}

			missileLocationChan <- cellPosition
		}
	}
}
