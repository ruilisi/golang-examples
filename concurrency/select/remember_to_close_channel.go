package main

import (
	"fmt"
	"strconv"
)

func makeCakeAndSend(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := flavor + " Cake " + strconv.Itoa(i)
		cs <- cakeName //send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack(strbry_cs chan string, choco_cs chan string) {
	var cakeName string
	strbry_ok, choco_ok := true, true

	for {
		//if both channels are closed then we can stop
		if !strbry_ok && !choco_ok {
			break
		}
		fmt.Println("Waiting for a new cake ...")
		select {
		case cakeName, strbry_ok = <-strbry_cs:
			if strbry_ok {
				fmt.Println("Received from Strawberry channel.  Now packing", cakeName)
			}
		case cakeName, choco_ok = <-choco_cs:
			if choco_ok {
				fmt.Println("Received from Chocolate channel.  Now packing", cakeName)
			}
		}
	}
}

func main() {
	strbry_cs := make(chan string)
	choco_cs := make(chan string)

	//two cake makers
	go makeCakeAndSend(choco_cs, "Chocolate", 3)   //make 3 chocolate cakes and send
	go makeCakeAndSend(strbry_cs, "Strawberry", 4) //make 3 strawberry cakes and send

	//one cake receiver and packer
	receiveCakeAndPack(strbry_cs, choco_cs) //pack all cakes received on these cake channels
}
