package files

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func DoFiles0() {
	dataForFile := []byte("Тестовая строка, предназначенная для записи в файл")
	file_name := "./files/test_file.txt"

	// Создаем новый файл и записываем в него данные dataForFile
	if err := os.WriteFile(file_name, dataForFile, 0600); err != nil {
		log.Fatal(err)
	}

	// Читаем данные из того же файла
	dataFromFile, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}

	// Сравниваем исходные данные с записанными в файл и прочитанными из него
	fmt.Printf("dataForFile == dataFromFile: %v\n", bytes.Equal(dataFromFile, dataForFile))

	// Получаем текущую директорию
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println((currentDir))

	// Изучаем содержимое директории
	filesFromDir, err := os.ReadDir(currentDir)
	if err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, file := range filesFromDir {
		// Проходим по всем найденным файлам и печатаем их имя и размер
		info, _ := file.Info()
		fmt.Printf("|_name: %s, size: %d\n", file.Name(), info.Size())
	}

	// Output:
	// dataForFile == dataFromFile: true
	// /home/<user>/<pwd>
	// |_name: main.go, size: 1491
	// |_name: test.txt, size: 93
	// |_...
}
