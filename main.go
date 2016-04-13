package main

import (
	"github.com/mcprice30/addtactoe/game"
)

func main() {
    p1 := game.NewConsoleHumanPlayer(1, "A")
    p2 := game.NewConsoleHumanPlayer(2, "B")
    g := game.NewGame(p1, p2)
    g.Play()
}
