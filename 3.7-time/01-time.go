package time

import (
	"fmt"
	"time"
)

func DoTime01() {
	// Получаем текущее время
	now := time.Now()

	// Создаем время с помощью конкретных значений
	currentTime := time.Date(
		2020,     // год
		time.May, // месяц
		15,       // день
		10,       // часы
		13,       // минуты
		12,       // секунды
		45,       // наносекунды
		time.UTC, // временная зона
	)

	// Создаем время, используя секунды и наносекунды, прошедшие с начала эпохи Unix
	unixTime := time.Unix(
		150000, // секунды
		1,      // наносекунды
	)

	// Форматируем и выводим время в строковом виде
	fmt.Println(now.Format("02-01-2006 15:04:05"))         // 15-05-2020 09:58:16
	fmt.Println(currentTime.Format("02-01-2006 15:04:05")) // 15-05-2020 10:13:12
	fmt.Println(unixTime.Format("02-01-2006 15:04:05"))    // 02-01-1970 22:40:00

	// Функция Parse парсит строку в соответствии с заданным шаблоном
	firstTime, err := time.Parse("2006/01/02 15-04", "2020/05/15 17-45")
	if err != nil {
		panic(err)
	}

	// LoadLocation находит временную зону в справочнике IANA
	// https://www.iana.org/time-zones
	loc, err := time.LoadLocation("Asia/Yekaterinburg")
	if err != nil {
		panic(err)
	}

	// Функция ParseInLocation парсит строку в указанной временной зоне
	secondTime, err := time.ParseInLocation("Jan 2 06 03:04:05pm", "May 15 20 05:45:10pm", loc)
	if err != nil {
		panic(err)
	}

	fmt.Println(firstTime.Format("02-01-2006 15:04:05"))  // 15-05-2020 17:45:00
	fmt.Println(secondTime.Format("01-02-2006 15:04:05")) // 05-15-2020 17:45:10
}
