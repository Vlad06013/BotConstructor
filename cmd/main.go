package main

import (
	config "github.com/Vlad06013/BotConstructor.git"
	"github.com/Vlad06013/BotConstructor.git/domain/useCase"
)

func main() {
	err := config.SetEnvValues()
	if err != nil {
		panic(err)
	}

	useCase.Listen()
}
