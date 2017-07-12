package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/translate"
	"google.golang.org/api/googleapi/transport"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/oauth2"
	"golang.org/x/text/language"
)

// Configuration for APIs
type Configuration struct {
	GoogleKey    string
	DiscordToken string
}

var configuration Configuration

func main() {

	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Add config.json to this directory; %v\n", err)
	}
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error: ", err)
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + configuration.DiscordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Author.ID == "194259676488138763" {
		if m.Message.Content == "drop messages" {
			msgs, _ := s.ChannelMessages(m.ChannelID, 100, "", "", "")
			for _, msg := range msgs {
				if msg.Author.ID == s.State.User.ID {
					s.ChannelMessageDelete(m.ChannelID, msg.ID)
				}
			}
		}
	}

	msg, _ := s.ChannelMessage(m.ChannelID, m.Message.ID)
	for _, translation := range translator(msg.Content) {
		s.ChannelMessageSend("334562076972548103", translation.Text)
	}
}

func translator(text string) []translate.Translation {
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
		Transport: &transport.APIKey{Key: configuration.GoogleKey},
	})
	// create client
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Printf("failed to create client: %v", err)
	}

	target, err := language.Parse("ru")
	if err != nil {
		log.Printf("failed to parse language %v", err)
	}

	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Printf("failed to translate text %v", err)
	}

	return translations
}

/*
func sentimenter(text string) string {
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
		Transport: &transport.APIKey{Key: configuration.GoogleKey},
	})

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Detects the sentiment of the text.
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	if sentiment.DocumentSentiment.Score >= 0 {
		return fmt.Sprintf("Text: %v\n", text) + "\nSentiment: positive"
	}
	return fmt.Sprintf("Text: %v\n", text) + "\nSentiment: negative"

}
*/
