package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Flags is a logger struct that holds command line flags for customising the logging output
type Flags struct {
	DefinitionFile *string
	LogFile        *string
	IsDebug        *bool
}

var isDebug bool
var isWriteToFile bool
var logfile *os.File

// Init is a logger function that initialises the logger based on the given flags
func Init(flags Flags) {
	initFile(*flags.LogFile)
	initDebug(*flags.IsDebug)
}

func initFile(filename string) {
	if filename == "" {
		isWriteToFile = false
		return
	}

	isWriteToFile = true
	var err error
	logfile, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
}

func initDebug(debug bool) {
	isDebug = debug
}

// Close is a logger function that closes the log file
func Close() {
	if isWriteToFile {
		err := logfile.Close()
		if err != nil {
			panic(err)
		}
	}
}

// Logo is a logger function that prints the apidor logo
func Logo() {
	logo := `

	 █████╗ ██████╗ ██╗██████╗  ██████╗ ██████╗ 
	██╔══██╗██╔══██╗██║██╔══██╗██╔═══██╗██╔══██╗
	███████║██████╔╝██║██║  ██║██║   ██║██████╔╝
	██╔══██║██╔═══╝ ██║██║  ██║██║   ██║██╔══██╗
	██║  ██║██║     ██║██████╔╝╚██████╔╝██║  ██║
	╚═╝  ╚═╝╚═╝     ╚═╝╚═════╝  ╚═════╝ ╚═╝  ╚═╝
												
	`

	writeln(logo)
}

// RunInfo is a logger function that prints some base information about the current execution
func RunInfo(baseURI string, endpointsCount int, flags Flags) {
	writeln("API: " + baseURI)
	writeln("Endpoints: " + strconv.Itoa(endpointsCount))
	writeln("Time: " + time.Now().String())

	if *flags.DefinitionFile != "" {
		writeln("Definition: " + *flags.DefinitionFile)
	}
	if *flags.LogFile != "" {
		writeln("Log: " + *flags.LogFile)
	}
	if *flags.IsDebug {
		writeln("Debugging: on")
	}

	writeln("")
}

// Starting is a logger function that prints a start message
func Starting() {
	writeln("Starting...")
	writeln("")
}

// Finished is a logger function that prints a finished message
func Finished() {
	if !isDebug {
		writeln("")
	}
	writeln("Done, nice one 👊")
	writeln("")
}

// TestPrefix is a logger function that prints the endpoint and the test name that is taking place
func TestPrefix(endpoint string, testName string) {
	prefix := "[" + endpoint + "][" + testName + "] "
	if isDebug {
		writeln(prefix)
		writeln("")
	} else {
		write(prefix)
	}
}

// TestResult is a logger function that prints the result of the test
func TestResult(result string) {
	if !isDebug {
		writeln(result)
	}
}

// Message is a logger function that prints a given message
func Message(message string) {
	writeln(message)
}

// DebugMessage is a logger function that prints if the debug flag is set
func DebugMessage(message string) {
	if isDebug {
		writeln(message)
		writeln("")
	}
}

func write(message string) {
	fmt.Print(message)
	if isWriteToFile {
		logfile.WriteString(message)
	}
}

func writeln(message string) {
	fmt.Println(message)
	if isWriteToFile {
		logfile.WriteString(message + "\n")
	}
}