// // client.go
// package main

// import (
//     "bufio"
//     "fmt"
//     "net"
//     // "os"
//     "strings"

//     "github.com/gdamore/tcell/v2"
//     "github.com/rivo/tview"
// )

// var messageBox *tview.TextView
// var inputField *tview.InputField
// var app *tview.Application
// var conn net.Conn
// var userList *tview.TextView

// func main() {
//     var err error
//     conn, err = net.Dial("tcp", "localhost:8080")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     defer conn.Close()

//     app = tview.NewApplication()

// 	header := tview.NewTextView().
// 	SetText("Welcome to LAN Party").
// 	SetTextAlign(tview.AlignCenter).
// 	SetTextColor(tcell.ColorYellow) 
	
//     userList = tview.NewTextView().
//         SetTextAlign(tview.AlignRight).
//         SetText("Active Users:\n").
//         SetTextAlign(tview.AlignLeft)

//     messageBox = tview.NewTextView().
//         SetDynamicColors(true).
//         SetRegions(true).
//         SetWordWrap(true).
//         SetTextColor(tcell.ColorGreen)

		
// 		go readMessages()

// 		inputField = tview.NewInputField().
//         SetLabel("Enter username: ").
//         SetFieldWidth(0).
//         SetDoneFunc(func(key tcell.Key) {
//             if key == tcell.KeyEnter {
//                 text := inputField.GetText()
//                 if strings.TrimSpace(text) != "" {
//                     conn.Write([]byte(text + "\n"))
//                     inputField.SetLabel("Enter message: ")
//                     inputField.SetText("")
//                     inputField.SetDoneFunc(sendMessage)
//                 }
//             }
//         })

		

//     flex := tview.NewFlex().
//         SetDirection(tview.FlexRow).
// 		AddItem(header, 1, 1, false). 
//         AddItem(messageBox, 0, 1, false).
// 		AddItem(userList, 1, 1, false).
//         AddItem(inputField, 1, 0, true)


// 	// grid.AddItem(header, 0, 0, 1, 1, 0, 0, true)
// 	// grid.AddItem(flex, 1, 0, 1, 1, 0, 0, true)


//     if err := app.SetRoot(flex, true).Run(); err != nil {
//         panic(err)
//     }

// }

// func readMessages() {
//     reader := bufio.NewReader(conn)
//     for {
//         message, err := reader.ReadString('\n')
//         if err != nil {
//             app.Stop()
//             return
//         }
// 		if strings.HasPrefix(message, "[USERLIST]") {
// 			// fmt.Println("[USERLIST] received")
//             updateUsers(message)
//         } else {
//             // Regular message
//             app.QueueUpdateDraw(func() {
//                 fmt.Fprintf(messageBox, "%s", message)
//             })
//         }
//     }
// }

// func sendMessage(key tcell.Key) {
//     if key == tcell.KeyEnter {
//         text := inputField.GetText()
//         if strings.TrimSpace(text) != "" {
//             conn.Write([]byte(text + "\n"))
//             inputField.SetText("")
//         }
//     }
// }

// func updateUsers(message string) {
//     // fmt.Println("Updating users with message:", message) // Debug: Check what message is received

//     // Extract user list from the message
//     usersStr := strings.TrimPrefix(message, "[USERLIST]")
//     // userListText := "Active Users:\n"
//     usernames := strings.Split(strings.TrimSpace(usersStr), ",")
// 	// numActiveUsers := len(usernames)

//     // fmt.Println("Usernames:", usernames) // Debug: Check if usernames are correctly split

//     // Update userList TextView
// 	app.QueueUpdateDraw(func() {
//         userListText := fmt.Sprintf("Active Users: %d\n", len(usernames)-1)
//         userList.SetText(userListText)
//     })
// }


package main

import(
    "lan-party/client/network"
    "lan-party/client/ui"
)

func main(){
    network.Connect()
	ui.StartUI()
}