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


	// ########  declaracion bucle for  ########

	// con variable inicializada fuera del bucle
	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// con variable inicializada dentro del bucle
	for i := 1 ; i <= 10 ; i++ {
		fmt.Println(i)
	}

	// usando break para termina el programa 
	for i := 1 ; i <= 10 ; i++ {
		fmt.Println(i)
		if i ==5 {
			break  
		}
	}

	// usando continue para saltar a la siguiente iteracion   
	for i := 1 ; i <= 10 ; i++ {
		if i ==5 {
			continue  
		}
		fmt.Println(i)
	}

	// ########  LLAMANDO FUNCIONES PROPIAS  ########

	hello()

	hello2("Ale")

	temp := hello3("Alejo")
	fmt.Printf("Hola %s, soy una funcion propia con parametro usando return \n",temp)

	temp2 := calc(2,3)
	fmt.Printf("El calculo es %d \n", temp2)

	suma, multiplicacion := calc2(2,3)
	fmt.Printf("la suma es %d y la multiplicacion es %d \n", suma, multiplicacion)

	// otra forma de la funcion de arriba
	suma2, multiplicacion2 := calc3(2,3)
	fmt.Printf("la suma es %d y la multiplicacion es %d \n", suma2, multiplicacion2)
}


// ########  CREANDO FUNCIONES PROPIAS  ########

// funcion simple
func hello() {
		fmt.Println("Hola , soy una funcion propia simple")
}

// funcion con parametro
func hello2(name string) {
	fmt.Printf("Hola %s, soy una funcion propia con parametro \n",name)
}

// funcion con parametro y devolucion
func hello3(name string) string {
	return name
}

// funcion con varios parametros y devolucion de 1 valor
func calc(a, b int) int {
	return a+b
}

// funcion con varios parametros y devolucion de 2 valores
func calc2(a, b int) (int , int) {
	suma := a+b
	multiplicacion := a*b
	return suma, multiplicacion
}

// otra variante de la misma funcion de arriba
func calc3(a, b int) (suma , multiplicacion int) {
	suma = a+b
	multiplicacion = a*b
	return  // no es necesario espesificar q devuelve pq se especifico en los parametros
}