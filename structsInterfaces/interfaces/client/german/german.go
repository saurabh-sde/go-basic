package german

import "fmt"

type German struct {
	Name     string
	Language string
}

func (e German) NewClient(name string) *German {
	e = German{}
	e.Name = name
	e.Language = "German"
	return &e
}

func (e *German) GetClientName() string {
	return e.Name
}

func (e *German) GreetClient() {
	fmt.Printf("Hi! Welcome Client: %s. We are connecting you to %s speaker\n", e.Name, e.Language)
}
