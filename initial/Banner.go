package initial

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func hash(s string) string {
	h := md5.New()
	f, err := os.Open(s)
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

func Check_for_hash(s string) bool {
	switch s {
	case "standard.txt":
		{
			if hash(s) == "ac85e83127e49ec42487f272d9b9db8b" {
				return true
			}
		}
	case "shadow.txt":
		if hash(s) == "a49d5fcb0d5c59b2e77674aa3ab8bbb1" {
			return true
		}
	case "thinkertoy.txt":
		if hash(s) == "db448376863a4b9a6639546de113fa6f" {
			return true
		}
	}
	return false
}
