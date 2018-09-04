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

	// fixme: the results is different from the PHP's chr() function, i don't know why
	i0 := string(timeUnix >> 24)
	i1 := string(timeUnix >> 16)
	i2 := string(timeUnix >> 8)
	i3 := string(timeUnix >> 0)

	accountAsciiSum := 0
	for _, v := range username {
		accountAsciiSum += int(v)
	}
	timeModPsw := timeUnix % pswLen
	timeMod := 0
	if timeModPsw < 1 {
		timeMod = timeModPsw + 1
	} else {
		timeMod = timeModPsw
	}
	seed := 0
	if timeModPsw == pswLen {
		seed = timeMod - 1
	} else {
		seed = timeMod
	}
	accountLatFour := fmt.Sprintf("%04x", accountAsciiSum^seed)
	splitLen := 0
	if timeMod == pswLen {
		splitLen = timeMod
	} else {
		splitLen = timeMod + 1
	}
	splitPswOne := string([]rune(password)[0:splitLen])
	splitPswTwo := string([]rune(password)[splitLen:pswLen])
	splitKeyOne := string([]rune(key)[0 : 60-len(splitPswOne)])
	splitKeyTwo := string([]rune(key)[0 : 64-len(username)-len(splitPswTwo)])
	md5Str := fmt.Sprintf("%s%s%s%s%s%s%s%s%s", i0, i1, i2, i3, splitKeyOne, splitPswOne, username, splitKeyTwo, splitPswTwo)
	md5twice := hex.EncodeToString(MD5(MD5([]byte(md5Str))))
	return "~ghca" + strings.ToUpper(timeHex+"2023"+md5twice+accountLatFour+username)
}

func MD5(bytes []byte) []byte {
	hash := md5.New()
	hash.Write(bytes)
	sum := hash.Sum(nil)
	return sum
}
