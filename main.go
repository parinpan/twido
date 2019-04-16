/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T11:18:02+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-16T18:34:47+07:00
 */

package main

import (
	"log"
	. "twido/config"
	. "twido/engine"
)

var runTwitterVideoDownloaderApps = func() {
	statusesToReplyBack, _ := CollectMentionsAsStatusesToReplyBack()
	NotifyUserTheVideoDownloadLink(statusesToReplyBack)
}

func main() {
	if nil != TwidoConfigErr {
		log.Println("Could not run apps because config file wasn't loaded.")
		return
	}

	runTwitterVideoDownloaderApps()
}
