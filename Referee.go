package battleship

import (
	"fmt"
	"time"
)

func Play(p1 *Player, p2 *Player, minutTime int64) {
	playing := p1
	nonPlaying := p2

	// set the timer
	timer := time.NewTimer(time.Minute * time.Duration(minutTime))

	missileLocationChan := make(chan CellPosition)
	playing.chanInput = make(chan bool)
	nonPlaying.chanInput = make(chan bool)

	go playing.Play(missileLocationChan)
	go nonPlaying.Play(missileLocationChan)

	// inform player to hit the target
	playing.chanInput <- true
	for {
		select {
		case cellPosition := <-missileLocationChan:
			// fmt.Println(nonPlaying)

			missileResult := nonPlaying.HitTarget(cellPosition)
			fmt.Print(playing.Name, " fires a missile with target ", cellPosition.Y, cellPosition.X, " which got ", missileResult)

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
			return

		}
	}
}
