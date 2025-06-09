package power

import (
	"os"
	"log"
)

func Show() string {
	file, err := os.Open("/sys/class/power_supply/BAT0/capacity")
	if err != nil {
		log.Fatalln("Ошибка при попытке открыть файл по пути: /sys/class/power_supply/BAT0/capacity\n", err)
	}
	defer file.Close()

	stat, errStat := file.Stat()
	if errStat != nil {
		log.Fatalln("Ошибка при попытке прочесть файл по пути: /sys/class/power_supply/BAT0/capacity\n", errStat)
	}
	buf := make([]byte, stat.Size())
	_, errRead := file.Read(buf)
	if errRead != nil {
		log.Fatalln("Ошибка при попытке прочесть файл по пути: /sys/class/power_supply/BAT0/capacity\n", errRead)
	}
	return string(buf)
}
