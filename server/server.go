package main

import (
	"flag"
	"fmt"
	"github.com/ddunkin/launcher"
	"html"
	"log"
	"net/http"
	"time"
)

type command struct {
	command        byte
	durationMillis int64
}

func main() {
	l := launcher.Create()
	defer l.Destroy()

	commandChannel := make(chan command)
	go handleCommands(l, commandChannel)

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/launcher/", func(w http.ResponseWriter, r *http.Request) {
		const prefixLen = len("/launcher/")
		commandName := r.URL.Path[prefixLen:]
		fmt.Fprintf(w, "Missile Launcher: %s", html.EscapeString(commandName))
		log.Println(commandName)
		const turnDuration = 150
		var cmd command
		switch commandName {
		case "left":
			cmd = command{launcher.Left, turnDuration}
		case "right":
			cmd = command{launcher.Right, turnDuration}
		case "up":
			cmd = command{launcher.Up, turnDuration}
		case "down":
			cmd = command{launcher.Down, turnDuration}
		case "fire":
			cmd = command{launcher.Fire, 0}
		}
		commandChannel <- cmd
	})

	port := flag.Uint("p", 8888, "Port to listen on")
	flag.Parse()

	log.Printf("Listening on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func handleCommands(l *launcher.Launcher, commandChannel chan command) {
	for {
		cmd := <-commandChannel
		l.SendCommand(cmd.command)
		if (cmd.durationMillis != 0) {
			time.Sleep(time.Duration(cmd.durationMillis * int64(time.Millisecond)))
			l.SendCommand(launcher.Stop)
		}
	}
}
