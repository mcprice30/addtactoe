package game

type Player interface {
    ClaimCellMove(board Board) (int, int)
    NewMiniBoardMove(board Board) (int, int)
    PlayerNumber() int
}
