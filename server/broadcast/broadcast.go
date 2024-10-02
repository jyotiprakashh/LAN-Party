package broadcast

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var clients = make(map[net.Conn]string)
var mutex = &sync.Mutex{}

func BroadcastUserList() {
	for {
		time.Sleep(1 * time.Second)

		userList := getUsersList()

		mutex.Lock()
		for conn := range clients {
			conn.Write([]byte("[USERLIST]" + userList + "\n"))
		}
		mutex.Unlock()
	}
}

func getUsersList() string {
	mutex.Lock()
	defer mutex.Unlock()

	var userList strings.Builder
	for _, username := range clients {
		userList.WriteString(username)
		userList.WriteString(",")
	}
	return userList.String()
}

func Broadcast(message string, sender net.Conn) {
	fmt.Print(message)
	mutex.Lock()
	defer mutex.Unlock()
	for conn := range clients {
		if conn == sender {
			conn.Write([]byte("(me)" + message))
		} else {
			conn.Write([]byte(message))
		}
	}
}
