package game



import "fmt"

const TiedGame = 5

type Game struct {
    board Board
    players [2]Player
    turnCount int
}

func NewGame(p1, p2 Player) Game {
    game := Game{}
    game.board = NewBoard()
    game.players[0] = p1
    game.players[1] = p2
    game.turnCount = 1
    return game
}

func (game Game) Play() {
    newBoard := false

    for game.board.Winner() == 0 {
        playerIdx := (game.turnCount + 1) % 2
        player := game.players[playerIdx]

        fmt.Println("Turn number", game.turnCount)

        if (newBoard) {
            game.board.FindNextBoards()
            fmt.Print(game.board.ToString())
            fmt.Println("New Board!")
            for game.board.SetNextBoard(player.NewMiniBoardMove(game.board)) != nil {
                fmt.Println("Invalid Move")
            }
            newBoard = false
        } else {
            fmt.Print(game.board.ToString())
            for true {
                playerRow, playerCol := player.ClaimCellMove(game.board)
                nb, err := game.board.ClaimCellMove(playerRow, playerCol, playerIdx + 1)
                if err == nil {
                    newBoard = nb
                    break
                } else {
                    fmt.Println("Invalid Move")
                }
            }
        }
        game.turnCount++
    }
    fmt.Print(game.board.ToString())
    winner := game.board.Winner()
    if winner == TiedGame {
        fmt.Println("Tie!")
    } else {
        fmt.Println("Congrats to: ", game.board.Winner())
    }
}
