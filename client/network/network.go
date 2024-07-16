package network

import (
	"bufio"
	// "fmt"
	"net"
	"strings"
)

var conn net.Conn

func Connect() error {
	var err error
	conn, err = net.Dial("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	return nil
}

func ReadMessages(updateUsers func(int), updateMessages func(string), handleError func(error)) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			handleError(err)
			return
		}
		if strings.HasPrefix(message, "[USERLIST]") {
			usersStr := strings.TrimPrefix(message, "[USERLIST]")
			usernames := strings.Split(strings.TrimSpace(usersStr), ",")
			updateUsers(len(usernames) - 1)
		} else {
			updateMessages(message)
		}
	}
}

func SendMessage(message string) {
	conn.Write([]byte(message))
}
