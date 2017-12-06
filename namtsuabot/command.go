package namtsuabot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Command represents a bot command
type Command interface {
	Name() string
	Process([]string, *discordgo.Message, []int, *GuildInfo) (string, bool)
	Usage(*GuildInfo) *CommandUsage
	UsageShort() string
}

// CommandUsage defines the help features of a command
type CommandUsage struct {
	Description string
	Params      []CommandUsageParam
}

// CommandUsageParam defines the help features of a command param
type CommandUsageParam struct {
	Name        string
	Description string
	Optional    bool
	Variadic    bool
}

// HandleCommand processes a NamtsuaBot command
func HandleCommand(s *discordgo.Session, m *discordgo.Message, info *GuildInfo) {
	var prefix byte = '$'

	if len(m.Content) > 1 && m.Content[0] == prefix {

		args, indices := parseArguments(m.Content[1:])
		arg := strings.ToLower(args[0])
		c, ok := info.commands[arg]
		if ok {
			result, usePM := c.Process(args[1:], m, indices[1:], info)
			if len(result) > 0 {
				targetChannel := m.ChannelID
				if usePM {
					channel, err := s.UserChannelCreate(m.Author.ID)
					if err == nil {
						targetChannel = channel.ID
						info.SendMessage(m.ChannelID, "Check your PMs for my reply!")
					}
				}

				info.SendMessage(targetChannel, result)
			}
		}
	}
}
