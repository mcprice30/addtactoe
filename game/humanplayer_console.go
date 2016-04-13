package game

import (
    "fmt"
)

type ConsoleHumanPlayer struct {
    playerNumber int
    name string
}

func NewConsoleHumanPlayer(playerNumber int, name string) ConsoleHumanPlayer {
    player := ConsoleHumanPlayer{}
    player.playerNumber = playerNumber
    player.name = name
    return player
}

func (_ ConsoleHumanPlayer) ClaimCellMove(board Board) (int, int) {
    var row, col int
    fmt.Print("Enter row: ")
    _, err1 := fmt.Scanf("%d", &row)
    fmt.Print("Enter col: ")
    _, err2 := fmt.Scanf("%d", &col)

    if err1 != nil || err2 != nil {
        row = 0
        col = 0
    }

    return row-1, col-1
}

func (_ ConsoleHumanPlayer) NewMiniBoardMove(board Board) (int, int) {
    var row, col int
    fmt.Print("Enter row: ")
    _, err1 := fmt.Scanf("%d", &row)
    fmt.Print("Enter col: ")
    _, err2 := fmt.Scanf("%d", &col)

    if err1 != nil || err2 != nil {
        row = 0
        col = 0
    }

    return row-1, col-1

}

func (p ConsoleHumanPlayer) PlayerNumber() int {
    return p.playerNumber
}
