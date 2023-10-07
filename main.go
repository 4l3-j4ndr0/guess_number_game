package main

import (
	"fmt"
	"runtime"
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


	// ########  declaracion switch  ########

	// especificando la variable con la que se va a trabajar
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}

	// sin especificar la variable con la que se va a trabajar y simplemente
	// trabjando directamente con la variable
	switch os := runtime.GOOS; {
	case os == "windows":
		fmt.Println("Go run -> Windows")
	case os == "linux":
		fmt.Println("Go run -> Linux")
	case os == "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}

	// manera tradicional como en otros lenguajes
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}

	// Nota : No esta mal colocar el break al fial de cada caso pero no
	// es necesario

}
