package main

import (
	log "github.com/sirupsen/logrus"
)

func fatal(err error, level ...int) {
	if err != nil {
		if len(level) == 0 || level[0] == 0  {
			panic(err)
		} else if level[0] == 1 {
			log.Error(err)
		} else if level[0] > 1 {
			log.Info(err)
		}
	}
}
