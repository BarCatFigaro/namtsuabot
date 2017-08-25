package namtsuabot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Bot represents a Namtsuabot instance
type Bot struct {
	// Session is a discordgo session
	Session *discordgo.Session
}

// New creates a new bot instance
func New(token string) *Bot {
	nb, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("problem creating namtsuabot; %v\n", err)
	}

	b := &Bot{
		Session: nb,
	}
	return b
}

// Connect opens a websocket connection
// func will return after disconnecting
func (b *Bot) Connect(cli bool) {
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

	out := ""

	if m.Content == "!chess" {
		out = chessCommand()
	}

	s.ChannelMessageSend(m.ChannelID, out)
}

func chessCommand() string {
	/*
			board := `
						8 ║♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
						7 ║♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟
						6 ║
						5 ║
						4 ║
						3 ║
						2 ║♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
						1 ║♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
		                 —╚═══════════════
						—— a  b c  d e  f g  h`
	*/
	return buildBoard()
}

func buildBoard() string {
	board := "Chess board" + `    

		♜ ♞ ♝ ♛ ♚ ♝ ♜ ♞
		♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟
		
		
		
		
		♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
		♖ ♘ ♗ ♕ ♔ ♖ ♘ ♗`

	return board
}

// TODO
func santizeBoard(board string) string {
	return ""
}
