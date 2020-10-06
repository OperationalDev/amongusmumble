package main

import (
	"amongusmumble/mumble"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"layeh.com/gumble/gumble"
	"layeh.com/gumble/gumbleutil"
)

const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02 15:04 MST"
)

var deadplayers []string
var gamestate string
var gameup bool
var gamestatetime time.Time

type Player struct {
	//Action       PlayerAction `json:"Action"`
	Name         string `json:"Name"`
	Color        int    `json:"Color"`
	IsDead       bool   `json:"IsDead"`
	Disconnected bool   `json:"Disconnected"`
}

func socketioServer(client *gumble.Client) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "connect", func(s socketio.Conn, msg string) {
		log.Println("set connect code:", msg)
		s.Emit("reply", "set guildID successfully")
	})
	server.OnEvent("/", "state", func(s socketio.Conn, msg string) {
		log.Println("Phase received from capture: ", msg)

		switch msg {
		case "0":
			gamestate = "LOBBY"
		case "1":
			gamestate = "TASKS"
		case "2":
			gamestate = "DISCUSSION"
		}
		log.Println("Gamestate set:", gamestate)
		switch gamestate {
		case "MENU":
			log.Println("Gamemode: Menu")
			//mumble.Endgame(client)
			deadplayers = nil
			gameup = false
		case "LOBBY":
			log.Println("Gamemode: LOBBY")
			//mumble.Endgame(client)
			deadplayers = nil
			gameup = false
		case "DISCUSSION":
			log.Println("Gamemode: DISCUSSION")
			//mumble.Meeting(client, deadplayers)
		case "TASKS":
			log.Println("Gamemode: TASKS")
			gamestatetime = time.Now()
			log.Println("Game State Time:", gamestatetime)
			time.Sleep(5 * time.Second)
			if gameup == false {
				mumble.Startgame(client)
			} else {
				mumble.Resumegame(client, deadplayers)
			}
			gameup = true
		}
	})
	server.OnEvent("/", "player", func(s socketio.Conn, msg string) {
		log.Println("Player received from capture: ", msg)
		log.Println("Gamestate: ", gamestate)
		player := Player{}
		_ = json.Unmarshal([]byte(msg), &player)

		if gamestate == "LOBBY" {
			mumble.Namecheck(client, strings.TrimSpace(player.Name))
		} else {
			if player.IsDead == true {
				deadplayers = mumble.Kill(client, strings.TrimSpace(player.Name), gamestate, deadplayers)
				duration := time.Since(gamestatetime)
				if duration.Seconds() < 10 {
					log.Println("Move", player.Name, "to Dead now")
					dead := client.Channels.Find("AmongUs", "Dead")
					user := client.Users.Find(player.Name)
					user.Move(dead)
					user.SetMuted(false)
					user.SetDeafened(false)
				} else {
					log.Println("Move", player.Name, "to Dead at end of round")
				}
			}
		}
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:8123...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8123", nil))
}

func main() {
	var tlsConfig tls.Config

	tlsConfig.InsecureSkipVerify = true
	server := flag.String("server", "sg.whysomadtho.com:64738", "address of the server")
	config := gumble.NewConfig()
	config.Username = "6ix9ine"
	certificateFile := "6ix9ine.crt"
	keyFile := "6ix9ine.key"
	if certificateFile != "" {
		if keyFile == "" {
			keyFile = certificateFile
		}
		if certificate, err := tls.LoadX509KeyPair(certificateFile, keyFile); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
			os.Exit(1)
		} else {
			tlsConfig.Certificates = append(tlsConfig.Certificates, certificate)
		}
	}
	keepAlive := make(chan bool)
	exitStatus := 0
	config.Attach(gumbleutil.Listener{
		Disconnect: func(e *gumble.DisconnectEvent) {
			if e.Type != gumble.DisconnectUser {
				exitStatus = int(e.Type) + 1
			}
			keepAlive <- true
		},
	})
	client, err := gumble.DialWithDialer(new(net.Dialer), *server, config, &tlsConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		if reject, ok := err.(*gumble.RejectError); ok {
			os.Exit(100 + int(reject.Type))
		}
		os.Exit(99)
	}

	go socketioServer(client)

	<-keepAlive
	os.Exit(exitStatus)
}
