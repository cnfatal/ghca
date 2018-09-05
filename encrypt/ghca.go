package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

//GhcaEncode 算号
func GhcaEncode(username, password string) string {
	const key = "aI0fC8RslXg6HXaKAUa6kpvcAXszvTcxYP8jmS9sBnVfIqTRdJS1eZNHmBjKN28j"
	pswLen := len(password)
	username = strings.ToUpper(username)
	timeUnix := int(time.Now().Unix())
	timeHex := fmt.Sprintf("%x", timeUnix)

	accountAsciiSum := 0
	for _, v := range username {
		accountAsciiSum += int(v)
	}
	timeModPsw := timeUnix % pswLen
	timeMod := addOne(timeModPsw, timeModPsw < 1)
	seed := addOne(timeMod-1, timeModPsw != pswLen)
	splitLen := addOne(timeMod, timeMod != pswLen)

	usernameLastFour := fmt.Sprintf("%04x", accountAsciiSum^seed)

	splitPswOne := password[0:splitLen]
	splitPswTwo := password[splitLen:]
	splitKeyOne := key[0 : 60-splitLen]
	splitKeyTwo := key[0 : 64-len(username)-pswLen+splitLen]

	md5Str := splitKeyOne + splitPswOne + username + splitKeyTwo + splitPswTwo

	bytes := make([]byte, 0)
	bytes = append(bytes, byte(timeUnix>>24), byte(timeUnix>>16), byte(timeUnix>>8), byte(timeUnix>>0))
	bytes = append(bytes, []byte(md5Str)...)

	// md5 twice
	hash := md5.New()
	hash.Write(bytes)
	sum := hash.Sum(nil)
	hash.Reset()
	hash.Write(sum)
	md5twice := hex.EncodeToString(hash.Sum(nil))[:16]

	return "~ghca" + strings.ToUpper(timeHex+"2023"+md5twice+usernameLastFour+username)
}

func addOne(in int, isAddOne bool) int {
	if isAddOne {
		in++
	}
	return in
}
