package main

import (
	"os"

	"github.com/KevinLiddle/tic_tac_go/console"
)

func main() {
	console.Play(os.Stdin, os.Stdout)
}
