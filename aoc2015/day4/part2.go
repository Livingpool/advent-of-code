package day4

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

func Part2() {
	key := "ckczppom"

	i := 1
	for {
		input := key + strconv.Itoa(i)
		h := md5.New()
		io.WriteString(h, input)

		hashString := fmt.Sprintf("%x", h.Sum(nil))
		if hashString[:6] == "000000" {
			fmt.Println(i)
			break
		}
		i++
	}
}
