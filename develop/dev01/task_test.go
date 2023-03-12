package dev01

import (
	"fmt"
	"testing"
	"time"
)

func TestShowTime(t *testing.T) {
	currentTime, err := CurrentTime()
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Second * 1)
	fmt.Printf("%d:%02d\n", currentTime.Hour(), currentTime.Minute())
}
