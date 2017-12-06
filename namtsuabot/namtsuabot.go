package namtsuabot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var nb *NamtsuaBot
var channelRegex = regexp.MustCompile("<#[0-9]+>")
var roleRegex = regexp.MustCompile("<@&[0-9]+>")
var userRegex = regexp.MustCompile("<@!?[0-9]+>")
var mentionRegex = regexp.MustCompile("<@(!|&)?[0-9]+>")

// NamtsuaBot represents a Namtsuabot instance
type NamtsuaBot struct {
	Session     *discordgo.Session
	MainGuildID string
	SelfID      uint64
	Info        *GuildInfo
}

// New creates a new bot instance
func New(token, guildID string) *NamtsuaBot {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("problem creating namtsuabot; %v\n", err)
	}

	temp, _ := s.Guild(guildID)
	g := &GuildInfo{
		ID:       guildID,
		Name:     temp.Name,
		OwnerID:  sanitizeUser(temp.OwnerID),
		commands: make(map[string]Command),
	}

	b := &NamtsuaBot{
		Session:     s,
		MainGuildID: guildID,
		Info:        g,
	}
	g.Bot = b

	return b
}

// Connect opens a websocket connection
// func will return after disconnecting
func (b *NamtsuaBot) Connect(cli bool) {
	if err := b.Session.Open(); err != nil {
		fmt.Printf("error opening connection: %v\n", err)
		return
	}
	fmt.Println("Connection established.")

	b.Session.AddHandler(messageCreate)

	if cli {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc
	}

	fmt.Println("Connection closed.")
	b.Session.Close()
}

// messageCreate fires when a message is created
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
}
