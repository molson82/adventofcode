package utils

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}
