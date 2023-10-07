package main

import (
	"fmt"
	"time"
)

func main() {

	// ########  declaracion if  ########

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Mañana")
	} else if t.Hour() < 17 {
		fmt.Println("Tarde")
	} else {
		fmt.Println("Noche")
	}
	
}
