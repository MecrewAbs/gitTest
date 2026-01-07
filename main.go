package main

import (
	"fmt"
	"gitTest/api"
	"gitTest/bins"
	"gitTest/file"
	"gitTest/storage"
)

func main() {
	// 1. Создаём хранилище
	store := storage.NewStorage()

	// 2. Создаём bin
	bin := bins.NewBin("mybin123", false)
	bin.SetPrivate()
	fmt.Println("Bin:", bin)

	// 3. Сохраняем в хранилище
	err := store.CreateBin(bin)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	// 4. Сериализуем в JSON и сохраняем в файл
	jsonStr, err := bin.ToJSON()
	if err != nil {
		fmt.Println("Ошибка JSON:", err)
		return
	}

	err = file.WriteFile("bin.json", []byte(jsonStr))
	if err != nil {
		fmt.Println("Ошибка записи файла:", err)
	} else {
		fmt.Println("Bin сохранён в bin.json")
	}

	// 5. Читаем из файла
	content, err := file.ReadFile("bin.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	} else {
		fmt.Println("Из файла:", string(content))
	}

	// 6. Работаем с API (если есть ключ)
	client := api.NewClient("ВАШ_API_КЛЮЧ")
	data := map[string]interface{}{
		"Name":    bin.Name,
		"Private": bin.Private,
	}

	id, err := client.CreateBin(data)
	if err != nil {
		fmt.Println("Ошибка API:", err)
	} else {
		fmt.Println("Bin создан на JSON-BIN, ID:", id)
	}

	retievedData, err := client.GetBin(id)
	if err != nil {
		fmt.Println("Ошибка получения из API:", err)
	} else {
		fmt.Println("Получено из API:", retievedData)
	}
}
