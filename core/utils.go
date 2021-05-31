package core

import (
	"time"
)

func Now() uint {
	return uint(time.Now().UnixNano() / int64(time.Millisecond))
}

func Wait(start uint, fps uint) uint {
	speed := uint(1000 / fps)
	delta := Now() - start
	if delta >= speed {
		return 0
	}

	waitTime := speed - delta

	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	return waitTime
}
