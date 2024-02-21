package english

import "fmt"

type English struct {
	Name     string
	Language string
}

func (e English) NewClient(name string) *English {
	e = English{}
	e.Name = name
	e.Language = "English"
	return &e
}

func (e *English) GetClientName() string {
	return e.Name
}

func (e *English) GreetClient() {
	fmt.Printf("Hi! Welcome Client: %s. We are connecting you to %s speaker\n", e.Name, e.Language)
}
