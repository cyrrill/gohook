package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	auth "github.com/cyrrill/go-http-auth"
)

// Holds messages
var messages = make(chan string, 10)

func main() {

	// Listen for webhook requests on HTTP with basic auth
	htpasswd := auth.HtpasswdFileProvider("./.htpasswd")
	authenticator := auth.NewBasicAuthenticator("Basic Realm", htpasswd)
	http.HandleFunc("/webhook", authenticator.Wrap(hookHandler))
	go http.ListenAndServe(":8080", nil)

	// Listen for clients on TCP
	clientServer, err := net.Listen("tcp", ":7700")
	if err == nil {
		conn, err := clientServer.Accept()
		if err == nil {
			clientHandler(conn)
		}
	}
}

// Handles client connections
func clientHandler(conn net.Conn) {
	defer conn.Close()
	for {
		select {

		// If we have a message in the channel pass it along
		case d := <-messages:
			fmt.Fprintln(conn, d)

		// Ping every 30s in case of no messages
		case <-time.After(30 * time.Second):
			fmt.Fprintln(conn, "ping")
		}
	}
}

// When a webhook post comes in, pass into messages channel
func hookHandler(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	data, err := ioutil.ReadAll(r.Body)
	if err == nil {
		messages <- string(data)
	}
}
