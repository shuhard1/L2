package dev01

import (
	"time"

	"github.com/beevik/ntp"
)

// CurrentTime возвращает текущее время
func CurrentTime() (time.Time, error) {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}
	time := time.Now().Add(response.ClockOffset)
	return time, err
}
