package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	ta := time.Now()
	fmt.Println(ta.Format("2006-01-02T15:04:05-07:00"))
}
