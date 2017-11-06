// aky.alexa@gmail.com

package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type ticket struct {
	seat int
	id   string
}

func main() {

	fmt.Println("Launching server...")

	var availableTickets = 15

	var tickets = []ticket{}

	for i := 1; i <= availableTickets; i++ {
		newTicket := ticket{seat: i * 15, id: strconv.Itoa(i)}
		tickets = append(tickets, newTicket)
	}

	// listen on all interfaces
	ln1, _ := net.Listen("tcp", ":8081")
	ln2, _ := net.Listen("tcp", ":8082")

	// accept connection on port
	conn1, _ := ln1.Accept()
	conn2, _ := ln2.Accept()

	// run loop forever (or until ctrl-c)
	for {
		message1, _ := bufio.NewReader(conn1).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message1))
		var newMessage string

		receivedMess := string(message1)
		wantedTickets, err := strconv.Atoi(strings.TrimSpace(receivedMess))
		if err != nil {
			panic(err)
		}

		if len(tickets) == 0 {
			newMessage = "We are sorry! No more tickets!\n"
		} else if wantedTickets > len(tickets) {
			newMessage = "We are sorry! Only " + strconv.Itoa(len(tickets)) + " tickets left!\n"
		} else {
			newMessage = "All good!\n"
			for i := 0; i < wantedTickets; i++ {
				t := tickets[len(tickets)-1]
				newMessage = newMessage + "You have ticket id #" + t.id + ". Seat no. " + strconv.Itoa(t.seat) + ".\n"

				tickets = tickets[:len(tickets)-1]
				if i == wantedTickets-1 {
					newMessage = newMessage + "x"
				}
			}
		}

		conn1.Write([]byte(newMessage + "\n"))

		message2, _ := bufio.NewReader(conn2).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message2))
		var newMessage2 string

		receivedMess2 := string(message2)
		wantedTickets2, err := strconv.Atoi(strings.TrimSpace(receivedMess2))
		if err != nil {
			panic(err)
		}

		if len(tickets) == 0 {
			newMessage2 = "We are sorry! No more tickets!\n"
		} else if wantedTickets > len(tickets) {
			newMessage2 = "We are sorry! Only " + strconv.Itoa(len(tickets)) + " tickets left!\n"
		} else {
			newMessage2 = "All good!\n"
			for i := 0; i < wantedTickets2; i++ {
				t := tickets[len(tickets)-1]
				newMessage2 = newMessage2 + "You have ticket id #" + t.id + ". Seat no. " + strconv.Itoa(t.seat) + ".\n"

				tickets = tickets[:len(tickets)-1]
				if i == wantedTickets2-1 {
					newMessage = newMessage2 + "x"
				}
			}
		}

		conn2.Write([]byte(newMessage2 + "\n"))
	}
}
