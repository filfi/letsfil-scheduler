package timeUtils

import (
	"fmt"
	"testing"
	"time"
)

func TestMonthStart(t *testing.T) {
	fmt.Println(GetMonthStart(time.Now()))
	fmt.Println(GetMonthEnd(time.Now()))
	fmt.Println(GetHourEndStr(time.Now(), 0))
}
