package bot

import (
	"discord-bot/config"
	"discord-bot/games"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/patrickmn/go-cache"
	"log"
	"strings"
	"time"
)

var BotID string

var GamesCache *cache.Cache

func Start(cfg *config.Config) (*discordgo.Session, error) {

	GamesCache = cache.New(5*time.Minute, 10*time.Minute)

	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return nil, err
	}

	err = session.Open()
	if err != nil {
		return nil, err
	}

	fmt.Println("Bot is running!")

	user, err := session.User("@me")
	if err != nil {
		return nil, err
	}

	BotID = user.ID

	session.AddHandler(gamesHandler)

	return session, nil
}

func gamesHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content != "!games" {
		return
	}

	cachedGames, isCache := GamesCache.Get("games")
	if isCache {
		log.Println("Cache hit")
		_, _ = s.ChannelMessageSend(m.ChannelID, cachedGames.(string))
		return
	}

	gameList, err := games.GetGames()
	if err != nil {
		fmt.Println(err)
		_, _ = s.ChannelMessageSend(m.ChannelID, "error in getting posts")
		return
	}

	var sb strings.Builder

	for _, game := range gameList {
		sb.WriteString(fmt.Sprintf(
			"**%s** \n💰 De: %s\n💸 Por: %s\n🔗 [Link](%s)\n\n",
			game.Name,
			game.OriginalPrice,
			game.FinalPrice,
			game.URL,
		))
	}

	GamesCache.Set("games", sb.String(), time.Minute*5)

	_, err = s.ChannelMessageSend(m.ChannelID, sb.String())
	if err != nil {
		fmt.Println(err)
	}
}
