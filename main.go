package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	joueur := charCreation()
	for {
		joueur.dead()
		handleMenu(joueur)
	}
}

