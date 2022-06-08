package system

var (
	version     string = "_unknown"
	buildCommit string = "_unknown"
	buildDate   string = "_unknown"
)

// Version get service version
func Version() string {
	return version
}

// BuildCommit get service build commit
func BuildCommit() string {
	return buildCommit
}

// BuildDate get service build date
func BuildDate() string {
	return buildDate
}
