package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var clients = make(map[net.Conn]string)
var mutex = &sync.Mutex{}


func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer ln.Close()

    fmt.Println("Server started on port 8080")
    go broadcastUserList()

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        go handleConnection(conn)
    }
}

func broadcastUserList() {
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

func handleConnection(conn net.Conn) {
    defer conn.Close()
    // conn.Write([]byte("Enter your username: "))

    username, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        return
    }
    username = strings.TrimSpace(username)

    mutex.Lock()
    clients[conn] = username
    clientCount := len(clients)
    mutex.Unlock()

    // broadcast(fmt.Sprintf("%s joined the chat\n", username), conn)
    if clientCount == 1 {
        conn.Write([]byte("You are alone in the chat. Invite your friends to join!\n"))
    } else {
        broadcast(fmt.Sprintf("%s joined the chat\n", username), nil)
    }

    for {
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            mutex.Lock()
            delete(clients, conn)
            mutex.Unlock()
            broadcast(fmt.Sprintf("%s left the chat\n", username), conn)
            return
        }

        broadcast(fmt.Sprintf("%s: %s", username, message), conn)
    }
}

func broadcast(message string, sender net.Conn) {
    fmt.Print(message)
    mutex.Lock()
    defer mutex.Unlock()
    for conn := range clients {
        if conn == sender {
            conn.Write([]byte("(me)"+message ))
        } else{
            conn.Write([]byte(message))
        }
    }
}

// package main

// import (
// 	"fmt"
// 	"net"
// 	"os"

// 	"lan-party/server/broadcast"
// 	"lan-party/server/connection"
// )

// func main() {
// 	ln, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	defer ln.Close()

// 	fmt.Println("Server started on port 8080")
// 	go broadcast.BroadcastUserList()

// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}

// 		go connection.HandleConnection(conn)
// 	}
// }

