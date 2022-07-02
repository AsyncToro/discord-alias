package main

import (
	"fmt"

	"github.com/AsyncToro/discord-alias/bot"
	"github.com/AsyncToro/discord-alias/config"

)

var (
	Token string
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = bot.Start()
	if err != nil {
		fmt.Println("error starting bot")
		fmt.Println(err.Error())
		return
	}

	<-make(chan struct{})

	return
}
