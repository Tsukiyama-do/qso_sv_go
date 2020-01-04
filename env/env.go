package env

import (
	"log"
	"os"
)

var smode bool = true

func S_mode() bool {
	return smode
}

func S_mset(b bool) {
	smode = b
	log.Printf("Mode is changed to %t. \n", b)
}

// return host name depending on runnning environments
func S_host() string {
	if s_hostname, _ := os.Hostname(); s_hostname == "yuichi-x220" {
		return "localhost:"
	} else {
		return "jj1pow.com:"
	}
}
