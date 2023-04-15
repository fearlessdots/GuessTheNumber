package main

import (
	// Modules in GOROOT
	"fmt"
	"os"
	"math/rand"
	"time"

	// External modules
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/buger/goterm"
)

//
//// DISPLAY FUNCTIONS
//

func showInfo(msg string) {
    gray := color.New(color.FgHiBlack)
    gray.Println(msg)
}

func showAttention(msg string) {
    orange := color.New(color.FgHiYellow)
    orange.Println(msg)
}

func showSuccess(msg string) {
    color.Blue(msg)
}

func showError(msg string) {
    color.Red(msg)
}

func hr(factor float64) {
	// Get terminal width
	width, _, err := terminalSize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
		return
	}

	// Print horizontal rule
	fmt.Println("")
	for i := 0; i < int(float64(width)*factor); i++ {
		fmt.Print("-")
	}
	fmt.Println("")
	fmt.Println("")
}

//
//// COMPLEMENTARY FUNCTIONS
//

func finishProgram(code int) {
    os.Exit(code)
}

func terminalSize() (int, int, error) {
	terminalHeight := goterm.Height()
	terminalWidth := goterm.Width()

	return terminalWidth, terminalHeight, nil
}

//
////
//

type Game struct {
	victories   int
	failures    int
	vicAlerts   []string
	faiAlerts   []string
	levels      []Level
	level 		Level
	number      int
	max_chances	int
	chances     int
}

type Level struct {
	name		string
	max_number	int
}

func NewGame() *Game {
	g := &Game{
		victories: 0,
		failures:  0,
		vicAlerts: []string{
			"You are very good!",
			"Super!",
			"Congratulations, you won!",
			"You're a winner, way to go!",
			"That's it, you got it!",
			"Excellent job, you nailed it!",
			"Bravo, you solved it!",
			"Outstanding, you are correct!",
			"Amazing, you're a genius!",
			"You're a natural at this.",
		},
		faiAlerts: []string{
			"Nope, that's not it.",
			"Not even close, try again.",
			"Oops, better luck next time!",
			"Sorry, that's not quite right.",
			"That's a miss.",
			"A valiant effort, but no.",
			"You're not quite there yet.",
		},
		levels: []Level{
			{
				name: "Easy",
				max_number: 15,
			},
			{
				name: "Medium",
				max_number: 25,
			},
			{
				name: "Hard",
				max_number: 35,
			},
		},
	}
	return g
}

func (g *Game) menu() {
	// Max chances
	maxchances_prompt := []*survey.Question{
		{
			Name: "custom max chances",
			Prompt: &survey.Confirm{
				Message: "Do you want to use the default max number of chances (5)?",
			},
		},
	}
	
	var choice bool
	
	err := survey.Ask(maxchances_prompt, &choice)
	if err != nil {
		showAttention("Error displaying menu: " + err.Error())
		finishProgram(1)			
	}

	if !choice{
		// Custom max chances
		customchances_prompt := &survey.Input{
			Message: "Insert a custom max chances:",
		}

		var number int

		err := survey.AskOne(customchances_prompt, &number)
		if err != nil {
			showAttention("Error displaying menu: " + err.Error())
			finishProgram(1)			
		}

		g.max_chances = number
	} else {
		g.max_chances = 5
	}
	
	// Game level
	options := make([]string, len(g.levels) + 1)
	for i, level := range g.levels {
		options[i] = level.name
	}
	options[len(g.levels)] = "Custom"

	prompt := &survey.Select{
		Message: "Select a game level:",
		Options: options,
	}

	var option int
	err = survey.AskOne(prompt, &option)
	if err != nil {
		showAttention("Error displaying menu: " + err.Error())
		finishProgram(1)
	}

	if option == len(g.levels) {
		// Custom game level
		prompt := &survey.Input{
			Message: "Insert a custom max number:",
		}

		var number int

		err := survey.AskOne(prompt, &number)
		if err != nil {
			showAttention("Error displaying menu: " + err.Error())
			finishProgram(1)			
		}

		g.level = Level{
			name: "Custom",
			max_number: number,
		}
	} else {
		// Set game level
		g.level = g.levels[option]
	}
	
	// Run game
	g.run()
}

func (g *Game) run() {
	// Generate random number
	g.number = rand.Intn(g.level.max_number) + 1

	g.chances = 0


	fmt.Println("")
	showAttention(fmt.Sprintf("> Guess a number between 1 and %d.\n", g.level.max_number))
	fmt.Println("")

	for g.chances < g.max_chances {
		var guess int
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		if guess == g.number {
			g.victories++
			hr(0.2)
			g.alert(true)
			fmt.Println("")
			fmt.Printf("You won: %d times\n", g.victories)
			fmt.Printf("You lost: %d times\n", g.failures)
			hr(0.2)
			g.restart()
		} else if guess < g.number {
			hr(0.2)
			fmt.Println("Your guess was too low")
		} else {
			hr(0.2)
			fmt.Println("Your guess was too high")
		}
		g.chances++
		fmt.Printf("Chances left: %d\n", g.max_chances-g.chances)
		hr(0.2)
	}

	g.failures++

	hr(0.2)
	g.alert(false)
	fmt.Println("")
	fmt.Printf("You won: %d times\n", g.victories)
	fmt.Printf("You lost: %d times\n", g.failures)
	hr(0.2)
	g.restart()
}

func (g *Game) alert(win bool) {
	var alerts []string
	if win {
		alerts = g.vicAlerts
	} else {
		alerts = g.faiAlerts
	}
	
	alert := alerts[rand.Intn(len(alerts))]

	if win {
		showSuccess(alert)
	} else {
		showError(alert)
	}
}

func (g *Game) restart() {
	prompt := []*survey.Question{
		{
			Name: "restart game",
			Prompt: &survey.Confirm{
				Message: "Play again?",
			},
		},
	}
	
	var choice bool
	
	err := survey.Ask(prompt, &choice)
	if err != nil {
		showAttention("Error displaying menu: " + err.Error())
		finishProgram(1)			
	}

	if choice {
		hr(0.5)
		fmt.Println("Starting a new game...")
		g.menu()
	} else {
		fmt.Println("")
		finishProgram(0)
	}
}

func main() {
    // Seed the random number generator with the current time
    rand.Seed(time.Now().UnixNano())
    
	game := NewGame()
	game.menu()
}
