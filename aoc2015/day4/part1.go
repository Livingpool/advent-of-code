package day4

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

func Part1() {
	key := "ckczppom"

	i := 1
	for {
		input := key + strconv.Itoa(i)
		h := md5.New()
		io.WriteString(h, input)

		hashString := fmt.Sprintf("%x", h.Sum(nil))
		if hashString[:5] == "00000" {
			fmt.Println(i)
			break
		}
		i++
	}
}
