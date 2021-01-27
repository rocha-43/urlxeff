package sdirectory

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Sdirectory(base string, wordlist string, outName string) {

	file, err := os.Open(wordlist)
	if err != nil {
		log.Fatal("[!] Can't open the wordlist...")
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	out := fmt.Sprintf("%s.txt", outName)

	outFile, err := os.Create(out)
	if err != nil {
		log.Fatal("Can't create the output file...")
		os.Exit(1)
	}

	defer outFile.Close()

	for scanner.Scan() {
		value := fmt.Sprintf("http://%s/%s\n", base, scanner.Text())
		_, err := outFile.WriteString(value)
		if err != nil {
			log.Fatal("[!] Can't write to the output file...")
		}
	}
}
