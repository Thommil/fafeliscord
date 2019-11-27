package main

import (
	fmt "fmt"
	os "os"
	signal "os/signal"
	syscall "syscall"

	discordgo "github.com/bwmarrin/discordgo"
	collector "github.com/thommil/fafeliscord/handlers/collector"
	pingpong "github.com/thommil/fafeliscord/handlers/pingpong"
)

func main() {
	// Create session
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	fmt.Println("Fafeliscord Bot ot created")

	// Handlers
	if handler, err := pingpong.New(); err != nil {
		fmt.Println("Failed to create PingPong handler", err)
	} else {
		discord.AddHandler(handler.Handler)
		fmt.Println("Handler PingPong added")
	}

	if handler, err := collector.New(); err != nil {
		fmt.Println("Failed to Collector handler", err)
	} else {
		discord.AddHandler(handler.Handler)
		fmt.Println("Handler Collector added")
	}

	// Open Socket
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Fafeliscord session started")

	// Wait
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// End session
	discord.Close()
}
