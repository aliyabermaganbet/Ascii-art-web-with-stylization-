package ascii

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Ascii_art(arg, banner string) (string, int) {
	t := "ascii/" + banner + ".txt"
	file, err := os.Open(t)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewScanner(file)
	var array []string
	for r.Scan() {
		array = append(array, r.Text())
	}
	if Check_the_hash(banner + ".txt") {
		if Check_the_argument(arg) {
			return To_ascii(array, arg), 0
		} else {
			return http.StatusText(http.StatusBadRequest), http.StatusBadRequest
		}
	} else {
		return http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError
	}
}

func To_ascii(array []string, arg string) string {
	myMap := make(map[rune][]string)
	var q rune = 32
	for i := 1; i < len(array); i += 9 {
		myMap[q] = array[i : i+8]
		q++
	}
	str := ""
	k := strings.ReplaceAll(arg, "\r\n", "\n")
	rk := strings.Split(k, "\n")
	for _, rn := range rk {
		if rn == "" {
			str += "\n"
		} else {
			for i := 0; i < 8; i++ {
				for d := 0; d < len(rn); d++ {
					str += myMap[rune(rn[d])][i]
				}
				str += "\n"
			}
		}
	}
	return str
}

func Check_the_hash(s string) bool {
	switch s {
	case "standard.txt":
		if Hash(s) == "ac85e83127e49ec42487f272d9b9db8b" {
			return true
		}
	case "thinkertoy.txt":
		if Hash(s) == "db448376863a4b9a6639546de113fa6f" {
			return true
		}
	case "shadow.txt":
		if Hash(s) == "a49d5fcb0d5c59b2e77674aa3ab8bbb1" {
			return true
		}
	}
	return false
}

func Hash(s string) string {
	h := md5.New()
	f, err := os.Open("ascii/" + s)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	a := fmt.Sprintf("%x", h.Sum(nil))
	return a
}

func Check_the_argument(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] < 32 || s[i] > 126) && s[i] != 10 && s[i] != 13 {
			return false
		}
	}
	return true
}
