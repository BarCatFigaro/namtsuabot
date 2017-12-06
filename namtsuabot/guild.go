package namtsuabot

import (
	"fmt"
	"strings"
)

// GuildInfo holds the data of a guild
type GuildInfo struct {
	ID      string
	Name    string
	OwnerID uint64
	Bot     *NamtsuaBot
	//config   BotConfig
	commands map[string]Command
}

// SendMessage sends a message to the given channel
func (info *GuildInfo) SendMessage(channelID, message string) {
	for len(message) > 1999 {

		if message[0:3] == "```" && message[len(message)-3:] == "```" {
			index := strings.LastIndex(message[:1995], "\n")

			if index < 10 {
				index = 1995
			}

			info.sendContent(channelID, message[:index]+"```")
			message = "```\n" + message[index:]
		} else {
			index := strings.LastIndex(message[:1999], "\n")

			if index < 10 {
				index = 1999
			}

			info.sendContent(channelID, message[:index])
			message = message[index:]
		}
	}
	info.sendContent(channelID, message)
}

func (info *GuildInfo) sendContent(channelID, message string) {
	_, err := info.Bot.Session.ChannelMessageSend(channelID, message)

	if err != nil {
		fmt.Printf("failed to send message: %s\n", err.Error())
	}
}
