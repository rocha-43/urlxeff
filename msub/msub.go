package msub

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Msub(targetsList string, wordlist string, outName string) {

	file, err := os.Open(wordlist)
	if err != nil {
		log.Fatal("[!] Can't open the wordlist...")
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var payload []string

	for scanner.Scan() {
		payload = append(payload, scanner.Text())
	}

	var targets []string

	twl, err := os.Open(targetsList)
	if err != nil {
		log.Fatal("[!] Can't open the target list...")
		os.Exit(1)
	}

	tscanner := bufio.NewScanner(twl)
	tscanner.Split(bufio.ScanWords)

	for tscanner.Scan() {
		targets = append(targets, tscanner.Text())
	}

	out := fmt.Sprintf("%s.txt", outName)

	outFile, err := os.Create(out)
	if err != nil {
		log.Fatal("Can't create the output file...")
		os.Exit(1)
	}

	defer outFile.Close()

	for _, base := range targets {
		for _, cont := range payload {
			value := fmt.Sprintf("http://%s.%s\n", cont, base)
			_, err := outFile.WriteString(value)
			if err != nil {
				log.Fatal("[!] Can't write the output")
			}
		}
	}
}
