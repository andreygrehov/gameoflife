package main

import "github.com/andreygrehov/gameoflife/life"

func main() {
	game := life.New()
	game.Start()
}
