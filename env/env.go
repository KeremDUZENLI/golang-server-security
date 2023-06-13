package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DECIDE int

	URL  string
	PORT string

	CONCURRENCY int
	NUMREQUEST  int

	COUNTER int
	LISTE   []int
)

func Load() {
	printScan("WELCOME TO THE HELL")
	printScan("PRESS 1 FOR LOCAL TEST", &DECIDE)

	if DECIDE == 1 {
		loadEnv()
	} else {
		printScan("URL", &URL)
	}

	printScan("CONCURRENCY", &CONCURRENCY)
	printScan("NUMREQUEST", &NUMREQUEST)
	printScan("PRESS ENTER", nil)
}

func loadEnv() {
	godotenv.Load(".env")
	URL = os.Getenv("URL")
	PORT = os.Getenv("PORT")
}

func printScan(varName string, name ...any) {
	if len(name) == 0 {
		fmt.Printf("%s\n", varName)
	} else {
		switch value := name[0].(type) {
		case *string:
			fmt.Printf("\n%s: ", varName)
			fmt.Scanln(value)
		case *int:
			fmt.Printf("\n%s: ", varName)
			fmt.Scanln(value)
		case nil:
			fmt.Printf("\n%s ", varName)
			fmt.Scanln()
		}
	}
}
