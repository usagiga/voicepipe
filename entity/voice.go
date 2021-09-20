package entity

import "github.com/bwmarrin/discordgo"

type Voice struct {
	Room   *Room
	Packet *discordgo.Packet
}

// NewVoice is constructor of Voice
func NewVoice(room *Room, packet *discordgo.Packet) *Voice {
	return &Voice{
		Room:   room,
		Packet: packet,
	}
}
