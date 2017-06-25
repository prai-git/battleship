package battleship

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
