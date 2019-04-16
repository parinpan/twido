/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T15:07:19+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-16T16:06:41+07:00
 */

package dataprovider

type ReplyBody struct {
	OriginalStatus Status
	VideoVariant   VideoVariant
}

type StatusesToReplyBack struct {
	Queue []ReplyBody
}

func (strb *StatusesToReplyBack) AddToQueue(rb ReplyBody) {
	strb.Queue = append(strb.Queue, rb)
}
