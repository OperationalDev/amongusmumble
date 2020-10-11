package main

import (
	"amongusmumble/socket"
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"github.com/spf13/viper"
	"layeh.com/gumble/gumble"
	"layeh.com/gumble/gumbleutil"
)

func main() {
	var tlsConfig tls.Config

	tlsConfig.InsecureSkipVerify = true

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if err, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Cannot Find Config file")
			os.Exit(3)
		} else {
			fmt.Println("error loading config", err)
		}
	}

	server := viper.GetString("mumbleserver")
	listenaddress := viper.GetString("listenaddress")
	listenport := viper.GetString("listenport")

	config := gumble.NewConfig()
	config.Username = viper.GetString("username")
	certificateFile := viper.GetString("cert")
	keyFile := viper.GetString("key")

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
	client, err := gumble.DialWithDialer(new(net.Dialer), server, config, &tlsConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		if reject, ok := err.(*gumble.RejectError); ok {
			os.Exit(100 + int(reject.Type))
		}
		os.Exit(99)
	}

	go socket.SocketioServer(client, listenaddress, listenport)

	<-keepAlive
	os.Exit(exitStatus)
}
