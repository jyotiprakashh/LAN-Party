package main

import(
    "lan-party/client/network"
    "lan-party/client/ui"
)

func main(){
    network.Connect()
	ui.StartUI()
}
