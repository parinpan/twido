package main

import (
	"fmt"
	"log"
	. "twido/client/twitter"
	. "twido/config"
)

func main() {
	if nil != TwidoConfigErr {
		log.Println("Could not run apps because config file wasn't loaded.")
		return
	}

	service := TwitterService{}

	result, err := service.Statuses.Show(map[string]string{
		"id": "1117989996189945857",
	})

	if nil != err {
		log.Println(err.Error())
	}

	fmt.Printf("%#v\n", result)
}
