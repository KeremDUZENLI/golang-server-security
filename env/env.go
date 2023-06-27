package env

import (
	"os"

	"seguro/common"

	"github.com/joho/godotenv"
)

var (
	NUMBERREQUEST int
	CONCURRENCY   int
	URL           string
	PORT          string
)

func LoadValuesGiven() {
	common.PrintScan("WELCOME TO THE HELL")

	common.PrintScan("NUMBERREQUEST", &NUMBERREQUEST)
	common.PrintScan("CONCURRENCY", &CONCURRENCY)
	common.PrintScan("URL (For Local Empty)", &URL)
}

func LoadValuesEnvFile() {
	godotenv.Load(".env")
	URL = os.Getenv("URL")
	PORT = os.Getenv("PORT")
}
