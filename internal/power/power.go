package power

import (
	"log"
	"time"
	"strings"
	"strconv"
	"io/ioutil"
)

type PowerManager struct {
	NfCnt int
	Delay time.Duration
}
func NewPM() *PowerManager {
	return &PowerManager{NfCnt: 0}
}

func Show() string {
	data, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/capacity")
	if err != nil {
		log.Fatalln("Ошибка при попытке открыть файл по пути: /sys/class/power_supply/BAT0/capacity\n", err)
	}

	return strings.TrimSpace(string(data))
}

func Check() bool {
	pwr, _ := strconv.Atoi(Show())
	if pwr > 10 {
		return true
	}
	return false
}
