package main

import (
	"fmt"

	"github.com/RioRizkyRainey/pokedex/internal/attack/service"
)

func main() {
	fmt.Println()

	fmt.Println("Attack service start")

	service.Start()
}
