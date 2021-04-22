package main

import (
	"fmt"

	"github.com/RioRizkyRainey/pokedex/internal/moves/service"
)

func main() {
	fmt.Println()

	fmt.Println("Moves service start")

	service.Start()
}
