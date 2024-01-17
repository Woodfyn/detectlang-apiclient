package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Woodfyn/detectlang-apiclient/detectlanguage/detectlanguage"
)

func main() {
	myClient, err := detectlanguage.NewClient(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	lang, err := myClient.GetLanguage("THAI")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(detectlanguage.Info(lang))

	langForAll, err := myClient.GetLanguages()
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range langForAll {
		fmt.Println(detectlanguage.Info(i))
	}

	accst, err := myClient.AccountStatus("a3d1067f9e52be7296b5e41715327bcf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(detectlanguage.Info(accst))

	wordOne, err := myClient.SingleDetect("a3d1067f9e52be7296b5e41715327bcf", "Hello+world")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(detectlanguage.Info(wordOne))

}
