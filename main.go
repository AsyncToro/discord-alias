package main

import (
	"fmt"

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

	_, errrr := newStore()
	if errrr != nil {
		fmt.Println(errrr.Error())
		return
	}

	err = startBot()
	if err != nil {
		fmt.Println("error starting bot")
		fmt.Println(err.Error())
		return
	}

	<-make(chan struct{})

	return
}
