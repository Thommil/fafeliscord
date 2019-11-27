package pingpong

import (
	fmt "fmt"

	discordgo "github.com/bwmarrin/discordgo"
)

type PingPong struct {
}

func New() (*PingPong, error) {
	fmt.Println("New PingPong handler")
	return &PingPong{}, nil
}

func (instance PingPong) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
