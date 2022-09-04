package global

// To set test and production, make it easy to change

var (
	debug bool = true
)

func IsDebugMode() bool {
	return debug
}
