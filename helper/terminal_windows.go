package helper

// IsTerminalContext checks if the current process is in a terminal context
// For windows terminal detection is currently not implement.
func IsTerminalContext() bool {
	return true
}
