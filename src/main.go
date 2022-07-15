package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func Decode(sd []byte) (res string) {
	for _, v := range string(sd) {
		res = string(v) + res
	}
	res = strings.Join(strings.Split(res, ":"), " ")
	return
}

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	whatIsIt = Decode(sd)
	fmt.Println(whatIsIt)
}
