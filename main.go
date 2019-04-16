/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T11:18:02+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-16T13:27:30+07:00
 */

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
