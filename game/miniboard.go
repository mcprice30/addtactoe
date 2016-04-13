package game

import (
    "errors"
)

const PlayerOneTile = 1
const PlayerTwoTile = 2
const UnplayableTile = -1
const UnclaimedTile = 0

const MiniBoardSize = 3

type MiniBoard struct {
	tiles        [MiniBoardSize][MiniBoardSize]int
	claimedTiles int
	status       int
}

func NewMiniBoard() MiniBoard {
	miniboard := MiniBoard{}
	miniboard.claimedTiles = 0
	miniboard.status = UnplayedBoard
	for i, row := range miniboard.tiles {
		for j := range row {
			miniboard.tiles[i][j] = UnclaimedTile
			if i == 1 && j == 1 {
				miniboard.tiles[i][j] = UnplayableTile
			}
		}
	}
	return miniboard
}

func (miniboard MiniBoard) GetTiles() [MiniBoardSize][MiniBoardSize]int {
	return miniboard.tiles
}

func (miniboard *MiniBoard) MakeMove(row, col, player int) (bool, error) {
    if row >= MiniBoardSize || row < 0 || col >= MiniBoardSize || col < 0 {
		return false, errors.New("Move out of bounds.")
	} else if player != PlayerOneTile && player != PlayerTwoTile {
		return false, errors.New("Invalid player.")
	} else if miniboard.tiles[row][col] != 0 {
		return false, errors.New("Tile already claimed!")
	} else {
		miniboard.claimedTiles++
		miniboard.tiles[row][col] = player
		return miniboard.claimedTiles == 8, nil
	}
}

func statusDescriptor(status int) string {
	switch status {
	case ActiveBoard:
		return "@"
	case UnplayedBoard:
		return "."
	case FinishedBoard:
		return "*"
	case PotentialBoard:
		return "?"
	default:
		return "!"
	}
}

func (miniboard MiniBoard) RowString(row int) (string, error) {
	if row >= MiniBoardSize || row < 0 {
		return "", errors.New("Row out of bounds.")
	}
	out := ""
	for col := range miniboard.tiles[row] {
		switch miniboard.tiles[row][col] {
		case PlayerOneTile:
			out += "X"
		case PlayerTwoTile:
			out += "O"
		case UnclaimedTile:
			if miniboard.status == ActiveBoard {
				out += "?"
			} else {
				out += " "
			}
		case UnplayableTile:
			out += statusDescriptor(miniboard.status)
		default:
			out += "!"
		}
	}
	return out, nil
}

func (miniboard MiniBoard) ToString() string {
	out := ""
	for row := range miniboard.tiles {
		rowString, _ := miniboard.RowString(row)
		out += rowString + "\n"
	}
	return out
}
