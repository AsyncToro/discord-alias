package main

import (
	"fmt"
	"strings"

	"github.com/AsyncToro/discord-alias/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotId   string
	discord *discordgo.Session
)

func startBot() error {
	discord, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		return fmt.Errorf("Can't initiate a new bot %w", err)
	}

	botUser, err := discord.User("@me")
	if err != nil {
		return fmt.Errorf("Bot can't be assigned as user %w", err)
	}

	BotId = botUser.ID

	discord.AddHandler(messageHandler)

	err = discord.Open()
	if err != nil {
		return fmt.Errorf("Error opening a session to discord %w", err)
	}

	return nil
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == BotId {
		return
	}

	if !strings.HasPrefix(message.Content, config.BotPrefix) {
		return
	}

	_, _ = session.ChannelMessageSend(message.ChannelID, "pong")
}
