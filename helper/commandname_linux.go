package helper

func GetCommandName() string {
	return commandFileName
}

func ParseCommandNameFromCommandLine() string {
	return filepath.Base(os.Args[0])
}

func GetCommandNameFilename(commandName string) string {
	return GetCommandNameFilename
}
