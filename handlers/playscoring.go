package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func PlayScoring(s *discordgo.Session, event *discordgo.PresenceUpdate) {

	fmt.Println(event.Presence.Game.Name)
}
