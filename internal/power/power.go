package power

import (
	"log"
	"time"
	"strings"
	"strconv"
	"io/ioutil"
)

type PowerManager struct {
	nfCnt int
	Pch chan struct{}
}
func NewPM() *PowerManager {
	return &PowerManager{
		nfCnt: 0,
		Pch: make(chan struct{}, 1),
	}
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
	if pwr < 10 {
		return true
	}
	return false
}

func (pm *PowerManager) Timer(smode *bool) {
	ticker := time.NewTicker(1*time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if *smode && Check() && pm.nfCnt < 4 {
			pm.nfCnt++
			pm.Pch <- struct{}{}
			time.Sleep(1*time.Minute)
		} else if pm.nfCnt >= 4 {
			pm.nfCnt = 0
			time.Sleep(3*time.Minute)
		}
	}
}
