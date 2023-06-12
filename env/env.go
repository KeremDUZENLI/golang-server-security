package env

var (
	URL         string
	PORT        string
	CONCURRENCY int
	NUMREQUEST  int
)

func Load() {
	URL = "http://localhost:8080/"
	PORT = ":8080"
	CONCURRENCY = 100
	NUMREQUEST = 1000
}
