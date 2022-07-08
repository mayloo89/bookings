// Change package name to main for running
package myapp

import (
	"errors"
	"fmt"
	"net/http"
)

// Globals
const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page.")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is the About page and 2 + 2 is %d", sum))
}

// Divide is the divide page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 10.0

	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%.2f divided by %.2f is %.2f", x, y, f))
}

// addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

// divideValues divide two float64 and returns the result float64
func divideValues(x, y float64) (float64, error) {
	result := x / y
	if y < 0 {
		err := errors.New("Cannot divide by 0")
		return 0, err
	}
	return result, nil
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Staring application... \nListening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
