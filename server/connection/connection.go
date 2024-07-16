package connection

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"

	"lan-party/server/broadcast"
)

var clients = make(map[net.Conn]string)
var mutex = &sync.Mutex{}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	username, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}
	username = strings.TrimSpace(username)

	mutex.Lock()
	clients[conn] = username
	clientCount := len(clients)
	mutex.Unlock()

	if clientCount == 1 {
		conn.Write([]byte("You are alone in the chat. Invite your friends to join!\n"))
	} else {
		broadcast.Broadcast(fmt.Sprintf("%s joined the chat\n", username), nil)
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			broadcast.Broadcast(fmt.Sprintf("%s left the chat\n", username), conn)
			return
		}

		broadcast.Broadcast(fmt.Sprintf("%s: %s", username, message), conn)
	}
}
