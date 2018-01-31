package cmd

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

/*
func prepareLogfile() {

		commandName, exPath := helper.GetExecutableInfo()

		val, _ := os.LookupEnv("COMMAND_PROXY_LOGGING")

		isLogginEnabled := (val == "1")

		if _, ok := conf.Commands[commandName]; ok {
			commandConfig = conf.Commands[commandName]
		}

		if conf.Commands[commandName].Logging || isLogginEnabled {

			fileName := path.Join(exPath, fmt.Sprintf("%s_proxy.log", commandName))

			var err error
			logfile, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				panic(err)
			}

			log.SetOutput(logfile)
		} else {
			log.SetOutput(ioutil.Discard)
		}

}

func buildAllInputString() string {
	bufferString := bytes.NewBufferString("")
	for _, val := range os.Args {
		bufferString.WriteString(val)
	}
	for _, val := range os.Environ() {
		bufferString.WriteString(val)
	}

	return bufferString.String()
}
*/
var logfile io.ReadCloser

func closeLogfile() {
	if logfile == nil {
		return
	}

	err := logfile.Close()
	if err != nil {
		log.Println(err)
	}
}

func logBlock(title string, f func()) {
	const totalHeaderWith = 40
	titleWidth := len(title)
	splitWidth := (totalHeaderWith - titleWidth/2)
	asciiart := bytes.Repeat([]byte{'-'}, splitWidth)
	log.Printf("%s%s%s", asciiart, title, asciiart)

	f()

	log.Printf("%s\n", bytes.Repeat([]byte{'-'}, splitWidth*2+titleWidth))
}

func logCommand(comandArgs []string) {
	logBlock(
		"docker command", func() {
			commandString := strings.Join(comandArgs, " ")
			log.Println(commandString)
		})
}

func initialLogging() {
	logBlock(
		"args",
		func() {
			for _, val := range os.Args {
				log.Println(val)
			}
		},
	)
	logBlock(
		"env",
		func() {
			for _, val := range os.Environ() {
				log.Println(val)
			}
		},
	)
}
