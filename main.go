package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	baseWord := flag.String("w", "", "Single word to use for concatenation")
	side := flag.String("side", "left", "String that determines the side of the single word")
	leftWordlistFileName := flag.String("l", "", "Wordlist file (left side)")
	rightWordlistFileName := flag.String("r", "", "Wordlist file (right side)")
	output := flag.String("o", "", "Output file (optional)")
	delimiter := flag.String("delimiter", "", "String delimiter to place between words")
	flag.Parse()

	go func() {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
		<-signalChannel

		fmt.Println("Program interrupted, exiting...")
		os.Exit(0)
	}()

	if *leftWordlistFileName == "" && *rightWordlistFileName == "" {
		fmt.Println("No input provided")
		os.Exit(1)
	}

	if *side != "left" && *side != "right" {
		fmt.Println("Side flag must be left or right!")
		os.Exit(0)
	}

	leftWordlist := make([]string, 0)
	rightWordlist := make([]string, 0)

	if *baseWord != "" {
		if *side == "left" {
			leftWordlist = append(leftWordlist, *baseWord)
		} else {
			rightWordlist = append(rightWordlist, *baseWord)
		}
	}

	if *leftWordlistFileName != "" {
		inputFile, err := os.Open(*leftWordlistFileName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer inputFile.Close()

		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			leftWordlist = append(leftWordlist, scanner.Text())
		}
	}

	if *rightWordlistFileName != "" {
		inputFile, err := os.Open(*rightWordlistFileName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer inputFile.Close()

		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			rightWordlist = append(rightWordlist, scanner.Text())
		}
	}

	if len(rightWordlist) == 0 || len(leftWordlist) == 0 {
		fmt.Println("No input provided")
		os.Exit(1)
	}

	var outputFile *os.File
	var err error
	if *output != "" {
		outputFile, err = os.Create(*output)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer outputFile.Close()
	}

	for _, l := range leftWordlist {
		for _, r := range rightWordlist {
			fmt.Println(l + *delimiter + r)
			if outputFile != nil {
				_, _ = outputFile.WriteString(l + *delimiter + r + "\n")
			}
		}
	}

}
