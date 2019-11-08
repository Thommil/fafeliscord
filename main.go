package main

import (
	fmt "fmt"
	os "os"
	signal "os/signal"
	syscall "syscall"

	discordgo "github.com/bwmarrin/discordgo"
	handlers "github.com/thommil/fafeliscord/handlers"
)

func main() {
	// Create session
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	fmt.Println("Fafeliscord Bot ot created")

	// Add handlers
	fmt.Println("Handler PingPong added")
	discord.AddHandler(handlers.PingPong)

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
