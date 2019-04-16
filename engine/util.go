/**
 * @Author: Fachrin Aulia Nasution <nakama>
 * @Date:   2019-04-17T00:10:48+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-17T00:35:21+07:00
 */

package engine

import (
	"strings"
)

func BuildStringFromFormat(format string, dictionary map[string]string) string {
	sentence := format

	for key, val := range dictionary {
		sentence = strings.Replace(sentence, "{"+key+"}", val, -1)
	}

	return sentence
}
