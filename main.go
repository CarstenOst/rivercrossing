package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var venstre []string
var noahsArk []string
var høyre []string

func main() {

	venstre := []string{"Rev", "Kylling", "Mann", "Korn"}
	noahsArk := []string{}
	høyre := []string{}

	inputStuff()
	showState(venstre, noahsArk, høyre)

	fmt.Printf("Flytter %v fra venstre til Noahs Ark", venstre)
	sleepyBoy(4000)
	clearTerminal()

	venstre = []string{}
	noahsArk = []string{"Rev", "Kylling", "Mann", "Korn"} // jajaja burde lære meg slice og pop osv...
	showState(venstre, noahsArk, høyre)

	fmt.Printf("Flytter %v fra Noahs Ark til høyre", noahsArk)
	sleepyBoy(4000)
	clearTerminal()

	noahsArk = []string{}
	høyre = []string{"Rev", "Kylling", "Mann", "Korn"}
	showState(venstre, noahsArk, høyre)

	exit()
}

func GjørIngenTing(enString string) string {
	return enString
}

func showState(venstre []string, noahsArk []string, høyre []string) {
	state(venstre, noahsArk, høyre)
	sleepyBoy(3000)
	clearTerminal()
}

func state(venstre []string, noahsArk []string, høyre []string) {
	fmt.Println("Status: ")
	fmt.Printf("Til venstre: %v", venstre)
	fmt.Println("") // hakke tid til å lære mer nå ass - Carsten 28.01.2022 22:42
	fmt.Printf("I Noahs Ark: %v", noahsArk)
	fmt.Println("")
	fmt.Printf("Til høyre: %v", høyre)
}

func inputStuff() {
	clearTerminal()
	scanner := bufio.NewScanner(os.Stdin) // Scanner objekt ugly yes stfu
	fmt.Println("")
	fmt.Printf("Skriv noe for å starte animasjon: ") // tar litt for lang tid å lage et spill
	scanner.Scan()                                   // input fra bruker
	input := scanner.Text()                          // lagrer input fra bruker som string
	fmt.Printf("Litt rart å skrive %q synes jeg ", input)
	sleepyBoy(1000)
	clearTerminal()
}

// Usefull stuff
func clearTerminal() { // obv
	fmt.Print("\033[H\033[2J") // What the fudge?
}

func sleepyBoy(milli int) { // obv
	time.Sleep(time.Duration(milli) * time.Millisecond)
}

func dotdotdot(nrDots int, milli int) { // fancy dot dot dot animasjon before exiting
	for i := 0; i < nrDots; i++ {
		sleepyBoy(milli)
		fmt.Printf(".")
	}
}

// useless shit
func exit() { // do I really need to explain?
	clearTerminal()
	fmt.Printf("Wait, somethings not right ")
	dotdotdot(50, 20)
	clearTerminal()
	os.Exit(0) // exit as succsess
}
