package gabbymux

import (
	"errors"
	"fmt"
	_ "fmt"
	"net/http"
)

func Router() {
	// This is the default router for Gabbymux
}

func ListenAndServe(port string) {
	portIsTaken := isPortTaken("8000")
	if !portIsTaken {

	} else {
		errors.New("Port has been taken!")
	}
	fmt.Println("Hey Buddy, I just want to satisfy this formatter")
	http.ListenAndServe(port)
}
