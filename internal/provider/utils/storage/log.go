package storage

import (
	"fmt"
	"time"
)

type Logger interface {
	Info(s string)
}

func DiffVersionMsg(provider string, diff int) (message string) {
	message = provider + " servers from file discarded because they are " +
		fmt.Sprint(diff) + " version"
	if diff > 1 {
		message += "s"
	}
	message += " behind"
	return message
}

func DiffTimeMsg(provider string, persistedUnix,
	hardcodedUnix int64) (message string) {
	diff := time.Unix(persistedUnix, 0).Sub(time.Unix(hardcodedUnix, 0))
	if diff < 0 {
		diff = -diff
	}
	diff = diff.Truncate(time.Second)
	message = "Using " + provider + " servers from file which are " +
		diff.String() + " more recent"
	return message
}
