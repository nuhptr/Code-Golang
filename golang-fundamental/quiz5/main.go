package main

import "fmt"

func main() {
	games := Games{
		Name: "Zelda",
	}

	games.addGame("Material Burst")
	games.addGame("Heavy Conduct")

	for _, game := range games.Game{
		fmt.Println(game)
	}
}

func (games *Games) addGame(game string)  {
	games.Game = append(games.Game, game)
}

type Games struct {
	Name string
	Game []string
}