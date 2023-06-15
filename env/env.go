package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	URL  string
	PORT string

	CONCURRENCY int
	NUMREQUEST  int

	COUNTER int = 0
	LISTE   []int
	LOOPER  bool = false
)

func PrintScan(varName string, name ...any) {
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

func LoadLocalEnvFile() {
	godotenv.Load(".env")
	URL = os.Getenv("URL")
	PORT = os.Getenv("PORT")
}
