/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T11:18:02+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-17T00:35:26+07:00
 */

package main

import (
	"log"
	. "twido/config"
	. "twido/engine"
)

var runTwitterVideoDownloaderBotApps = func() {
	statusesToReplyBack, _ := CollectMentionsAsStatusesToReplyBack()
	NotifyUserTheVideoDownloadLink(statusesToReplyBack)
}

func main() {
	if nil != TwidoConfigErr {
		log.Println("Could not run apps because config file wasn't loaded.")
		return
	}

	runTwitterVideoDownloaderBotApps()
}
