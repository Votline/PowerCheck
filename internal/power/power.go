package power

import (
	"log"
	"strings"
	"io/ioutil"
)

func Show() string {
	data, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/capacity")
	if err != nil {
		log.Fatalln("Ошибка при попытке открыть файл по пути: /sys/class/power_supply/BAT0/capacity\n", err)
	}

	return strings.TrimSpace(string(data))
}
