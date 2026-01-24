package main

import (
	"fmt"
	"gitTest/files"
	"gitTest/storage"
)

func main() {
	bins := []storage.Bin{
		{ID: "1", Value: "first"},
		{ID: "2", Value: "second"},
	}

	// Сохранение в файл
	err := storage.SaveToJSON(bins, "bins.json")
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	// Загрузка из файла
	loadedBins, err := storage.LoadFromJSON("bins.json")
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}

	fmt.Printf("Загружено: %+v\n", loadedBins)

	filename := "data.json"

	exists, err := files.IsFileExists(filename)
	if err != nil {
		fmt.Println("Ошибка проверки файла:", err)
		return
	}
	if !exists {
		fmt.Println("Файл не существует")
		return
	}

	// Проверка расширения
	if !files.IsJSONFile(filename) {
		fmt.Println("Файл не является JSON")
		return
	}

	// Чтение файла
	data, err := files.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Println(string(data))
}

//2dwdwd
