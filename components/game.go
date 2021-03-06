package components

import (
	"fmt"
	"time"
)

//LifeGame ライフゲーム
type LifeGame struct {
	currentField   *Field
	lastField      *Field
	intervalSecond int
	useRoutine     bool
	isPrint        bool
}

//NewLifeGame ゲームを新規作成
func NewLifeGame(height int, width int, initRate float64, interval int, useRoutine bool, isPrint bool) *LifeGame {
	lifeGame := new(LifeGame)
	gamefield := CreateFieldFrame(height, width).InitFieldStatus(initRate)

	lifeGame.currentField = gamefield
	lifeGame.intervalSecond = interval
	lifeGame.isPrint = isPrint
	lifeGame.useRoutine = useRoutine
	return lifeGame
}

func (game *LifeGame) computeNextFlameAsync() (*Field, *Field) {
	type routineParams struct {
		h         int
		nextCells []int
	}
	channel := make(chan routineParams, game.currentField.width)
	defer close(channel)
	for h := 0; h < game.currentField.height; h++ {
		go func(_h int, _field *Field) {
			nextCells := make([]int, _field.width)
			for _w := 0; _w < _field.width; _w++ {
				AroundLifeCount := _field.countAroundLife(_h, _w)
				var nextCell int
				if game.currentField.status[_h][_w] == 1 {
					if AroundLifeCount == 2 || AroundLifeCount == 3 {
						nextCell = 1
					} else {
						nextCell = 0
					}
				} else if game.currentField.status[_h][_w] == 0 {
					if AroundLifeCount == 3 {
						nextCell = 1
					} else {
						nextCell = 0
					}
				} else {
					fmt.Println("エラーが発生しました")
				}
				nextCells[_w] = nextCell
			}
			channel <- routineParams{_h, nextCells}
		}(h, game.currentField)
	}
	nextFrame := CreateFieldFrame(game.currentField.height, game.currentField.width)
	for _h := 0; _h < game.currentField.height; _h++ {
		tmp := <-channel
		nextFrame.status[tmp.h] = tmp.nextCells
	}
	return nextFrame, game.currentField
}

func (game *LifeGame) nextFrame() (*Field, *Field) {
	nextFrame := CreateFieldFrame(game.currentField.height, game.currentField.width)
	for h := 0; h < nextFrame.height; h++ {
		for w := 0; w < nextFrame.width; w++ {
			AroundLifeCount := game.currentField.countAroundLife(h, w)
			var nextCell int
			if game.currentField.status[h][w] == 1 {
				if AroundLifeCount == 2 || AroundLifeCount == 3 {
					nextCell = 1
				} else {
					nextCell = 0
				}
			} else if game.currentField.status[h][w] == 0 {
				if AroundLifeCount == 3 {
					nextCell = 1
				} else {
					nextCell = 0
				}
			} else {
				fmt.Println("エラーが発生しました")
			}
			nextFrame.status[h][w] = nextCell

		}
	}
	return nextFrame, game.currentField
}

func (game *LifeGame) isChange(lastField *Field) bool {
	for hIndex, line := range game.currentField.status {
		for wIndex, cell := range line {
			if cell == lastField.status[hIndex][wIndex] {
				continue
			} else {
				return true
			}
		}
	}
	return false
}

//MainLoop ゲームのメインのループ
func (game *LifeGame) MainLoop() {
	i := 1
	var timeSum time.Duration
	for {
		if game.useRoutine {
			fmt.Print("** USING ROUTINE **")
		}
		fmt.Println("step", i)
		if game.isPrint {
			game.currentField.printField()
		}
		startTime := time.Now()
		if game.useRoutine {
			game.currentField, game.lastField = game.computeNextFlameAsync()
		} else {
			game.currentField, game.lastField = game.nextFrame()
		}
		endTime := time.Now()
		timeSum += endTime.Sub(startTime)
		fmt.Printf("time duration for  next flame computing %s \n", endTime.Sub(startTime))
		fmt.Printf("AVG time duration for  next flame computing %s \n", time.Duration(float64(timeSum)/float64(i))*time.Nanosecond)
		time.Sleep(time.Duration(game.intervalSecond) * time.Second)
		if !game.isChange(game.lastField) {
			break
		}
		if game.isPrint {
			fmt.Printf("\033[%dA", game.currentField.height+3)
		}
		i++
	}
}
