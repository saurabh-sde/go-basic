package main

import (
	"fmt"

	"github.com/saurabh-sde/go-basic/structsInterfaces/interfaces/client/english"
	"github.com/saurabh-sde/go-basic/structsInterfaces/interfaces/client/german"
)

type Client interface {
	GetClientName() string
	GreetClient()
}

func NewClient(c Client) {
	fmt.Println("Setup done for Client: ", c.GetClientName())
	c.GreetClient()

}

func main() {
	eng := (&english.English{}).NewClient("ABCD")
	NewClient(eng)

	german := (&german.German{}).NewClient("PQRS")
	NewClient(german)

}
