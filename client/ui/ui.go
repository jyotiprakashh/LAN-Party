package ui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"lan-party/client/network"
)

var (
	messageBox *tview.TextView
	inputField *tview.InputField
	app        *tview.Application
	userList   *tview.TextView
)

func StartUI() {
	app = tview.NewApplication()

	header := tview.NewTextView().
		SetText("Welcome to LAN Party !!!").
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorYellow)
		
	userList = tview.NewTextView().
		SetTextAlign(tview.AlignRight).
		SetText("Active Users: 0\n").
		SetTextAlign(tview.AlignLeft)

	messageBox = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetTextColor(tcell.ColorGreen)

	inputField = tview.NewInputField().
		SetLabel("Enter username: ").
		SetFieldWidth(0).
		SetDoneFunc(handleInput)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(header, 1, 1, false).
		AddItem(messageBox, 0, 1, false).
		AddItem(userList, 1, 1, false).
		AddItem(inputField, 1, 0, true)

	if err := network.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	go network.ReadMessages(updateUsers, updateMessages, handleError)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}


func handleInput(key tcell.Key) {
	if key == tcell.KeyEnter {
		text := inputField.GetText()
		if strings.TrimSpace(text) != "" {
			network.SendMessage(text + "\n")
			inputField.SetText("")
			inputField.SetLabel("Enter message: ")
			inputField.SetDoneFunc(handleMessage)
		}
	}
}

func handleMessage(key tcell.Key) {
	if key == tcell.KeyEnter {
		text := inputField.GetText()
		if strings.TrimSpace(text) != "" {
			network.SendMessage(text + "\n")
			inputField.SetText("")
		}
	}
}

func updateMessages(message string) {
	app.QueueUpdateDraw(func() {
		fmt.Fprintf(messageBox, "%s", message)
	})
}

func updateUsers(userCount int) {
	app.QueueUpdateDraw(func() {
		userList.SetText(fmt.Sprintf("Active Users: %d\n", userCount))
	})
}

func handleError(err error) {
	fmt.Println(err)
	app.Stop()
}
