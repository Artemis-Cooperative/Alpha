package alpha

import (
	"fmt"

	beta "github.com/Artemis-Cooperative/beta"
	packages "github.com/Artemis-Cooperative/interfaces/packages"
)

var betaInstance packages.Beta = beta.Beta{}

type Alpha struct{}

func (a Alpha) Run() {
	fmt.Println("alpha")
	betaInstance.Run()
}
