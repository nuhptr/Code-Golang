package main

import "fmt"

func main() {
	// tidak langsung di isi
	//var gamingConsole []string
	//gamingConsole = append(gamingConsole, "Playstation4")
	//gamingConsole = append(gamingConsole, "Nintendo")
	//gamingConsole = append(gamingConsole, "Switch")
	//fmt.Println(gamingConsole)

	// langsung diisi
	gamingConsole := []string {
		"PS4", "Nintendo", "Xbox",
	}

	fmt.Println(gamingConsole)

	for _, console := range gamingConsole {
		fmt.Println(console)
	}
}
