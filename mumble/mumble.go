package mumble

import (
	"log"
	"strings"

	"layeh.com/gumble/gumble"
)

// Kill marks a player as dead
func Kill(c *gumble.Client, player string, gamestate string, deadplayers []string) []string {
	alive := c.Channels.Find("AmongUs", "Alive")

	log.Println("In game player:", player)
	aliveusers := c.Channels[alive.ID].Users
	var player2 string
	for _, element := range aliveusers {
		if element.Comment == player {
			player2 = strings.TrimSpace(element.Name)
			log.Println("Mumble user:", player2)
		}
	}

	duplicateplayer := 0

	for _, s := range deadplayers {
		if s == player2 {
			duplicateplayer = duplicateplayer + 1
		}
	}

	if player2 != "" {
		if duplicateplayer == 0 {
			log.Println(player2, "is now dead")
			deadplayers = append(deadplayers, strings.TrimSpace(player2))
		} else {
			log.Println(player2, "is already dead")
		}
	} else {
		log.Println("Ignoring blank player name")
	}

	log.Println("Dead Players:", deadplayers)

	return deadplayers
}

// Startgame starts the game
func Startgame(c *gumble.Client) {
	lobby := c.Channels.Find("AmongUs", "Lobby")
	alive := c.Channels.Find("AmongUs", "Alive")
	lobbyusers := c.Channels[lobby.ID].Users
	for _, element := range lobbyusers {
		element.Move(alive)
		element.SetMuted(true)
		element.SetDeafened(true)
		log.Println("Moving", element.Name, "to #alive")
	}
}

// Meeting starts discussion phase
func Meeting(c *gumble.Client, deadplayers []string) {
	alive := c.Channels.Find("AmongUs", "Alive")
	aliveusers := c.Channels[alive.ID].Users

	for _, element := range aliveusers {
		log.Println("Unmute player", element.Name)
		element.SetMuted(false)
		element.SetDeafened(false)
		log.Println(element.Name, "is alive")
	}

	for _, deadplayer := range deadplayers {
		user := c.Users.Find(deadplayer)
		if user != nil {
			log.Println("Mute player", user.Name)
			user.SetMuted(true)
			user.SetDeafened(false)
			user.Move(alive)
			log.Println(user.Name, "is dead")
		}
	}
}

// Resumegame Resumes Game (tasks)
func Resumegame(c *gumble.Client, deadplayers []string) {
	alive := c.Channels.Find("AmongUs", "Alive")
	dead := c.Channels.Find("AmongUs", "Dead")

	aliveusers := c.Channels[alive.ID].Users

	log.Println("Resuming game")

	for _, element := range aliveusers {
		log.Println("Mute player", element.Name)
		element.SetMuted(true)
		element.SetDeafened(true)
		log.Println(element.Name, "is alive")
	}

	for _, deadplayer := range deadplayers {
		user := c.Users.Find(deadplayer)
		if user != nil {
			log.Println("Unmute player", user.Name)
			user.SetMuted(false)
			user.SetDeafened(false)
			user.Move(dead)
			log.Println(user.Name, "is dead")
		}
	}
}

// Endgame ends game
func Endgame(c *gumble.Client) {
	lobby := c.Channels.Find("AmongUs", "Lobby")
	alive := c.Channels.Find("AmongUs", "Alive")
	dead := c.Channels.Find("AmongUs", "Dead")

	aliveusers := c.Channels[alive.ID].Users
	deadusers := c.Channels[dead.ID].Users

	for _, element := range aliveusers {
		element.Move(lobby)
		element.SetMuted(false)
		element.SetDeafened(false)
		log.Println("Unmute player", element.Name)
	}

	for _, element := range deadusers {
		element.Move(lobby)
		element.SetMuted(false)
		element.SetDeafened(false)
		log.Println("Unmute player", element.Name)
	}
}

// Namecheck makes sure player has a valid comment
func Namecheck(c *gumble.Client, player string) {
	var player2 string

	lobby := c.Channels.Find("AmongUs", "Lobby")
	lobbyusers := c.Channels[lobby.ID].Users

	log.Println("Checking if", player, "has a mumble user set")
	for _, element := range lobbyusers {
		if element.Comment == player {
			player2 = element.Name
			log.Println("User set:", player2, "==", player)
			return
		}
	}
	lobby.Send("Player "+player+" does not have a mumble user set.", true)
	log.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	log.Println("Player", player, "does not have a mumble user set.")
	log.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}
