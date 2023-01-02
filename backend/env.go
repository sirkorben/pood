package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func setENV() {
	envFile, err := os.Open("./.env")
	if err != nil {
		log.Fatalln(err)
	}
	defer envFile.Close()
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		envVar := strings.Split(scanner.Text(), "=")
		os.Setenv(envVar[0], envVar[1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(os.Getenv("ENV_VAR"))
}
