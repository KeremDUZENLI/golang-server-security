package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	URL         string
	PORT        string
	CONCURRENCY int
	NUMREQUEST  int
)

func Load() {
	godotenv.Load(".env")

	URL = "http://localhost:" + os.Getenv("URL")
	PORT = ":" + os.Getenv("PORT")
	CONCURRENCY, _ = strconv.Atoi(os.Getenv("CONCURRENCY"))
	NUMREQUEST, _ = strconv.Atoi(os.Getenv("NUMREQUEST"))
}
