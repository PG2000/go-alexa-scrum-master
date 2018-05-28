package main

import (
	"go-scrum-master/alexa"
	"os"
)

var (
	arn string
)

func main() {
	arn = os.Getenv("ARN_FOR_START")
	if arn == "" {
		panic("NO Arn Given")
	}

	alexa.Start(arn)
}
