package command

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	err = iota
	nick
)

const ident = "$"

// Command holds the context of a command
type Command interface {
	handleMessage(message []string) error
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
}

// Error holds the data for an error command
type Error struct {
	log string
}

// Execute returns the error log
func (ec *Error) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, ec.log)
}

func (ec *Error) handleMessage(message []string) error {
	if 1 != len(message) {
		return fmt.Errorf("not a valid message")
	}
	ec.log = fmt.Sprintf("\"%s\" is not a command, please try again.", message[0])
	return nil
}

// Nickname holds the data for a nickname command
type Nickname struct {
	from string
	to   string
}

// Execute changes the nickname of a member
func (nc *Nickname) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	channel, _ := s.State.Channel(m.Message.ChannelID)
	users := m.Message.Mentions
	if len(users) <= 0 {
		s.ChannelMessageSend(m.ChannelID, "no one tagged to change")
	}
	s.GuildMemberNickname(channel.GuildID, users[0].ID, nc.to)
}

func (nc *Nickname) handleMessage(message []string) error {
	if 3 != len(message) {
		return fmt.Errorf("not valid nickname inputs")
	}
	nc.from = message[1]
	nc.to = message[2]
	return nil
}

// Is returns true if the given message is a command
func Is(message string) bool {
	return strings.HasPrefix(message, ident)
}

// New returns a new command
func New(typ int, message []string) (Command, error) {
	var cmd Command
	switch typ {
	case 0:
		cmd = &Error{}

	case 1:
		cmd = &Nickname{}
	default:
		fmt.Println("Should not have been here!")
	}

	err := cmd.handleMessage(message)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// Find returns the command type
func Find(typ string) int {
	var out int

	switch typ {
	case "nick":
		out = nick
	default:
		out = err
	}

	return out
}
