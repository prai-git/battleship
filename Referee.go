package battleship

import (
	"errors"
	"fmt"
	"time"
)

func Play(p1 *Player, p2 *Player, input1 []CellPosition, input2 []CellPosition, minutTime int64) (err error) {
	playing := p1
	nonPlaying := p2

	// set the timer
	timer := time.NewTimer(time.Second * time.Duration(minutTime))

	missileLocationChan := make(chan CellPosition)
	playing.chanInput = make(chan bool)
	nonPlaying.chanInput = make(chan bool)

	go playing.Play(missileLocationChan, input1)
	go nonPlaying.Play(missileLocationChan, input2)

	// inform player to hit the target
	playing.chanInput <- true
	for {
		select {
		case cellPosition := <-missileLocationChan:
			// fmt.Println(nonPlaying)

			missileResult := nonPlaying.HitTarget(cellPosition)
			fmt.Printf("%s fires a missile with target %c%d which got %s \n", playing.Name, convertNumberToChar(cellPosition.Y), cellPosition.X, missileResult)

			if missileResult == HIT {
				if nonPlaying.IsAllShunk() {
					fmt.Println("\n", playing.Name, "won the battle ")
					return
				}
			} else {
				// now Player will get the chance
				temp := playing
				playing = nonPlaying
				nonPlaying = temp
			}

			// inform player to hit the target
			playing.chanInput <- true

			// listen the timer
		case <-timer.C:
			fmt.Println("Oh! Time out, Both player has played nice game. Thank you")
			err = errors.New("Time out Error")
			return

		}
	}
}
