package game

import (
    "errors"
)

const BoardDim = 7
const WinLen = 4
const CellDim = BoardDim * MiniBoardSize

const ActiveBoard = 1
const UnplayedBoard = 2
const FinishedBoard = 3
const PotentialBoard = 4


type Board struct {
	gameBoard            [BoardDim][BoardDim]MiniBoard
	activeRow, activeCol int
	cells                [BoardDim * MiniBoardSize][BoardDim * MiniBoardSize]int
}

func NewBoard() Board {
	board := Board{}
	for i, row := range board.gameBoard {
		for j := range row {
			board.gameBoard[i][j] = NewMiniBoard()
		}
	}
	board.activeRow = BoardDim / 2
	board.activeCol = BoardDim / 2
	board.gameBoard[board.activeRow][board.activeCol].status = ActiveBoard

	for i := 0; i < BoardDim*MiniBoardSize; i++ {
		for j := 0; j < BoardDim*MiniBoardSize; j++ {
			miniRow := i / MiniBoardSize
			miniCol := j / MiniBoardSize
			subRow := i % MiniBoardSize
			subCol := j % MiniBoardSize
			board.cells[i][j] = board.gameBoard[miniRow][miniCol].tiles[subRow][subCol]
		}
	}

	return board
}

func (board Board) ToString() string {
	out := ""
	for i := 0; i < BoardDim*MiniBoardSize; i++ {
		for j := 0; j < BoardDim; j++ {
			rowString, _ := board.gameBoard[i/MiniBoardSize][j].RowString(i % MiniBoardSize)
			out += rowString
		}
		out += "\n"
	}
	return out
}

func (board *Board) ClaimCellMove(row, col, player int) (bool, error) {
    newBoard, err := board.gameBoard[board.activeRow][board.activeCol].MakeMove(row, col, player)
    if err != nil {
        return false, err
    } else {
        cellRow := board.activeRow * MiniBoardSize + row
        cellCol := board.activeCol * MiniBoardSize + col
        board.cells[cellRow][cellCol] = player
        return newBoard, nil
    }
}

func (board *Board) FindNextBoards() {
    var directions [4][2]int = [4][2]int {{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    board.gameBoard[board.activeRow][board.activeCol].status = FinishedBoard

    for i, row := range board.gameBoard {
        for j, miniboard := range row {
            if miniboard.status == FinishedBoard {
                for _, move := range directions {
                    newRow := i + move[0]
                    newCol := j + move[1]
                    if newRow >= 0 && newRow < BoardDim && newCol >= 0 && newCol < BoardDim {
                        if board.gameBoard[newRow][newCol].status == UnplayedBoard {
                            board.gameBoard[newRow][newCol].status = PotentialBoard
                        }
                    }
                }
            }
        }
    }
}

func (board *Board) SetNextBoard(row, col int) (error) {
    if row < 0 || row >= BoardDim || col < 0 || col >= BoardDim {
        return errors.New("Coordinates out of range.")
    } else if board.gameBoard[row][col].status != PotentialBoard {
        return errors.New("Cannot play in this board.")
    } else {
        board.activeRow = row
        board.activeCol = col
        board.gameBoard[row][col].status = ActiveBoard
        for i, row := range board.gameBoard {
            for j:= range row {
                if board.gameBoard[i][j].status == PotentialBoard {
                    board.gameBoard[i][j].status = UnplayedBoard
                }
            }
        }
        return nil;
    }
}

func (board Board) Winner() int {

    var directions [4][2]int = [4][2]int {{0, 1}, {1, 0}, {1, 1}, {1, -1}}
    tieGame := true

    for i := 0; i < CellDim; i++ {
        for j := 0; j < CellDim; j++ {
            startVal := board.cells[i][j]
            if startVal != PlayerOneTile && startVal != PlayerTwoTile {
                if startVal == UnclaimedTile {
                    tieGame = false
                }
                continue
            }
            for k := range directions {
                win := true
                for l := 1; l < WinLen; l++ {
                    nRow := i + l * directions[k][0]
                    nCol := j + l * directions[k][1]
                    if nRow >= 0 && nRow < CellDim && nCol >= 0 && nCol < CellDim {
                        win = win && (board.cells[nRow][nCol] == startVal)
                    } else {
                        win = false
                    }
                }
                if win {
                    return startVal
                }
            }
        }
    }

    if tieGame {
        return TiedGame
    } else {
        return UnclaimedTile
    }

}
