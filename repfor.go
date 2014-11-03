package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func runCommand(command string, t time.Time, stdOut chan<- string, stdError chan<- string) {
	args := strings.Split(command, " ")
	c := args[0]
	p := args[1:]
	cmd := exec.Command(c, p...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		stdError <- fmt.Sprintf("%v | %v >>>\n%v\n<<<\n", t.Format("Jan 2 15:04:05 2006"), command, err.Error())
	}
	stdOut <- fmt.Sprintf("%v | %v >>>\n%v\n<<<\n", t.Format("Jan 2 15:04:05 2006"), command, out.String())
}

func main() {
	fmt.Println("+==============+")
	fmt.Println("|REPeat FORever|")
	fmt.Println("+==============+")

	stdOut := make(chan string)
	stdError := make(chan string)
	var delay int64
	var command, outlog, errorlog string
	if len(os.Args) == 1 {
		fmt.Println("Usage: repfor [-d=3] [-o=/dev/stdout] [-e=/dev/stderr] commandToRepeatForeverWithDelayProvided")
		fmt.Println("Examples:\n # repfor 'cat /etc/issue' \n # repfor -d=5 -o=output.log -e=error.log 'ps -e'")
		fmt.Println("Visit https://github.com/vodolaz095/repfor for more information")
		os.Exit(1)
	}
	command = os.Args[len(os.Args)-1]
	flag.Int64Var(&delay, "d", 3, "Interval between executing command")
	flag.StringVar(&outlog, "o", "/dev/stdout", "Where to stream stdout")
	flag.StringVar(&errorlog, "e", "/dev/stderr", "Where to stream stderr")

	flag.Parse()
	fmt.Printf("We execute `%v` with interval of  %v seconds...\n", command, delay)
	fmt.Printf("StdOut goes to `%v`\n", outlog)
	fmt.Printf("StdError goes to `%v`\n", errorlog)

	outLogFile, err := os.OpenFile(outlog, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	outErrorFile, err := os.OpenFile(errorlog, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	defer func() {
		outErrorFile.Close()
		outLogFile.Close()
	}()

	timerChannel := time.Tick(time.Duration(delay) * time.Second)
	for {
		select {
		case outMessage := <-stdOut:
			outLogFile.Write([]byte(outMessage))
		case outError := <-stdError:
			outErrorFile.Write([]byte(outError))
		case t := <-timerChannel:
			go runCommand(command, t, stdOut, stdError)
		default:
			//do nothing
		}
	}
}
