package main

import (
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Missile Launcher")
	})

	http.HandleFunc("/launcher/", func(w http.ResponseWriter, r *http.Request) {
		const prefixLen = len("/launcher/")
		commandName := r.URL.Path[prefixLen:]
		fmt.Fprintf(w, "Missile Launcher: %s", html.EscapeString(commandName))
		log.Println(commandName)
		const turnDuration = 250
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
			cmd = command{launcher.Fire, 3000}
		}
		commandChannel <- cmd
	})

	log.Println("Listening")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCommands(l *launcher.Launcher, commandChannel chan command) {
	for {
		cmd := <-commandChannel
		l.SendCommand(cmd.command)
		time.Sleep(time.Duration(cmd.durationMillis * int64(time.Millisecond)))
		l.SendCommand(launcher.Stop)
	}
}
