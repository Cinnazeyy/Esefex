package bot

import (
	"esefexapi/bot/commands"
	"esefexapi/service"
	"esefexapi/sounddb"

	"log"

	"github.com/bwmarrin/discordgo"
)

var _ service.Service = &DiscordBot{}

// DiscordBot implements Service
type DiscordBot struct {
	Session *discordgo.Session
	cmdh    *commands.CommandHandlers
	db      sounddb.ISoundDB
	stop    chan struct{}
	ready   chan struct{}
}

func NewDiscordBot(s *discordgo.Session, db sounddb.ISoundDB) *DiscordBot {
	return &DiscordBot{
		Session: s,
		cmdh:    commands.NewCommandHandlers(db),
		stop:    make(chan struct{}, 1),
		ready:   make(chan struct{}),
	}
}

func (b *DiscordBot) run() {
	defer close(b.stop)
	log.Println("Starting bot...")

	s := b.Session

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	log.Println("Adding commands...")
	b.RegisterComands(s)
	defer b.DeleteAllCommands(s)
	// defer actions.LeaveAllChannels(s)

	log.Println("Bot Ready.")
	close(b.ready)
	<-b.stop
}

func (b *DiscordBot) Start() <-chan struct{} {
	go b.run()
	return b.ready
}

func (b *DiscordBot) Stop() <-chan struct{} {
	b.stop <- struct{}{}
	return b.stop
}
